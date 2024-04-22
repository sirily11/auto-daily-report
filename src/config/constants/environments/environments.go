package environments

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	AdminApiKey = os.Getenv("ADMIN_API_KEY")
)
