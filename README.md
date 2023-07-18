# Env Server

## Endpoints

- /env returns environment variables on the server
- /env/<key> returns the value of a specific env variable key

## How to test

- Run the tests by running:

```sh
go test ./...
```

- If all tests pass, the output indicate that the tests have passed. if there is failure, the output will provide information about it.
