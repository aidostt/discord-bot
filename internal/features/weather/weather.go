package weather

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func (cmd *WeatherCommand) Description() string {
	return "this command returns weather in desired city. Usage -weather <city>"
}

// Execute fetches and displays weather information.
func (cmd *WeatherCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: `-weather <city>`")
		return
	}

	city := args[1]
	// Inside the Execute method
	weather, err := cmd.Client.GetWeather(city)
	if err != nil {
		// Adjust the error message based on the context
		var errMsg string
		if err.Error() == "request timed out or city not found" {
			errMsg = "Invalid input or city not found. Please check your input and try again."
		} else {
			errMsg = fmt.Sprintf("Failed to get weather: %v", err)
		}
		s.ChannelMessageSend(m.ChannelID, errMsg)
		return
	}

	response := fmt.Sprintf("Current weather in %s: %.2f°C", weather.Name, weather.Main.Temp)
	s.ChannelMessageSend(m.ChannelID, response)

}
