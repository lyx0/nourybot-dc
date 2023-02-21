package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type InfoCommand struct{}

var (
	_ ken.SlashCommand = (*InfoCommand)(nil)
	_ ken.DmCapable    = (*InfoCommand)(nil)
)

func (c *InfoCommand) Name() string {
	return "info"
}

func (c *InfoCommand) Description() string {
	return "Returns information about your user account"
}

func (c *InfoCommand) Version() string {
	return "1.0.0"
}

func (c *InfoCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *InfoCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *InfoCommand) IsDmCapable() bool {
	return true
}

func (c *InfoCommand) Run(ctx ken.Context) (err error) {
	err = ctx.RespondEmbed(&discordgo.MessageEmbed{
		Title: fmt.Sprintf("Info for %v", ctx.User().Username),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:    ctx.User().AvatarURL("64x64"),
			Width:  64,
			Height: 64,
		},
		Description: fmt.Sprintf("Username: %v#%v\nUser ID: %v\n", ctx.User().Username, ctx.User().Discriminator, ctx.User().ID),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Links",
				Value: fmt.Sprintf("- [Avatar URL](%v)\n", ctx.User().AvatarURL("")) +
					fmt.Sprintf("- [Discord Lookup](https://discordlookup.com/user/%v)\n", ctx.User().ID),
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
