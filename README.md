## Downloading project dependencies
To download the project dependencies, run the following commands:
```Shell
go mod download

go mod verify
```

## Starting Project with Go
Run this command to start the project.
```Shell
go run .
```

Note: the database must be running before starting the application.

## Starting environment with Docker Compose
To start all the environment:
```Shell
docker-compose up
```

### Starting only database
If you want to start only the database MySQL container:
```Shell
docker-compose up db
```

### Starting only app
If you want to start only the Golang application:
```Shell
docker-compose up app
```
