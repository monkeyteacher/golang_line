#!/bin/bash
go get -u github.com/gin-gonic/gin
go get -u github.com/line/line-bot-sdk-go/v7/linebot
go get github.com/spf13/viper
go get -u github.com/spf13/cobra@latest
go get go.mongodb.org/mongo-driver/mongo

read -p "Please input your mongodb username: " username      # mongodb username
read -p "Please input your mongodb password: " password       # mongodb password

docker run --name mongodb -d  -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=${username} -e MONGO_INITDB_ROOT_PASSWORD=${password} mongo:4.4