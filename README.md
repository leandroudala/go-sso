# Opening the project using devcontainer
Run the following command to start the project:
```Shell
devcontainer up
```

## Installing the devcontainer
To install the devcontainer globally, please execute the following command:
```Shell
npm install -g @devcontainers/cli
```

# Opening the project locally
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

## Update Swagger
Run the following command:
```Shell
swag init
```

Note: if the console reports that the command `swag` was not found, check where the Go is installed.
If you are using asdf, run the following command to find out where is the Golang bin folder:
```Shell
asdf where golang
```


# References:
devcontainer: https://github.com/devcontainers/cli