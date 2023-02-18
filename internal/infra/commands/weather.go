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
	"github.com/yuno-obsessed/shikimori/internal/logs"
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
		WindSpeed     string `json:"windspeedKmph"`
		WindDirection string `json:"winddir16Point"`
	} `json:"current_condition,omitempty"`
	NearestArea []struct {
		Country []struct {
			Value string `json:"value,omitempty"`
		} `json:"country,omitempty"`
	} `json:"nearest_area,omitempty"`
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

func countryInfo(country string) string {
	response, err := http.Get("https://countryflagapi.herokuapp.com/country/" + country)
	if err != nil {
		log.Println(logs.ErrReachingAPI, "countryInfo")
	}
	type CountryInfo []struct {
		Name    string `json:"name,omitempty"`
		Code    string `json:"code,omitempty"`
		Capital string `json:"capital,omitempty"`
		Flag    string `json:"flag,omitempty"`
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var countryData CountryInfo
	err = json.Unmarshal(body, &countryData)
	if err != nil {
		log.Println(err)
	}
	return countryData[0].Code
}
func weatherImage(city string) string {
	var link string
	c := colly.NewCollector()
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link = "https://wttr.in/" + e.Attr("src")
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
				Name:        "type",
				Description: "type of a weather info",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{Name: "brief",
						Value: "briefWeather",
					},
					{
						Name:  "detailed",
						Value: "detailedWeather",
					},
				},
				Required: true,
			},
			{
				Name:        "city",
				Description: "city you want weather of",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			city := ctx.Interaction.ApplicationCommandData().Options[1].StringValue()
			weather := weatherInfo(city)
			if ctx.Interaction.ApplicationCommandData().Options[0].StringValue() == "briefWeather" {

				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Title: "Brief weather in " + city,
								Color: 0xff00a4,
								Type:  discordgo.EmbedTypeRich,
								Image: &discordgo.MessageEmbedImage{
									URL: weatherImage(city),
								},
								Thumbnail: &discordgo.MessageEmbedThumbnail{
									URL: "https://www.countryflagicons.com/FLAT/64/" + countryInfo(weatherInfo(city).NearestArea[0].Country[0].Value) + ".png",
								},
								Fields: []*discordgo.MessageEmbedField{
									{
										Name:   "Description:",
										Value:  weather.CurrentCondition[0].WeatherDesc[0].Value,
										Inline: false,
									},
									{
										Name:   "Feels Like:",
										Value:  weather.CurrentCondition[0].FeelsLikeC,
										Inline: true,
									},
									{
										Name:   "Temperature:",
										Value:  weather.CurrentCondition[0].TempC,
										Inline: true,
									},
									{
										Name:   "Humidity:",
										Value:  weather.CurrentCondition[0].Humidity,
										Inline: true,
									},
									{
										Name:   "Date:",
										Value:  weather.CurrentCondition[0].LocalObsDateTime,
										Inline: false,
									},
								},
							},
						},
					},
				})
			} else {
				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Title: "Detailed weather in " + city,
								Color: 0xff00a4,
								Type:  discordgo.EmbedTypeRich,
								Thumbnail: &discordgo.MessageEmbedThumbnail{
									URL: "https://www.countryflagicons.com/FLAT/64/" + countryInfo(weatherInfo(city).NearestArea[0].Country[0].Value) + ".png",
								},
								Image: &discordgo.MessageEmbedImage{
									URL: weatherImage(city),
								},
								Fields: []*discordgo.MessageEmbedField{
									{
										Name:   "Description:",
										Value:  weather.CurrentCondition[0].WeatherDesc[0].Value,
										Inline: false,
									},
									{
										Name:   "Feels Like:",
										Value:  weather.CurrentCondition[0].FeelsLikeC,
										Inline: true,
									},
									{
										Name:   "Temperature:",
										Value:  weather.CurrentCondition[0].TempC,
										Inline: true,
									},
									{
										Name:   "Max Temperature:",
										Value:  weather.Weather[0].MaxTempC,
										Inline: true,
									},
									{
										Name:   "Humidity:",
										Value:  weather.CurrentCondition[0].Humidity,
										Inline: true,
									},
									{
										Name:   "Wind Speed:",
										Value:  weather.CurrentCondition[0].WindSpeed,
										Inline: true,
									},
									{
										Name:   "Date:",
										Value:  weather.CurrentCondition[0].ObservationTime,
										Inline: true,
									},
									{
										Name:   "Sunrise:",
										Value:  weather.Weather[0].Astronomy[0].Sunrise,
										Inline: true,
									},
									{
										Name:   "Sunset:",
										Value:  weather.Weather[0].Astronomy[0].Sunset,
										Inline: true,
									},
									{
										Name:   "Moon Phase:",
										Value:  weather.Weather[0].Astronomy[0].MoonPhase,
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
	})
	session.AddHandler(session.Router.HandleInteraction)
	session.AddHandler(session.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
		Prefixes:      []string{"d.", "dis.", "disgolf."},
		MentionPrefix: true,
	}))
}
