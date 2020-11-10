# GO BOT TELEGRAM

## Create your bot token
* create telegram bot search in your telegram `BotFather` and typing `/newbot` and enter
* typing your naming boot example tes_bot and enter `_bot` is required in your last name
* if your success your have recived message token save your token
* copy environment variable with your terminal
```sh
cp .env-sample .env
```
* fill the TELEGRAM_TOKEN your bot token in `.env`

## RUN APP

```go
go run server.go
```

## RUNNING WEBHOOK
* using `localhost.run`
* in your terminal
```sh
ssh -R 80:localhost:8080 ssh.localhost.run
```
after your running ssh yout have tunnel url 
ex `root-fa5d3df2.localhost.run`

running curl

```sh
curl -F "url=your-tunnel-localhost.run" https://api.telegram.org/botYOUR_BOT_TOKEN_HERE/setWebhook
```
