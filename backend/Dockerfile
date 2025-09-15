FROM golang:1.25

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build main.go

EXPOSE 3001

CMD ["./main"]
