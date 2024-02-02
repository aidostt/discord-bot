package commands

import (
	"github.com/bwmarrin/discordgo"
	"sync"
)

var mu sync.Mutex

func callServiceAsync(s *discordgo.Session, m *discordgo.MessageCreate, service func(s *discordgo.Session, m *discordgo.MessageCreate)) {
	go func() {
		mu.Lock()
		defer mu.Unlock()
		service(s, m)
	}()
}

func (r *CommandRegistry) Register(name string, command Command) {
	r.commands[name] = command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{commands: make(map[string]Command)}
}
