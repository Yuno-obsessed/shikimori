package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"github.com/yuno-obsessed/shikimori/internal/logs"
)

func ChadAdvice() string {
	resp, err := http.Get("http://fucking-great-advice.ru/api/random")
	if err != nil {
		fmt.Println("No response")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	type response struct {
		Id    int    `json:"id"`
		Text  string `json:"text"`
		Sound string `json:"sound"`
	}
	var data response
	err = json.Unmarshal(body, &data)
	if err != nil {
		logs.LogErr(logs.ErrUnmarshalingJSON, "ChadAdvice")
	}
	return data.Text
}

func NormalAdvice() string {
	resp, err := http.Get("https://api.adviceslip.com/advice")
	if err != nil {
		fmt.Println("No response")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	type response struct {
		Slip struct {
			ID     int    `json:"id,omitempty"`
			Advice string `json:"advice,omitempty"`
		} `json:"slip,omitempty"`
	}
	var data response
	err = json.Unmarshal(body, &data)
	if err != nil {
		logs.LogErr(logs.ErrUnmarshalingJSON, "NormalAdvice")
	}
	return data.Slip.Advice
}

func AdviceCommand(session *disgolf.Bot) {
	session.Router.Register(&disgolf.Command{
		Name:        "advice",
		Description: "generate great-advice",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "advice-type",
				Description: "choose the advice you want",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "chad-advice",
						Value: "give chad-advice",
					},
					{
						Name:  "normal-advice",
						Value: "give normal-advice",
					},
				},
			},
		},
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			var author string
			if ctx.Interaction.Member.Nick != "" {
				author = "Advice for " + ctx.Interaction.Member.User.Username + " (" + ctx.Interaction.Member.Nick + ")"
			} else {
				author = "Advice for " + ctx.Interaction.Member.User.Username
			}
			if ctx.Interaction.ApplicationCommandData().Options[0].StringValue() == "give chad-advice" {

				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Color: 0x1e201e,
								Type:  discordgo.EmbedTypeRich,
								Author: &discordgo.MessageEmbedAuthor{
									Name: author,
								},
								Thumbnail: &discordgo.MessageEmbedThumbnail{
									URL: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi.pinimg.com%2Foriginals%2F25%2Fbd%2F8b%2F25bd8b7f6e57cdfd17747b25d753b2ce.jpg&f=1&nofb=1&ipt=d651d0f9582515894007dfede6e35cfa9796c137fbec24eee754f0a4161700f5&ipo=images",
								},

								Fields: []*discordgo.MessageEmbedField{
									{
										Name:  ChadAdvice(),
										Value: "*Chad-advice*",
									},
								},
							},
						},
					},
				})

			} else if ctx.Interaction.ApplicationCommandData().Options[0].StringValue() == "give normal-advice" {
				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Color: 0x00ff5c,
								Type:  discordgo.EmbedTypeRich,
								Author: &discordgo.MessageEmbedAuthor{
									Name: author,
								},
								Fields: []*discordgo.MessageEmbedField{
									{
										Name:  NormalAdvice(),
										Value: "*Normal-advice*",
									},
								},
								Thumbnail: &discordgo.MessageEmbedThumbnail{
									URL: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fimages.flickreel.com%2Fwp-content%2Fuploads%2F2016%2F06%2FYoda_SWSB.png&f=1&nofb=1&ipt=7602823f2df62016fb3afc04db265f8efd8e0a3fef694d6b2b0216c1e5843425&ipo=images",
								},
							},
						},
					},
				})
			}
		}),
		MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
			_, _ = ctx.Reply("advice", false)

		}),
		MessageMiddlewares: []disgolf.MessageHandler{
			disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				fmt.Println("ez")
				ctx.Next()
			}),
		},
		Middlewares: []disgolf.Handler{
			disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				fmt.Println("AdviceCommand was invoked")
				ctx.Next()
			}),
		},
	})
	session.AddHandler(session.Router.HandleInteraction)
	session.AddHandler(session.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
		Prefixes:      []string{"d.", "dis.", "disgold."},
		MentionPrefix: true,
	}))

}
