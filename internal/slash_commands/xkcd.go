package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/pkg/providers/xkcd"
	"github.com/zekrotja/ken"
)

type XkcdCommand struct{}

var (
	_ ken.SlashCommand = (*XkcdCommand)(nil)
	_ ken.DmCapable    = (*XkcdCommand)(nil)
)

func (c *XkcdCommand) Name() string {
	return "xkcd"
}

func (c *XkcdCommand) Description() string {
	return "Returns an Xkcd comic."
}

func (c *XkcdCommand) Version() string {
	return "1.0.0"
}

func (c *XkcdCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *XkcdCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "latest",
			Description: "Returns the latest xkcd comic.",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "random",
			Description: "Returns a random xkcd comic.",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "specific",
			Description: "Returns the xkcd comic for the provided number",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "number",
					Description: "Number of the comic",
					Required:    true,
				},
			},
		},
	}
}

func (c *XkcdCommand) IsDmCapable() bool {
	return true
}

func (c *XkcdCommand) Run(ctx ken.Context) (err error) {
	err = ctx.HandleSubCommands(
		ken.SubCommandHandler{Name: "latest", Run: c.latest},
		ken.SubCommandHandler{Name: "random", Run: c.random},
		ken.SubCommandHandler{Name: "specific", Run: c.specific},
	)

	return
}

func (c *XkcdCommand) latest(ctx ken.SubCommandContext) (err error) {
	if err = ctx.Defer(); err != nil {
		return
	}
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

func (c *XkcdCommand) random(ctx ken.SubCommandContext) (err error) {
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

func (c *XkcdCommand) specific(ctx ken.SubCommandContext) (err error) {
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
