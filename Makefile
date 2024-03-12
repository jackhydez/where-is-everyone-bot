build-prod:
	docker-compose --env-file env.list build
run-prod:
	docker-compose --env-file env.list up -d --no-deps
stop-and-remove-container:
	docker-compose --env-file env.list stop
	yes | docker-compose --env-file env.list rm
clean-images:
	docker images -q -f dangling=true | xargs --no-run-if-empty docker rmi