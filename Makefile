env:
	cp .env.example .env

bot:
	docker-compose up -d --build