FROM golang:1.20-alpine as build
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# Install time zone data
RUN apk add --no-cache tzdata

# Set the desired time zone
ENV TZ=Asia/Bangkok

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:3.13.0
WORKDIR /app
COPY --from=build app/.env .
COPY --from=build app/main .

EXPOSE 8000

ENTRYPOINT ./main ${VALIDATOR_ADDR} ${RPC} ${WEBHOOK} 
