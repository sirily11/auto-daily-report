package environments

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type RunMode string

const (
	ModeProduction RunMode = "production"
)

var (
	AdminApiKey            = os.Getenv("ADMIN_API_KEY")
	RunEnvironment         = RunMode(os.Getenv("RUN_ENV"))
	TelegramBotApiKey      = os.Getenv("TELEGRAM_BOT_API_KEY")
	TelegramChatEndpoint   = os.Getenv("TELEGRAM_CHAT_ENDPOINT")
	MongoDBDataAPIEndpoint = os.Getenv("MONGODB_DATA_API_ENDPOINT")
	MongoDBDataAPIKey      = os.Getenv("MONGODB_DATA_API_KEY")
	MongoDBDataSource      = os.Getenv("MONGODB_DATA_SOURCE")
	OpenAIApiKey           = os.Getenv("OPENAI_API_KEY")
)
