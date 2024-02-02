package weather

import (
	"discord-bot.aidostt.me/internal/api"
	"strings"
)

// WeatherCommand handles weather queries.
type WeatherCommand struct {
	Client *api.WeatherClient
}

// NewWeatherCommand creates a new instance of WeatherCommand.
func NewWeatherCommand() *WeatherCommand {
	return &WeatherCommand{
		Client: api.NewWeatherClient(),
	}
}

// splitArgs is a helper function to parse command arguments.
func splitArgs(content string) []string {
	return strings.Fields(content)
}
