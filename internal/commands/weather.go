package commands

import (
	"log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func WeatherCommand(session *disgolf.Bot) {
	session.Router.Register(&disgolf.Command{
		Name:        "weather",
		Description: "get weather forecast",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:         "city",
				Description:  "city you want forecast for",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
			},
		},
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						&discordgo.MessageEmbed{
							Title:       "WeatherCommand",
							Description: "this embedded message provides weather info",
							Color:       0xff00a4,
							Type:        discordgo.EmbedTypeRich,
							Author: &discordgo.MessageEmbedAuthor{
								Name: ctx.Interaction.Member.User.Username + " (" + ctx.Interaction.Member.Nick + ")",
							},
						},
					},
				},
			})
		}),
		MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
			_, _ = ctx.Reply("It worked", true)
		}),
		MessageMiddlewares: []disgolf.MessageHandler{
			disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				log.Println("Everything works")
				ctx.Next()
			}),
		},
		Middlewares: []disgolf.Handler{
			disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				log.Println("WeatherCommand was invoked")
				ctx.Next()
			}),
		},
	})
	session.AddHandler(session.Router.HandleInteraction)
	session.AddHandler(session.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
		Prefixes:      []string{"d.", "dis.", "disgolf."},
		MentionPrefix: true,
	}))
}
