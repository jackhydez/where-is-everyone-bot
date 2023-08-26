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

solution for:  FATAL:  role "postgres" does not exist
```
docker exec -it postgres_container bash

psql postgres

\du

CREATE USER postgres SUPERUSER;

CREATE DATABASE postgres WITH OWNER postgres;

\du

\q

exit
```