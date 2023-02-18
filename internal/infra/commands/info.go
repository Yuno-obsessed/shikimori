package commands

import (
	"fmt"
	"strconv"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func PresenceInfo(session *disgolf.Bot, ctx *disgolf.Ctx, user *discordgo.User) (string, int) {
	pr, err := session.State.Presence(ctx.Interaction.GuildID, user.ID)
	if pr == nil {
		return ":white_circle: Offline", 0xffffff
	}
	if err != nil {
		fmt.Println(err)
	}
	var color int
	presence := string(pr.Status)
	var userPresence string
	switch presence {
	case "dnd":
		userPresence = ":red_circle: Do not disturb"
		color = 0xe60025
		break
	case "idle":
		userPresence = ":yellow_circle: Idle"
		color = 0xe6dd00
		break
	case "online":
		userPresence = ":green_circle: Online"
		color = 0x00b340
		break
	}
	return userPresence, color
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
					guild, _ := session.Guild(ctx.Interaction.GuildID)
					usernameId := guild.OwnerID
					guildCounts, err := session.GuildWithCounts(ctx.Interaction.GuildID)
					y := strconv.Itoa(guildCounts.ApproximatePresenceCount) + "/" + strconv.Itoa(guildCounts.ApproximateMemberCount)
					if err != nil {
						fmt.Println(err)
					}
					gtimestamp, err := discordgo.SnowflakeTimestamp(ctx.Interaction.GuildID)
					username, err := ctx.User(usernameId)
					// if err != nil {
					// 	fmt.Println(err)
					// }
					_ = ctx.Respond(&discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Embeds: []*discordgo.MessageEmbed{
								{
									Color: 0xb1cc00,
									Title: "Info about guild: ",
									Type:  discordgo.EmbedTypeRich,
									Thumbnail: &discordgo.MessageEmbedThumbnail{
										URL: guild.IconURL(),
									},
									Fields: []*discordgo.MessageEmbedField{
										{
											Name:   "Guild name:",
											Value:  guild.Name,
											Inline: true,
										},
										{
											Name:   "Guild's owner: ",
											Value:  username.Username,
											Inline: true,
										},
										{
											Name:   "Created: ",
											Value:  gtimestamp.Format("02 Jan 2006 15:04:05"),
											Inline: false,
										},
										{
											Name:   "Current Member Count: ",
											Value:  y,
											Inline: true,
										},
										{
											Name:   "Emoji count",
											Value:  strconv.Itoa(len(guild.Emojis)),
											Inline: true,
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
					var nickname string
					presence, color := PresenceInfo(session, ctx, userStruct)
					timestamp, err := discordgo.SnowflakeTimestamp(userStruct.ID)
					if err != nil {
						fmt.Println(err)
					}
					memberStruct, _ := session.GuildMember(ctx.Interaction.GuildID, userStruct.ID)
					if memberStruct.Nick != "" {
						nickname = memberStruct.Nick
					} else {
						nickname = userStruct.Username
					}
					if nickname != "" {
						_ = ctx.Respond(&discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Embeds: []*discordgo.MessageEmbed{
									{
										Title: "Info about " + userStruct.Username,
										Color: color,
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
												Name:   "AKA: ",
												Value:  nickname,
												Inline: true,
											},

											{
												Name:   "Status: ",
												Value:  presence,
												Inline: false,
											},
											{
												Name:   "Joined server: ",
												Value:  ctx.Interaction.Member.JoinedAt.Format("02 Jan 2006 15:04:05"),
												Inline: true,
											},

											{
												Name:   "Joined discord: ",
												Value:  timestamp.Format("02 Jan 2006 15:04:05"),
												Inline: true,
											},
										},
									},
								},
							},
						})
					}
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
