package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func NiceAdvice() string {
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
	json.Unmarshal(body, &data)
	return data.Text
}

func AdviceCommand(session *disgolf.Bot) {
	session.Router.Register(&disgolf.Command{
		Name:        "advice",
		Description: "generate great-advice",
		Type:        discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: NiceAdvice(),
				},
			})
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
