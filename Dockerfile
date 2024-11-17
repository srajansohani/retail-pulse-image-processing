
FROM golang:1.23.3-alpine
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /main main.go
ENTRYPOINT [ "/main" ]
