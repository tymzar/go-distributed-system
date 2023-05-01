package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"os"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file", err)
	}

	fmt.Println(viper.Get("DATABASE_HOST"))
}

func main() {
	host := viper.GetString("DATABASE_HOST")
	port := viper.GetInt("DATABASE_PORT")
	//user := viper.GetString("DATABASE_USER")
	password := viper.GetString("DATABASE_PASSWORD")
	//dbname := viper.GetString("DATABASE_DBNAME")
	//sslmode := viper.GetString("DATABASE_SSLMODE")

	//connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

	connStr := fmt.Sprintf("postgres://postgres:%s@%s:%d/postgres", password, host, port)

	fmt.Println(connStr)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Connected")
	}
	defer conn.Close(context.Background())

	// Your code here
}
