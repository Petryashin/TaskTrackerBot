# TaskTrekerBot

Name - @MyDailyTaskTrackerBot

![Pipeline status](https://github.com/Petryashin/TaskTrackerBot/actions/workflows/main.yml/badge.svg?event=push&branch=main)

## Migration commands

1. Add new migration - ```migrate create -ext sql -dir migration -seq <name>```
2. Exec migration -
```migrate -path /migration -database postgres://<user>:<psswd>@localhost:5432/<dbname>?sslmode=disable up/down```
   in our case
```migrate -path ./migration -database postgres://app:app@localhost:5432/app?sslmode=disable up``` 
