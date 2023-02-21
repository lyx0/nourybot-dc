package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/pkg/providers/xkcd"
	"github.com/zekrotja/ken"
)

type XkcdLatestCommand struct{}

var (
	_ ken.SlashCommand = (*XkcdLatestCommand)(nil)
	_ ken.DmCapable    = (*XkcdLatestCommand)(nil)
)

func (c *XkcdLatestCommand) Name() string {
	return "xkcd_latest"
}

func (c *XkcdLatestCommand) Description() string {
	return "Returns the latest xkcd comic."
}

func (c *XkcdLatestCommand) Version() string {
	return "1.0.0"
}

func (c *XkcdLatestCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *XkcdLatestCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *XkcdLatestCommand) IsDmCapable() bool {
	return true
}

func (c *XkcdLatestCommand) Run(ctx ken.Context) (err error) {

	num, title, img := xkcd.Latest()
	msg := fmt.Sprint("Latest Xkcd comic number #", num, " Title: ", title, " ", img)

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
	return
}
