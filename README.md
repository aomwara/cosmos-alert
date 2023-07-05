# Node X - Watch signature Bot

- Build binary
```sh
go build
```

- Run
```sh
# Run dev
go run main.go <validator_pubkey> <rpc> <discord_webhook>

# Run Prod
./alertbot <validator_pubkey> <rpc> <discord_webhook>
```

- Build with Dockerfile
```sh
docker build -t alertbot:latest .
```

- Run with Docker
```sh
# make .env file
make env

# edit .env file
nano .env

# run with docker compose
docker compose up -d --build
```