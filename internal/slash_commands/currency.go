package slashcommands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/pkg/providers/currency"
	"github.com/zekrotja/ken"
	"go.uber.org/zap"
)

type CurrencyCommand struct{}

var (
	_ ken.SlashCommand = (*CurrencyCommand)(nil)
	_ ken.DmCapable    = (*CurrencyCommand)(nil)
)

func (c *CurrencyCommand) Name() string {
	return "currency"
}

func (c *CurrencyCommand) Description() string {
	return "Returns the current exchange rate for two given currencies."
}

func (c *CurrencyCommand) Version() string {
	return "1.0.0"
}

func (c *CurrencyCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *CurrencyCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "amount",
			Required:    true,
			Description: "Amount of money to convert in three letter abbreviation (USD, EUR, PLN)",
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "from",
			Required:    true,
			Description: "Currency from which to convert from in three letter abbreviation (USD, EUR, PLN)",
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "to",
			Required:    true,
			Description: "Currency to convert to",
		},
	}
}

func (c *CurrencyCommand) IsDmCapable() bool {
	return true
}

func (c *CurrencyCommand) Run(ctx ken.Context) (err error) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	currAmount := ctx.Options().GetByName("amount").StringValue()
	currFrom := ctx.Options().GetByName("from").StringValue()
	currTo := ctx.Options().GetByName("to").StringValue()

	msg, err := currency.Convert(currAmount, currFrom, currTo)
	if err != nil {
		sugar.Errorw("Error converting currencies:",
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
