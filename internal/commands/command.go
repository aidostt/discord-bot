package commands

import (
	"discord-bot.aidostt.me/internal/features/game"
	"discord-bot.aidostt.me/internal/features/reminder"
	"discord-bot.aidostt.me/internal/features/weather"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
	Description() string
}

type CommandRegistry struct {
	commands map[string]Command
}

func (r *CommandRegistry) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.TrimSpace(m.Content)
	if strings.HasPrefix(content, "-") {
		args := strings.Fields(content[1:])
		if len(args) > 0 {
			if command, exists := r.commands[args[0]]; exists {
				callServiceAsync(s, m, func(s *discordgo.Session, m *discordgo.MessageCreate) {
					command.Execute(s, m, args)
				})
			}
		}
	}
}

func RegisterCommands(session *discordgo.Session, registry *CommandRegistry) {
	registry.Register("help", NewHelpCommand(registry))
	registry.Register("weather", weather.NewWeatherCommand())
	registry.Register("reminder", reminder.NewReminderCommand())
	registry.Register("tictactoe", game.NewTicTacToeCommand())

}
