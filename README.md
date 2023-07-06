# Bot for downloading active instagram stories from your feed
## Launch
Requires GoLang and PostgreSQL on your machine

1. Replace `{your telegram API token}` with your telegram bot token in the 10th line of [cmd/bot/main.go](cmd/bot/main.go/#L10) file
2. Replace brackets -> {} with your data in 40th line of [pkg/telegram/handlers.go](pkg/telegram/handlers.go) file
3. Run `make build && make run` command in the root directory of the project
