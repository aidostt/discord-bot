package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type HelpCommand struct {
	CommandRegistry *CommandRegistry
}

func (c *HelpCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, c.CommandRegistry.ListCommands())
}

func (r *CommandRegistry) ListCommands() string {
	var result strings.Builder
	for name, command := range r.commands {
		result.WriteString("-" + name + ": " + command.Description() + "\n")
	}
	return result.String()
}

func NewHelpCommand(registry *CommandRegistry) *HelpCommand {
	return &HelpCommand{CommandRegistry: registry}
}

func (c *HelpCommand) Description() string {
	return "displays all possible commands. Usage: -help"
}
