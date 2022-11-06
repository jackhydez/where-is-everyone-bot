build:
	docker build -t where-is-everyone-bot-$(BUILD_NUMBER) .
run:
	docker run -d --env-file env.list -name where-is-everyone-bot-$(BUILD_NUMBER)
all:
	docker build -t where-is-everyone-bot .
	docker run --env-file env.list where-is-everyone-bot 