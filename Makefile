all:
	sudo docker build -t where-is-everyone-bot .
	sudo docker run --env-file env.list where-is-everyone-bot &