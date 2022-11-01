package commands

import (
	"fmt"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"github.com/yuno-obsessed/shikimori/internal/logs"
)

func PresenceInfo(session *disgolf.Bot, ctx *disgolf.Ctx) string {
	presence, err := ctx.State.Presence(getGuildStructure(session, ctx).ID, ctx.Interaction.Member.User.ID)
	if err != nil {
		fmt.Println(err)
	}
	userPresence := string(presence.Status)
	if userPresence == "" {
		return "yESSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSss"
	}
	return userPresence
}
func getGuildStructure(session *disgolf.Bot, ctx *disgolf.Ctx) *discordgo.Guild {
	guild, _ := session.Guild(ctx.Interaction.GuildID)
	return guild
}

func getUserStructure(session *disgolf.Bot, ctx *disgolf.Ctx) *discordgo.User {
	user, err := session.User(ctx.Interaction.ApplicationCommandData().Options[0].UserValue(ctx.Session).ID)
	if err != nil {
		logs.LogErr(logs.ErrUserNotExist, "getUserStructure")
	}
	return user
}

func InfoCommand(session *disgolf.Bot) {
	session.Router.Register(&disgolf.Command{
		Name:        "info",
		Description: "show info about specific thing",
		Type:        discordgo.ChatApplicationCommand,

		MessageMiddlewares: []disgolf.MessageHandler{
			disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				fmt.Println("Everything works")
				ctx.Next()
			}),
		},
		Middlewares: []disgolf.Handler{
			disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				fmt.Println("InfoCommand was invoked")
				ctx.Next()
			}),
		},
		SubCommands: disgolf.NewRouter([]*disgolf.Command{
			{
				Name:        "guild",
				Description: "guild info",
				Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
					_ = ctx.Respond(&discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Embeds: []*discordgo.MessageEmbed{
								{
									Color: 0xb1cc00,
									Title: "Guild info command",
									Type:  discordgo.EmbedTypeRich,
									Thumbnail: &discordgo.MessageEmbedThumbnail{
										URL: getGuildStructure(session, ctx).IconURL(),
									},
									Author: &discordgo.MessageEmbedAuthor{
										Name: ctx.Interaction.Member.User.Username,
									},
									Fields: []*discordgo.MessageEmbedField{
										{
											Name:  "Guild name",
											Value: getGuildStructure(session, ctx).Name,
										},
									},
								},
							},
						},
					})
				}),
				MessageMiddlewares: []disgolf.MessageHandler{
					disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
						fmt.Println("guild middleware")
						ctx.Next()
					}),
				},
				Middlewares: []disgolf.Handler{
					disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						fmt.Println("guild middleware")
						ctx.Next()
					}),
				},
				MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
					_, _ = ctx.Reply("that works", false)
				}),
			},
			{
				Name:        "user",
				Description: "user info",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "username",
						Description: "user you want info about",
						Type:        discordgo.ApplicationCommandOptionUser,
						Required:    true,
					},
				},
				Type: discordgo.ChatApplicationCommand,
				Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
					userStruct := ctx.Interaction.ApplicationCommandData().Options[0].Options[0].UserValue(ctx.Session)
					nickname := userStruct.Username
					if ctx.Interaction.Member.Nick != "" {
						nickname = ctx.Interaction.Member.Nick
					}
					_ = ctx.Respond(&discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Embeds: []*discordgo.MessageEmbed{
								{
									Title: "Info about " + userStruct.Username,
									Color: 0xcd0101,
									Type:  discordgo.EmbedTypeRich,
									Thumbnail: &discordgo.MessageEmbedThumbnail{
										URL: "https://cdn.discordapp.com/avatars/" + userStruct.ID + "/" + userStruct.Avatar + ".png?size=1024",
									},

									Fields: []*discordgo.MessageEmbedField{
										{
											Name:   "Username: ",
											Value:  userStruct.Username,
											Inline: true,
										},
										{
											Name:   "Nickname: ",
											Value:  nickname,
											Inline: true,
										},
										{
											Name:   "Joined at: ",
											Value:  ctx.Interaction.Member.JoinedAt.Format("02 Jan 2006 15:04:05"),
											Inline: false,
										},

										{
											Name:   "Status: ",
											Value:  PresenceInfo(session, ctx),
											Inline: true,
										},
										// 										{
										// Name: "Joined discord",
										// 											Value: ctx.,
										// 										},
									},
								},
							},
						},
					})
				}),
				MessageMiddlewares: []disgolf.MessageHandler{
					disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
						fmt.Println("user message middleware")
						ctx.Next()
					}),
				},
				Middlewares: []disgolf.Handler{
					disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						fmt.Println("user middleware")
						ctx.Next()
					}),
				},
				MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
					_, _ = ctx.Reply("that worked", false)
				}),
			},
		}),
	})
	session.AddHandler(session.Router.HandleInteraction)
	session.AddHandler(session.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
		Prefixes:      []string{"d.", "dis.", "disgolf."},
		MentionPrefix: true,
	}))
}
