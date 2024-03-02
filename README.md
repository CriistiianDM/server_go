# Go web server canvas

This is a server oriented towards abstraction rather than implementation, with a clean and scalable architecture. It features a PostgreSQL database connection with generalized queries using advanced SQL concepts and offers some security by returning encrypted responses.

# Architecture
![Tablero en blanco](https://github.com/CriistiianDM/server_go/assets/62184928/1c10a53d-78f2-4c9f-9143-52dc0dbddb50)

# Technologies
* Golang
* Postgresql
* Docker
* Kubernetes

# Future requirements
* Proxys
* Middlewares
* Deploy to google cloud whit k8

# How to run with Docker

```
  docker-compose -p project_name up
```

# HOW TO INSTALL LOCALLY

* Step 1:
  
  ```
     go mod tidy
  ```

* Step 2:
  ```
     mv .env.local .env
  ```

* Step 3:
  ```
     psql -U postgres -h localhost -a -f /route/to/archive/sql.sql -W
  ```
  
* Step 3:
  ```
     go run ./src/main.go
  ```
