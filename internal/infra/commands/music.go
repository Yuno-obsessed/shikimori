package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

type Music struct {
	Tracks struct {
		Hits []struct {
			Track struct {
				Layout   string `json:"layout"`
				Type     string `json:"type"`
				Key      string `json:"key"`
				Title    string `json:"title"`
				Subtitle string `json:"subtitle"`
				Share    struct {
					Subject  string `json:"subject"`
					Text     string `json:"text"`
					Href     string `json:"href"`
					Image    string `json:"image"`
					Twitter  string `json:"twitter"`
					HTML     string `json:"html"`
					Avatar   string `json:"avatar"`
					Snapchat string `json:"snapchat"`
				} `json:"share"`
				Images struct {
					Background string `json:"background"`
					Coverart   string `json:"coverart"`
					Coverarthq string `json:"coverarthq"`
					Joecolor   string `json:"joecolor"`
				} `json:"images"`
				Hub struct {
					Type    string `json:"type"`
					Image   string `json:"image"`
					Actions []struct {
						Name string `json:"name"`
						Type string `json:"type"`
						ID   string `json:"id,omitempty"`
						URI  string `json:"uri,omitempty"`
					} `json:"actions"`
					Options []struct {
						Caption string `json:"caption"`
						Actions []struct {
							Name string `json:"name"`
							Type string `json:"type"`
							URI  string `json:"uri"`
						} `json:"actions"`
						Beacondata struct {
							Type         string `json:"type"`
							Providername string `json:"providername"`
						} `json:"beacondata"`
						Image               string `json:"image"`
						Type                string `json:"type"`
						Listcaption         string `json:"listcaption"`
						Overflowimage       string `json:"overflowimage"`
						Colouroverflowimage bool   `json:"colouroverflowimage"`
						Providername        string `json:"providername"`
					} `json:"options"`
					Providers []struct {
						Caption string `json:"caption"`
						Images  struct {
							Overflow string `json:"overflow"`
							Default  string `json:"default"`
						} `json:"images"`
						Actions []struct {
							Name string `json:"name"`
							Type string `json:"type"`
							URI  string `json:"uri"`
						} `json:"actions"`
						Type string `json:"type"`
					} `json:"providers"`
					Explicit    bool   `json:"explicit"`
					Displayname string `json:"displayname"`
				} `json:"hub"`
				Artists []struct {
					ID     string `json:"id"`
					Adamid string `json:"adamid"`
				} `json:"artists"`
				URL string `json:"url"`
			} `json:"track"`
			Snippet string `json:"snippet,omitempty"`
		} `json:"hits"`
	} `json:"tracks"`
	Artists struct {
		Hits []struct {
			Artist struct {
				Avatar   string `json:"avatar"`
				ID       string `json:"id"`
				Name     string `json:"name"`
				Verified bool   `json:"verified"`
				Weburl   string `json:"weburl"`
				Adamid   string `json:"adamid"`
			} `json:"artist,omitempty"`
		} `json:"hits"`
	} `json:"artists"`
}

func MusicInfo(name string) *Music {
	url := "https://shazam.p.rapidapi.com/search?term=" + name + "&locale=en-US&offset=0&limit=5"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-RapidAPI-Key", "92fefacf56mshf9ff7e2db031aa3p1e4724jsn29b4bed240b4")
	req.Header.Add("X-RapidAPI-Host", "shazam.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	type responseData *Music
	var MusicData responseData
	err = json.Unmarshal(body, &MusicData)
	if err != nil {
		fmt.Println(err)
	}
	return MusicData
}

func MusicCommand(session *disgolf.Bot) {
	// session.Router.Register(&disgolf.Command{
	// 	Name:        "music",
	// 	Description: "some music action",
	// 	Type:        discordgo.ChatApplicationCommand,
	// 	Options: []*discordgo.ApplicationCommandOption{
	// 		{
	// 			Name:        "action",
	// 			Description: "what to do with music",
	// 			Type:        discordgo.ApplicationCommandOptionString,
	// 			Choices: []*discordgo.ApplicationCommandOptionChoice{
	// 				{
	// 					Name:  "info",
	// 					Value: "musicInfo",
	// 				},
	// 				{
	// 					Name:  "link",
	// 					Value: "musicLink",
	// 				},
	// 			},
	// 			Required: true,
	// 		},
	// 		{
	// 			Name:        "songName",
	// 			Description: "the songs title",
	// 			Type:        discordgo.ApplicationCommandOptionString,
	// 			Required:    true,
	// 		},
	// 	},
	// 	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
	// 		// songName := ctx.Interaction.ApplicationCommandData().Options[1].StringValue()
	// 		if ctx.Interaction.ApplicationCommandData().Options[0].StringValue() == "musicInfo" {
	// 			_ = ctx.Respond(&discordgo.InteractionResponse{
	// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 				Data: &discordgo.InteractionResponseData{
	// 					Embeds: []*discordgo.MessageEmbed{
	// 						{
	// 							Color: 0x4aff00,
	// 							Type:  discordgo.EmbedTypeRich,
	// 							Title: "Music info",
	// 							Fields: []*discordgo.MessageEmbedField{
	// 								{
	// 									Name: "Title",
	// 									// Value:  MusicInfo(songName).Tracks.Hits[0].Track.Title,
	// 									Value:  "yes",
	// 									Inline: true,
	// 								},
	// 								{
	// 									Name: "Artist",
	// 									// Value:  MusicInfo(songName).Tracks.Hits[0].Track.Subtitle,
	// 									Value:  "no",
	// 									Inline: true,
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			})
	// 		} else {
	// 			_ = ctx.Respond(&discordgo.InteractionResponse{
	// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 				Data: &discordgo.InteractionResponseData{
	// 					Embeds: []*discordgo.MessageEmbed{
	// 						{
	// 							Color: 0x4aff00,
	// 							Type:  discordgo.EmbedTypeRich,
	// 							Title: "Music info",
	// 							Fields: []*discordgo.MessageEmbedField{
	// 								{
	// 									Name: "Title",
	// 									// Value:  MusicInfo(songName).Tracks.Hits[0].Track.Title,
	// 									Value:  "yes",
	// 									Inline: true,
	// 								},
	// 								{
	// 									Name: "Artist",
	// 									// Value:  MusicInfo(songName).Tracks.Hits[0].Track.URL,
	// 									Value:  "no",
	// 									Inline: true,
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			})
	// 		}
	// 	}),
	session.Router.Register(&disgolf.Command{
		Name:        "music",
		Description: "get music info",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "action",
				Description: "type of action",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "info",
						Value: "music info",
					},
					{
						Name:  "links",
						Value: "music links",
					},
				},
				Required: true,
			},
			{
				Name:        "name",
				Description: "the name of a song",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			songName := ctx.Interaction.ApplicationCommandData().Options[1].StringValue()
			song := MusicInfo(songName)
			if ctx.Interaction.ApplicationCommandData().Options[0].StringValue() == "music info" {

				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Color: 0xff00a4,
								Type:  discordgo.EmbedTypeRich,
								Fields: []*discordgo.MessageEmbedField{
									{
										Name:   "Title:",
										Value:  song.Tracks.Hits[0].Track.Title,
										Inline: true,
									},
									{
										Name:   "Author:",
										Value:  song.Tracks.Hits[0].Track.Subtitle,
										Inline: true,
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
								Color: 0xff00a4,
								Type:  discordgo.EmbedTypeRich,
								Fields: []*discordgo.MessageEmbedField{
									{
										Name:   "Title:",
										Value:  song.Tracks.Hits[0].Track.Title,
										Inline: true,
									},
									{
										Name:   "Links:",
										Value:  song.Tracks.Hits[0].Track.URL,
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
