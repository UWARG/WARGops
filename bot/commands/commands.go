package commands

import (
	"strings"

	"bitbucket.org/mikehouston/asana-go"
	"github.com/bwmarrin/discordgo"
	"github.com/bwmarrin/lit"
)

const (
	// Version is a constant that stores the WARGops version information.
	Version       = "v0.0.1"
	ephemeralFlag = 64
)

var (
	LeadRoleID  string
	WorkspaceID string
	AsanaClient *asana.Client

	ProjectMap   = make(map[string]*asana.Project)
	Projects     []*asana.Project
	ProjectNames []string

	UserMap   = make(map[string]*asana.User)
	Users     []*asana.User
	UserNames []string

	Commands = make(map[string]*Command)
)

type Command struct {
	*discordgo.ApplicationCommand
	Handler          func(*discordgo.Session, *discordgo.InteractionCreate) (*discordgo.InteractionResponseData, error)
	Autocomplete     func(*discordgo.Session, *discordgo.InteractionCreate) ([]*discordgo.ApplicationCommandOptionChoice, error)
	MessageComponent func(*discordgo.Session, *discordgo.InteractionCreate, []string) (*discordgo.InteractionResponseData, error)
}

func OnAutocomplete(ds *discordgo.Session, ic *discordgo.InteractionCreate) {
	if ic.Type != discordgo.InteractionApplicationCommandAutocomplete {
		return
	}

	data := ic.ApplicationCommandData()
	cmd, ok := Commands[data.Name]
	if !ok || cmd.Autocomplete == nil {
		return
	}

	choices, err := cmd.Autocomplete(ds, ic)
	if err != nil {
		return
	}
	if len(choices) > 25 {
		choices = choices[:25]
	}

	err = ds.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})
	if err != nil {
		lit.Error("responding to autocomplete: %v", err)
	}
}

func OnInteractionCommand(ds *discordgo.Session, ic *discordgo.InteractionCreate) {
	if ic.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := ic.ApplicationCommandData()
	cmd, ok := Commands[data.Name]
	if !ok {
		return
	}

	res, err := cmd.Handler(ds, ic)
	if err != nil {
		res = &discordgo.InteractionResponseData{
			Flags: ephemeralFlag,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Error",
					Description: err.Error(),
					Color:       0xEE2211,
				},
			},
		}
	}

	typ := discordgo.InteractionResponseChannelMessageWithSource
	if res.Title != "" {
		typ = discordgo.InteractionResponseModal
	}
	err = ds.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
		Type: typ,
		Data: res,
	})
	if err != nil {
		lit.Error("responding to interaction %s: %v", data.Name, err)
	}
}

func OnModalSubmit(ds *discordgo.Session, ic *discordgo.InteractionCreate) {
	if ic.Type != discordgo.InteractionModalSubmit {
		return
	}

	data := ic.ModalSubmitData()
	cmd := Commands["asana"]

	res, err := cmd.Handler(ds, ic)
	if err != nil {
		res = &discordgo.InteractionResponseData{
			Flags: ephemeralFlag,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Error",
					Description: err.Error(),
					Color:       0xEE2211,
				},
			},
		}
	}
	if res == nil {
		err = ds.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		})
	} else {
		err = ds.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: res,
		})
	}

	if err != nil {
		lit.Error("responding to modal submit %s: %v", data.CustomID, err)
	}
}

func OnMessageComponent(ds *discordgo.Session, ic *discordgo.InteractionCreate) {
	if ic.Type != discordgo.InteractionMessageComponent {
		return
	}

	data := ic.MessageComponentData()
	parts := strings.Split(data.CustomID, ",")
	if len(parts) < 2 {
		return
	}
	cmd := Commands[parts[0]]

	res, err := cmd.MessageComponent(ds, ic, parts)
	if err != nil {
		res = &discordgo.InteractionResponseData{
			Flags: ephemeralFlag,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Error",
					Description: err.Error(),
					Color:       0xEE2211,
				},
			},
		}
	}
	if res == nil {
		err = ds.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredMessageUpdate,
		})
		if err != nil {
			lit.Error("responding to interaction %s: %v", parts, err)
		}
		return
	}

	typ := discordgo.InteractionResponseChannelMessageWithSource
	if res.Title != "" {
		typ = discordgo.InteractionResponseModal
	}
	err = ds.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
		Type: typ,
		Data: res,
	})
	if err != nil {
		lit.Error("responding to interaction %s: %v", parts, err)
	}
}

func ContentResponse(c string) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Content: c,
	}
}

func EphemeralResponse(c string) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Flags:   ephemeralFlag,
		Content: c,
	}
}

func EmbedResponse(e discordgo.MessageEmbed) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{&e},
	}
}

func Autocomplete(options ...string) []*discordgo.ApplicationCommandOptionChoice {
	var choices []*discordgo.ApplicationCommandOptionChoice
	for _, opt := range options {
		if len(opt) > 100 {
			opt = opt[:100]
		}

		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  opt,
			Value: opt,
		})
	}
	return choices
}
