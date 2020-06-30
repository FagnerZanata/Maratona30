FROM golang:1.14
ADD ./src/go /go/src/
WORKDIR /go/src/
COPY . .
RUN GOOS=linux go build main.go
EXPOSE 8081
ENTRYPOINT ["./main"]
