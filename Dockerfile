FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
ENV PATH="/go/bin:$PATH"
RUN export PATH

EXPOSE 5000
RUN swag init -g internal/app/handler/docs.go -o docs --parseDependency --parseInternal
CMD ["go", "run", "./cmd/main.go", "-docker", "-migrate", "-start"]
