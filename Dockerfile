FROM golang:1.20

WORKDIR /app


COPY . ./


RUN go build -o app cmd/main.go


EXPOSE 8080

CMD [ "./app","-p","8080" ]