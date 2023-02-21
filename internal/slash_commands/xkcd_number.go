package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/pkg/providers/xkcd"
	"github.com/zekrotja/ken"
)

type XkcdNumberCommand struct{}

var (
	_ ken.SlashCommand = (*XkcdNumberCommand)(nil)
	_ ken.DmCapable    = (*XkcdNumberCommand)(nil)
)

func (c *XkcdNumberCommand) Name() string {
	return "xkcd_number"
}

func (c *XkcdNumberCommand) Description() string {
	return "Returns a specific xkcd comic by number (from 1 to around 2700)."
}

func (c *XkcdNumberCommand) Version() string {
	return "1.0.0"
}

func (c *XkcdNumberCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *XkcdNumberCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        "number",
			Required:    true,
			Description: "Comic number",
		},
	}
}

func (c *XkcdNumberCommand) IsDmCapable() bool {
	return true
}

func (c *XkcdNumberCommand) Run(ctx ken.Context) (err error) {
	comicNum := ctx.Options().GetByName("number").IntValue()

	num, title, img := xkcd.Number(fmt.Sprint(comicNum))
	msg := fmt.Sprint("Xkcd comic number #", num, " Title: ", title, " ", img)

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
	return
}
