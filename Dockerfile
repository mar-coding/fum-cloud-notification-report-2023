FROM golang:1.17-alpine AS BuildStage 

ENV GOPROXY=https://goproxy.io,direct 
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 1234

# Deploy stage

FROM alpine:3.14

WORKDIR /app

COPY --from=BuildStage /app/main .
# COPY --from=BuildStage /app/envs /app/envs

EXPOSE 1234

CMD ["./main"]
LABEL maintainer="MohammadAmin Rahimi <marcoding78@gmail.com>"