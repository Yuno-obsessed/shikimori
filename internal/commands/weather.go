package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

type Weather struct {
	CurrentCondition []struct {
		FeelsLikeC       string `json:"FeelsLikeC"`
		Cloudcover       string `json:"cloudcover,omitempty"`
		Humidity         string `json:"humidity"`
		LocalObsDateTime string `json:"localObsDateTime"`
		ObservationTime  string `json:"observation_time"`
		Pressure         string `json:"pressure"`
		TempC            string `json:"temp_C"`
		WeatherDesc      []struct {
			Value string `json:"value,omitempty"`
		}
	} `json:"current_condition,omitempty"`
	Weather []struct {
		Astronomy []struct {
			MoonIllumination string `json:"moon_illumination,omitempty"`
			MoonPhase        string `json:"moon_phase,omitempty"`
			Moonrise         string `json:"moonrise,omitempty"`
			Moonset          string `json:"moonset,omitempty"`
			Sunrise          string `json:"sunrise,omitempty"`
			Sunset           string `json:"sunset,omitempty"`
		} `json:"astronomy,omitempty"`
		AvgTempC string `json:"avgtempC,omitempty"`
		Date     string `json:"date,omitempty"`
		Hourly   []struct {
			WindGustKmph string `json:"WindGustKmph,omitempty"`
		}
		MaxTempC    string `json:"maxtempC"`
		MinTempC    string `json:"mintempC"`
		SunHour     string `json:"sunHour,omitempty"`
		TotalSnowCm string `json:"totalSnow_cm,omitempty"`
	} `json:"weather,omitempty"`
}

func weatherInfo(city string) *Weather {
	response, err := http.Get("https://wttr.in/" + city + "?format=j1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	type responseData *Weather
	var data responseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// func weatherGet(weatherInfo *Weather) string {
//
// }

func weatherImage(city string) string {
	var link string
	c := colly.NewCollector()
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link = "https://wttr.in/" + e.Attr("src")
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://wttr.in/" + city + "?format=v2")
	return link
}

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
			var nickname string
			city := ctx.Interaction.ApplicationCommandData().Options[0].StringValue()
			weather := weatherInfo(city)
			if ctx.Interaction.Member.Nick != "" {
				nickname = " (" + ctx.Interaction.Member.Nick + ")"
			} else {
				nickname = ""
			}
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "WeatherCommand",
							Description: "this embedded message provides weather info",
							Color:       0xff00a4,
							Type:        discordgo.EmbedTypeRich,
							Author: &discordgo.MessageEmbedAuthor{
								Name: ctx.Interaction.Member.User.Username + nickname,
							},
							Image: &discordgo.MessageEmbedImage{
								URL: weatherImage(city),
							},
							Fields: []*discordgo.MessageEmbedField{
								{
									Name:  "FeelsLike:",
									Value: weather.CurrentCondition[0].FeelsLikeC,
								},
								{
									Name:  "Temp:",
									Value: weather.CurrentCondition[0].TempC,
								},
								{
									Name:  "MaxInfo",
									Value: weather.Weather[0].MaxTempC,
								},
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
