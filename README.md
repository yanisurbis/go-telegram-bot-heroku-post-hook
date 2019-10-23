# Telegram bot for heroku deploy hooks

## General

- you should know `BOT_TOKEN`
- you should know `CHAT_ID` (the bot will push notifications there)
- to find out your `CHAT_ID` please check `https://api.telegram.org/bot{BOT_TOKEN}/getUpdates` and find a chat
you are interested in

## Development

- set `APP_ENV=dev`
- put `BOT_TOKEN` and `CHAT_ID` to `.env` file

## Deployment to now

- install `now.sh` with `npm i -g now`
- login with `now login`
- add secrets
```cmd
now secret add bot-token BOT_TOKEN
now secret add chat-id CHAT_ID
```
- deploy with `now -e BOT_TOKEN=@bot-token -e CHAT_ID=@chat-id`
- you will get the `BOT_URL`

## Connect to heroku

- go your app resources
- find add-on `deploy hooks`, choose `http post hook`
- add `{BOT_URL}/bot.go` if you use `now.sh` as a hook url
