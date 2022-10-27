package commands

import (
	"fmt"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func AvatarCommand(session *disgolf.Bot) {
	session.Router.Register(&disgolf.Command{
		Name:        "avatar",
		Description: "returns avatar",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "user you want the avatar of",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    false,
			},
		},
		Type: discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			if ctx.Interaction.ApplicationCommandData().Options == nil {
				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "https://cdn.discordapp.com/avatars/" + ctx.Interaction.Member.User.ID + "/" + ctx.Interaction.Member.User.Avatar + ".png?size=1024",
					},
				})
			} else {
				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "https://cdn.discordapp.com/avatars/" + ctx.Interaction.ApplicationCommandData().Options[0].UserValue(ctx.Session).ID + "/" + ctx.Interaction.ApplicationCommandData().Options[0].UserValue(ctx.Session).Avatar + ".png?size=1024",
					},
				})
			}
		}),

		MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
			_, _ = ctx.Reply("hi", true)
		}),
		MessageMiddlewares: []disgolf.MessageHandler{
			disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				fmt.Println("Message Middleware worked")
				ctx.Next()
			}),
		},
		Middlewares: []disgolf.Handler{
			disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				fmt.Println("AvatarCommand was invoked")
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
