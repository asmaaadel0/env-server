FROM golang:1.20

WORKDIR /app

COPY . ./

RUN go build -o bin/app cmd/main.go

EXPOSE 8080

CMD [ "./cmd/app","-p","8080" ]