# read-redis
Implementation of nsq handling redis data update

## Setting Up
To run, clone this project and use
```bash
docker-compose up
```

## Environment Variables
Currently we need the following in an `.env` file:
```
redisConn=redis:6379

dbHost=db
dbPort=5432
dbUser=postgres
dbPass=password
dbName=postgres

messageChannel=message_channel
messageTopic=message_topic
nsqdAddr=nsqd:4151
lookupAddr=nsqlookupd:4161
```

There is definitely a better way to do it. Feel free to create a PR on improving it!