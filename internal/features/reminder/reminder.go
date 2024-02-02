package reminder

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

type ReminderCommand struct {
	//TODO:Implement storage or schedule interface
}

// NewReminderCommand creates a new instance of the ReminderCommand
func NewReminderCommand() *ReminderCommand {
	return nil
	//	return &ReminderCommand{
	//		CommandRegistry: registry,
	//		// Initialize any necessary fields
	//	}
}

// Execute method for the ReminderCommand
func (cmd *ReminderCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// Splitting the message content to parse command arguments
	if len(args) < 3 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Usage: `-reminder <duration> <message>`. Example: `-reminder 10m Take a break`")
		return
	}

	duration, err := time.ParseDuration(args[1])
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid duration. Please specify a valid duration (e.g., 30s, 10m, 1h).")
		return
	}

	reminderMessage := strings.Join(args[2:], " ")
	userID := m.Author.ID // Capture the user's ID

	// Use a goroutine for the delay
	go func() {
		time.Sleep(duration)

		reminderText := fmt.Sprintf("<@%s>, Reminder: %s", userID, reminderMessage)
		_, _ = s.ChannelMessageSend(m.ChannelID, reminderText)
	}()

	_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("I'll remind you in %s: %s", duration, reminderMessage))
}

// Description returns the description of the ReminderCommand
func (cmd *ReminderCommand) Description() string {
	return "sets a reminder. Usage: `-reminder <time> <message>`"
}

// parseArgs is a helper function to parse the command arguments
func parseArgs(content string) []string {
	return strings.Fields(content)
}
