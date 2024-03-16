build:
	docker-compose build
run:
	docker-compose up -d
stop:
	docker-compose stop
	yes | docker-compose rm
clean:
	docker images -q -f dangling=true | xargs --no-run-if-empty docker rmi