package slashcommands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/pkg/common"
	"github.com/zekrotja/ken"
)

type CoinflipCommand struct{}

var (
	_ ken.SlashCommand = (*CoinflipCommand)(nil)
	_ ken.DmCapable    = (*CoinflipCommand)(nil)
)

func (c *CoinflipCommand) Name() string {
	return "coinflip"
}

func (c *CoinflipCommand) Description() string {
	return "Basic Coinflip Command"
}

func (c *CoinflipCommand) Version() string {
	return "1.0.0"
}

func (c *CoinflipCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *CoinflipCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *CoinflipCommand) IsDmCapable() bool {
	return true
}

func (c *CoinflipCommand) Run(ctx ken.Context) (err error) {
	flip := common.GenerateRandomNumber(2)

	var msg string
	switch flip {
	case 0:
		msg = "Heads"
	case 1:
		msg = "Tails"
	default:
		msg = "Something went wrong :("
	}

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
	return
}
