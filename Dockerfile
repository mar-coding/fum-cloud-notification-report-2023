FROM golang:1.20.4-alpine3.18 AS BuildStage 

ENV GOPROXY=https://goproxy.io,direct 
WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 1234

# Deploy stage

FROM alpine:3.18

WORKDIR /app

COPY --from=BuildStage /app/main .
# COPY --from=BuildStage /app/.env .

EXPOSE 1234

CMD ["./main"]
LABEL maintainer="MohammadAmin Rahimi <marcoding78@gmail.com>"