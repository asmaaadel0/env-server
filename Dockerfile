FROM golang:1.20

WORKDIR /app


COPY . ./


RUN go build -o cmd/app cmd/main.go


EXPOSE 8080

CMD [ "./dmc/app","-p","8080" ]
