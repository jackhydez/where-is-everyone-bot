build:
	docker build -t where-is-everyone-bot .
run:
	docker run -d --env-file env.list --rm where-is-everyone-bot
	docker rmi $(docker images -a -q)
all:
	docker build -t where-is-everyone-bot .
	docker run --env-file env.list where-is-everyone-bot 