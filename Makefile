build:
	docker build -t where-is-everyone-bot .
run:
	docker run -d --env-file env.list where-is-everyone-bot
all:
	docker build -t where-is-everyone-bot .
	docker run --env-file env.list where-is-everyone-bot 