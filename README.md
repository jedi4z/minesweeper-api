# Minesweeper-API

This is API rest written in GO that allows you to play the classic game [Minesweeper](https://en.wikipedia.org/wiki/Minesweeper_(video_game))

# Top level project structure
- **app/adapter:** this is the interface between your application and outside data service, for example another Rest or gRPC service. All the data conversion and transformation happened here, so your business logic code doesn’t need to be aware of the detail implementation (whether it gRPC or REST) of outside services.
- **app/container:** the dependency injection container, which is responsible for creating concrete types and injecting them into each function.
- **app/models:** domain module layer, which has domain structs. All other layers depend on them and they don’t depend on any other layers.
- **app/repositories:** persistence layer, which is responsible for retrieving and modifying data for the domain model. It only depends on the model layer.
- **app/usecases:** This is an important layer and the entry point of business logic. Each business feature is implemented by a use case. It is the top level layer, so no other layer depends on it ( except “cmd”), but it depends on other layers.
- **cmd:** the command. All different types of “main.go” are here and you can have multiple ones. This is the starting point of the application.

# How to run the project
```
# Start the service
make run

# Build and start service using docker-compose
make up

# Delete docker containers
make down
```
As an alternative you can use docker-compose command

```
docker-compose up
docker-compose down
```

# API
Here are the endpoints that need to be used to play the game, also there are a Postman collection and environment to test the API.
Please, see [the postman folder](https://github.com/jedi4z/minesweeper-api/tree/master/postman). 

## User

### Register a User
```http request
POST /v1/users/register HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Content-Type: application/json

{ 
    "email": "jesusdiazbc@gmail.com",
    "password": "demo"
}
```

### Authenticate User
```http request
POST /v1/users/auth HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Content-Type: application/json

{ 
    "email": "jesusdiazbc@gmail.com",
    "password": "demo"
}
```

## Game

### Create a Game
```http request
POST /v1/games HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Authorization: Bearer <access_token>
Content-Type: application/json

{
    "number_of_rows": 15,
    "number_of_cols": 15,
    "number_of_mines": 20
}
```

### Retrieve a Game
```http request
GET /v1/games/<game_id> HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Authorization: Bearer <access_token>

```

### List Games
```http request
GET /v1/games HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Authorization: Bearer <access_token>
Content-Type: application/json

```

### Hold a Game
```http request
POST /v1/games/<game_id>/hold HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Authorization: Bearer <access_token>

```

### Resume a Game
```http request
POST /v1/games/<game_id>/resume HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Authorization: Bearer <access_token>

```

### Uncover a Cell
```http request
POST /v1/games/<game_id>/uncover/<cell_id> HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Authorization: Bearer <access_token>

```

### Flag a Cell
```http request
POST /v1/games/<game_id>/flag/<cell_id> HTTP/1.1
Host: https://minesweeper-jd.herokuapp.com
Authorization: Bearer <access_token>

```

# Client Library (Python)
I created a python library to make requests to the API. For more information please go to this repository: https://github.com/jedi4z/minesweeper-api-lib

# TODO
- Add a configuration module using [Viper](https://github.com/spf13/viper)
- Testing (unit tests and functional tests).
- Document the API using swagger.
- Configure CI/CD with any tool like CircleCI, Jenkins or Github Actions.  
