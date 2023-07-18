# Env Server

## Endpoints

- /env returns environment variables on the server
- /env/<key> returns the value of a specific env variable key

## Installation

- Clone the repository:

```sh
$ git clone https://github.com/codescalersinternships/enserver-Asmaa.git
```

- Change into the project directory:

```sh
$ cd envserver-Asmaa
```

- Build the project:

```sh
go build -o "bin/server" cmd/main.go
```

- Change into the project directory:

```sh
$ cd bin
```

- Run the server.

```sh
./server -p [port]
```

## How to use

- To get all the environment variables open:

```sh
 http://localhost:8080/env
```

- To get environment variable with specific "key" open:

```sh
 http://localhost:8080/env/key
```

## How to use Docker

## How to test

- Run the tests by running:

```sh
go test ./...
```

- If all tests pass, the output indicate that the tests have passed. if there is failure, the output will provide information about it.
