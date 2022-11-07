build:
	docker build -t where-is-everyone-bot .
run:
	docker run --name=where-is-everyone-bot-main -d --env-file env.list where-is-everyone-bot
clean-containers:
	docker stop where-is-everyone-bot-main
	docker rm where-is-everyone-bot-main
clean-images:
	docker images -q -f dangling=true | xargs --no-run-if-empty docker rmi