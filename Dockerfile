#Golang image
FROM golang:1.19.8

WORKDIR /server_go

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./src/main.go

EXPOSE 4700

CMD ["./server"]
