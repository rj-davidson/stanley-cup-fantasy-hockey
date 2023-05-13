FROM golang:latest
WORKDIR /app/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .
RUN go generate ./db/ent && go build -o main .
EXPOSE 8080
CMD ["/app/main"]