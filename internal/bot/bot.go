package bot

import (
	"discord-bot.aidostt.me/internal/commands"
	"discord-bot.aidostt.me/pkg/config"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

type Bot struct {
	Session  *discordgo.Session
	Commands *commands.CommandRegistry
}

func NewBot(cfg *config.Config) (*Bot, error) {
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		Session:  dg,
		Commands: commands.NewCommandRegistry(),
	}

	// Register all command handlers here
	commands.RegisterCommands(dg, bot.Commands)

	return bot, nil
}

func (b *Bot) Start() error {
	b.setupHandlers()
	return b.Session.Open()
}

func (b *Bot) AwaitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop

	b.Session.Close()
}

func (b *Bot) setupHandlers() {
	// Add the messageCreate function as a handler for MessageCreate events.
	b.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore messages sent by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}
		// Delegate to the CommandRegistry to handle the message
		b.Commands.HandleMessage(s, m)
	})
}
