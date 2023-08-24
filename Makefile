build-prod:
	docker-compose --env-file env.list build backend
run-prod:
	docker-compose --env-file env.list up -d --no-deps backend
stop-and-remove-container:
	docker-compose --env-file env.list stop backend
	yes | docker-compose --env-file env.list rm backend 
clean-images:
	docker images -q -f dangling=true | xargs --no-run-if-empty docker rmi