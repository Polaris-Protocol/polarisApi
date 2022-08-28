# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.16-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download
 
COPY *.go ./

RUN go build -o /polaris-api

EXPOSE 6000

CMD [ "/polaris-api"]

