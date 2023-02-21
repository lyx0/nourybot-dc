package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/pkg/providers/xkcd"
	"github.com/zekrotja/ken"
)

type XkcdRandomCommand struct{}

var (
	_ ken.SlashCommand = (*XkcdRandomCommand)(nil)
	_ ken.DmCapable    = (*XkcdRandomCommand)(nil)
)

func (c *XkcdRandomCommand) Name() string {
	return "xkcd_random"
}

func (c *XkcdRandomCommand) Description() string {
	return "Returns a random xkcd comic."
}

func (c *XkcdRandomCommand) Version() string {
	return "1.0.0"
}

func (c *XkcdRandomCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *XkcdRandomCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *XkcdRandomCommand) IsDmCapable() bool {
	return true
}

func (c *XkcdRandomCommand) Run(ctx ken.Context) (err error) {
	num, title, img := xkcd.Random()
	msg := fmt.Sprint("Random Xkcd comic number #", num, " Title: ", title, " ", img)

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
	return
}
