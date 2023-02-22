package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/pkg/common"
	"github.com/zekrotja/ken"
)

type TwitchCommand struct{}

var (
	_ ken.SlashCommand = (*TwitchCommand)(nil)
	_ ken.DmCapable    = (*TwitchCommand)(nil)
)

func (c *TwitchCommand) Name() string {
	return "twitch"
}

func (c *TwitchCommand) Description() string {
	return "Returns information about your user account"
}

func (c *TwitchCommand) Version() string {
	return "1.0.0"
}

func (c *TwitchCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *TwitchCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "preview",
			Description: "Returns a screenshot of a currently live channel.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "channel",
					Description: "Twitch username of a live channel",
					Required:    true,
				},
			},
		},
	}
}

func (c *TwitchCommand) IsDmCapable() bool {
	return true
}

func (c *TwitchCommand) Run(ctx ken.Context) (err error) {
	err = ctx.HandleSubCommands(
		ken.SubCommandHandler{Name: "preview", Run: c.preview},
	)

	return
}
func (c *TwitchCommand) preview(ctx ken.SubCommandContext) (err error) {
	channel := ctx.Options().GetByName("channel").StringValue()
	imageHeight := common.GenerateRandomNumberRange(1040, 1080)
	imageWidth := common.GenerateRandomNumberRange(1890, 1920)

	link := fmt.Sprintf("https://static-cdn.jtvnw.net/previews-ttv/live_user_%v-%vx%v.jpg", channel, imageWidth, imageHeight)
	err = ctx.RespondEmbed(&discordgo.MessageEmbed{
		Title: fmt.Sprintf("Preview of %v's stream", channel),
		Image: &discordgo.MessageEmbedImage{
			URL: link,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Links",
				Value: fmt.Sprintf("- [Link](%v)\n", link) +
					fmt.Sprintf("- [Twitch](https://twitch.tv/%v)\n", channel),
			},
		},
		//	Footer: &discordgo.MessageEmbedFooter{
		//		Text: "Data might be outdated",
		//	},
	})
	// err = ctx.Respond(&discordgo.InteractionResponse{
	// 	Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 	Data: &discordgo.InteractionResponseData{
	// 		Content: msg,
	// 	},
	// })
	return
}
