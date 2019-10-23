package handler

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
)

var telegramBot *tgbotapi.BotAPI

func ValidateEnvVars() {
	if os.Getenv("APP_ENV") == "dev" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		panic(fmt.Errorf("BOT_TOKEN env is missing\n"))
	}

	chatIdStr := os.Getenv("CHAT_ID")
	if chatIdStr == "" {
		panic(fmt.Errorf("CHAT_ID env is missing\n"))
	}
}

func GetBotTokenAndChatId() (botToken string, chatId int64) {
	ValidateEnvVars()
	botToken = os.Getenv("BOT_TOKEN")
	chatIdStr := os.Getenv("CHAT_ID")

	chatIdInt, _ := strconv.Atoi(chatIdStr)
	chatId = int64(chatIdInt)

	return botToken, chatId
}

func GetBot(token string) *tgbotapi.BotAPI {
	if telegramBot == nil {
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
		return bot
	} else {
		fmt.Printf("Bot already created \n")
		return telegramBot
	}
}

func Handler(_ http.ResponseWriter, r *http.Request) {
	token, chatId := GetBotTokenAndChatId()
	bot := GetBot(token)
	err := r.ParseForm()
	if err != nil {
		return
	}

	fmt.Printf("Params: %v\n", r.Form)

	keys := []string{"app", "url", "release", "user", "head", "head_long", "prev_head", "git_log"}

	message := "Deployment completed \n\n"
	for _, key := range keys {
		value := r.Form.Get(key)
		if value != "" {
			message += key + ": " + value + "\n"
		}
	}

	_, err = bot.Send(tgbotapi.NewMessage(chatId, message))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	_ = http.ListenAndServe(":8080", mux)
}
