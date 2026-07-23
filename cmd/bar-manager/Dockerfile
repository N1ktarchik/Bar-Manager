FROM golang:1.25.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/cmd/bar-manager/exe ./cmd/bar-manager/main.go

CMD ["/app/cmd/bar-manager/exe"]