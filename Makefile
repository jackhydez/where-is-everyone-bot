build:
	docker stop $(docker ps -a -q)
	docker rm $(docker ps -a -q)
	docker rmi $(docker images -a -q)
	
	docker build -t where-is-everyone-bot .
run:
	docker run -d --env-file env.list where-is-everyone-bot
all:
	docker build -t where-is-everyone-bot .
	docker run --env-file env.list where-is-everyone-bot 