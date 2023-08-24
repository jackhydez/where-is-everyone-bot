# Project telegram-bot where-is-everyone

create file for environment variables:
```
cp env-example.list env.list
```

build and run all:
```
docker-compose --env-file env.list up --build
```

stop and remove all:
```
docker-compose --env-file env.list down
```

link on my dev:
```
t.me/where_is_everyone_dev_bot
```

link on my prod:
```
t.me/where_is_everyone_bot
```