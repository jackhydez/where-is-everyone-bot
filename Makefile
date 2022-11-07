build:
	docker build -t where-is-everyone-bot .
run:
	docker run --name=where-is-everyone-bot-main -d --env-file env.list where-is-everyone-bot
all:
	docker build -t where-is-everyone-bot .
	docker run --env-file env.list where-is-everyone-bot 
clean:
	docker stop where-is-everyone-bot-main
	docker rm where-is-everyone-bot-main
	docker images -q -f dangling=true | xargs --no-run-if-empty docker rmi
	docker volume ls -qf dangling=true | xargs -r docker volume rm