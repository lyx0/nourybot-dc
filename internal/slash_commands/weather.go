package slashcommands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/pkg/providers/weather"
	"github.com/zekrotja/ken"
	"go.uber.org/zap"
)

type WeatherCommand struct{}

var (
	_ ken.SlashCommand = (*WeatherCommand)(nil)
	_ ken.DmCapable    = (*WeatherCommand)(nil)
)

func (c *WeatherCommand) Name() string {
	return "weather"
}

func (c *WeatherCommand) Description() string {
	return "Returns the current weather for a given location."
}

func (c *WeatherCommand) Version() string {
	return "1.0.0"
}

func (c *WeatherCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *WeatherCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "location",
			Required:    true,
			Description: "Location",
		},
	}
}

func (c *WeatherCommand) IsDmCapable() bool {
	return true
}

func (c *WeatherCommand) Run(ctx ken.Context) (err error) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	val := ctx.Options().GetByName("location").StringValue()

	msg, err := weather.Get(val)
	if err != nil {
		sugar.Errorw("Error getting the weather:",
			"err", err)
	}

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
	return
}
