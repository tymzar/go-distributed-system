# All-purpose backend distributed systems

Idea for this system is to achieve a fault-tolerant, fast, scalable and secure system.

## Proposed stack 
- Chi: A lightweight, idiomatic, and composable router for building Go HTTP services. Chi offers similar functionality to Gorilla Mux, such as URL parameters, middleware, and route grouping.
- Go Kit: A programming toolkit for building microservices in Go. Go Kit provides a set of packages and best practices, allowing your service to be more robust, scalable, and maintainable.
- Envoy: A high-performance, programmable L7 proxy and communication bus designed for large modern service-oriented architectures. Envoy is used to manage and route traffic between microservices, enabling load balancing, service discovery, and observability.
- gRPC: A high-performance, open-source universal RPC framework developed by Google. gRPC allows you to define service contracts using Protocol Buffers, which can then be used to generate client and server stubs in Go.
- etcd: A distributed, reliable key-value store for distributed key locking, storing configuration data, and service discovery. etcd is used for coordinating distributed systems, providing a simple and secure method for service discovery and configuration management.
- GORM: An ORM library for Go that supports a variety of databases. GORM can be used to perform CRUD operations, run migrations, and interact with the database in a more intuitive way.
- Viper: A complete configuration solution for Go applications. Viper allows you to manage your microservice's configuration using a variety of sources such as environment variables, configuration files, and remote key-value stores.
- Prometheus: A monitoring system and time-series database that can be used to track the performance of your microservices. Prometheus provides Go client libraries to instrument your code for monitoring purposes.
- Docker: A platform for developing, shipping, and running applications in containers. Docker allows you to package your Go microservice and its dependencies into a container that can be deployed and managed easily.

```
               +---------+
               |         |
               |  Client |
               |         |
               +----+----+ 
                    |
                    v
               +---------+
               |         |
               |  Envoy  |
               |         |
               +----+----+ 
                    |
                    v
               +---------+
               |         |
               +  Broker +
               |         |
               +----+----+
                    |
                    v
        +-----------+-----------+
        |           |           |
   +----+----+ +----+----+ +----+----+
   |         | |         | |         |
   | Worker1 | | Worker2 | | WorkerN |
   |         | |         | |         |
   +----+----+ +----+----+ +----+----+
        |           |           |
        +-----------+-----------+
                    |
                    v
             +------+------+
             |             |
             | Microservice|
             |             |
             +------+------+
                    |
                    v
              +-----+-----+
              |           |
              | Database  |
              |           |
              +-----------+
```