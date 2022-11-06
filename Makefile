build:
	docker build -t where-is-everyone-bot .
run:
	docker run -d --env-file env.list where-is-everyone-bot 
stop:
	docker stop $(docker ps -a -q)
	docker rm $(docker ps -a -q)
	docker rmi $(docker images -a -q)
all:
	make stop
	docker build -t where-is-everyone-bot .
	docker run --env-file env.list where-is-everyone-bot 