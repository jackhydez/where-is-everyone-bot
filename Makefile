C=1
build:
	docker build -t where-is-everyone-bot .
run:
	docker run --name=where-is-everyone-bot -d --env-file env.list where-is-everyone-bot
clean:
	docker stop where-is-everyone-bot
	docker rm where-is-everyone-bot
	docker images -q -f dangling=true | xargs --no-run-if-empty docker rmi
run-prod:
	docker run --name=where-is-everyone-bot-$(BUILD_NUMBER) -d --env-file env.list where-is-everyone-bot
clean-containers:
	docker stop where-is-everyone-bot-$$(($(BUILD_NUMBER) - $(C)))
	docker rm where-is-everyone-bot-$$(($(BUILD_NUMBER) - $(C)))
clean-images:
	docker images -q -f dangling=true | xargs --no-run-if-empty docker rmi