build:
	docker build -t where-is-everyone-bot .
run:
	docker run --name=where-is-everyone-bot-name -d --env-file env.list --rm where-is-everyone-bot
all:
	docker build -t where-is-everyone-bot .
	docker run --env-file env.list where-is-everyone-bot 
clean:
	docker stop where-is-everyone-bot-name