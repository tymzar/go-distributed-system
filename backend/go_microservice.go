package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"github.com/tymzar/go-distributed-system/proto/hello"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var conn *pgx.Conn
var connMutex = &sync.Mutex{}

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
		viper.Set("IS_DOCKER", "false")
	} else {
		fmt.Println("Initialize DB connection")
	}
}

type myHelloServiceServer struct {
	hello.UnimplementedHelloServiceServer
}

func (s *myHelloServiceServer) Create(ctx context.Context, req *hello.CreateRequest) (*hello.CreateResponse, error) {
	connMutex.Lock()
	defer connMutex.Unlock()

	if conn == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	// Select from database
	rows, err := conn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		log.Println("Error executing select query: ", err)
		return nil, err
	}
	defer rows.Close()

	fmt.Println("Current users:")
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println("Error scanning rows: ", err)
			return nil, err
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Insert into database
	_, err = conn.Exec(context.Background(), "INSERT INTO users(name) VALUES ($1)", req.GetName())
	if err != nil {
		log.Println("Error executing insert query: ", err)
		return nil, err
	}

	return &hello.CreateResponse{Message: "Welcome, " + req.GetName()}, nil
}

func main() {
	isDocker := viper.GetBool("IS_DOCKER")

	if isDocker {
		host := viper.GetString("DATABASE_HOST")
		port := viper.GetInt("DATABASE_PORT")
		password := viper.GetString("DATABASE_PASSWORD")

		connStr := fmt.Sprintf("postgres://postgres:%s@%s:%d/postgres", password, host, port)

		fmt.Println(connStr)

		var err error
		conn, err = pgx.Connect(context.Background(), connStr)
		if err != nil {
			log.Fatalf("Unable to connect to the database: %v", err)
		} else {
			fmt.Println("Connected")
			defer conn.Close(context.Background())
		}

		// Check if 'users' table exists, if not, create it
		_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT
		)`)
		if err != nil {
			log.Fatalf("Unable to create 'users' table: %v", err)
		}
	}

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("cannot connect to the listener: %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &myHelloServiceServer{}

	hello.RegisterHelloServiceServer(serverRegister, service)

	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
