# build:
# 	docker build -t where-is-everyone-bot .
# run:
# 	docker run --env-file env.list where-is-everyone-bot &
# stop:
# 	docker stop $(docker ps -a -q)
# 	docker rm $(docker ps -a -q)
# 	docker rmi $(docker images -a -q)
# all:
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker rmi $(docker images -a -q)
docker build -t where-is-everyone-bot .
docker run --env-file env.list where-is-everyone-bot &