package commands

import (
	"fmt"
	"strings"
	"time"

	"bitbucket.org/mikehouston/asana-go"
	"github.com/bwmarrin/discordgo"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

func init() {
	Commands[cmdAsana.Name] = &Command{
		ApplicationCommand: cmdAsana,
		Autocomplete:       handleAsanaAutocomplete,
		Handler:            handleAsanaRaw,
		MessageComponent:   handleAsanaComponent,
	}
}

var cmdAsana = &discordgo.ApplicationCommand{
	Type:        discordgo.ChatApplicationCommand,
	Name:        "asana",
	Description: "WARG Asana Integration",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "create",
			Description: "Create an Asana task and assign it to a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "project",
					Description:  "Select the Project for the task",
					Required:     true,
					Autocomplete: true,
				},
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "assignee",
					Description:  "Select the Assignee for the task (defaults to yourself)",
					Required:     true,
					Autocomplete: true,
				},
			},
		},
	},
}

func handleAsanaRaw(ds *discordgo.Session, ic *discordgo.InteractionCreate) (*discordgo.InteractionResponseData, error) {
	var lead bool
	for _, role := range ic.Member.Roles {
		if role == LeadRoleID {
			lead = true
			break
		}
	}

	if !lead {
		return nil, fmt.Errorf("only team leads can use this command for now")
	}

	switch ic.Type {
	case discordgo.InteractionApplicationCommand:
		return handleAsanaCmd(ds, ic)
	case discordgo.InteractionModalSubmit:
		return handleAsanaSubmit(ds, ic)
	}

	return nil, nil
}

func handleAsanaAutocomplete(ds *discordgo.Session, ic *discordgo.InteractionCreate) ([]*discordgo.ApplicationCommandOptionChoice, error) {
	options := ic.ApplicationCommandData().Options[0].Options
	active := options[0]
	if len(options) == 2 {
		if options[1].Focused {
			active = options[1]
		}
	}
	val := active.StringValue()
	var results []*discordgo.ApplicationCommandOptionChoice
	switch active.Name {
	case "project":
		ranks := fuzzy.RankFindNormalizedFold(val, ProjectNames)
		for _, r := range ranks {
			results = append(results, &discordgo.ApplicationCommandOptionChoice{
				Name:  r.Target,
				Value: Projects[r.OriginalIndex].ID,
			})
		}
	case "assignee":
		if val == "" {
			name := ic.Member.Nick
			if i := strings.Index(name, "["); i != -1 {
				name = strings.TrimSpace(name[:i])
			}
			if l := strings.Index(name, "("); l != -1 {
				if r := strings.Index(name, ")"); r != -1 && r > l {
					name = strings.TrimSpace(name[:l]) + " " + strings.TrimSpace(name[:r])
				}
			}

			ranks := fuzzy.RankFindNormalizedFold(name, UserNames)
			if len(ranks) > 0 {
				results = []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  ranks[0].Target,
						Value: Users[ranks[0].OriginalIndex].ID,
					},
				}
			}
			break
		}
		ranks := fuzzy.RankFindNormalizedFold(val, UserNames)
		for _, r := range ranks {
			results = append(results, &discordgo.ApplicationCommandOptionChoice{
				Name:  r.Target,
				Value: Users[r.OriginalIndex].ID,
			})
		}
	}
	return results, nil
}

func handleAsanaCmd(ds *discordgo.Session, ic *discordgo.InteractionCreate) (*discordgo.InteractionResponseData, error) {
	options := ic.ApplicationCommandData().Options[0].Options
	project, ok := ProjectMap[options[0].StringValue()]
	if !ok {
		return nil, fmt.Errorf("invalid project provided")
	}
	assignee, ok := UserMap[options[1].StringValue()]
	if !ok {
		return nil, fmt.Errorf("invalid assignee provided")
	}

	return &discordgo.InteractionResponseData{
		Title:    "Task for " + project.Name,
		CustomID: project.ID + "," + assignee.ID,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.TextInput{
						Label:       "Task Title",
						Style:       discordgo.TextInputShort,
						CustomID:    "title",
						Value:       "",
						Required:    true,
						Placeholder: "Try Pineapple Pizza",
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					&discordgo.TextInput{
						Label:       "Task Description",
						Style:       discordgo.TextInputParagraph,
						CustomID:    "desc",
						Value:       "",
						Placeholder: "Praise Pineapple Pizza",
					},
				},
			},
		},
	}, nil
}

func handleAsanaUpdate(ds *discordgo.Session, ic *discordgo.InteractionCreate) {
	data := ic.ModalSubmitData()
	args := strings.Split(data.CustomID, ",")

	task := &asana.Task{ID: args[2]}
	if err := task.Fetch(AsanaClient); err != nil {
		ds.FollowupMessageCreate(ic.Interaction, false, &discordgo.WebhookParams{
			Flags: 64,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Error",
					Description: err.Error(),
					Color:       0xEE2211,
				},
			},
		})
		return
	}

	switch args[1] {
	case "deadline":
		deadline := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
		d, err := time.ParseInLocation("2006-01-02", deadline.Value, time.UTC)
		if err != nil {
			d, err = time.ParseInLocation("2006/01/02", deadline.Value, time.UTC)
			if err != nil {
				ds.FollowupMessageCreate(ic.Interaction, false, &discordgo.WebhookParams{
					Flags: 64,
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "Error",
							Description: err.Error(),
							Color:       0xEE2211,
						},
					},
				})
				return
			}
		}

		ad := asana.Date(d)
		task.Update(AsanaClient, &asana.UpdateTaskRequest{
			TaskBase: asana.TaskBase{
				DueOn: &ad,
			},
		})
		ds.FollowupMessageCreate(ic.Interaction, false, &discordgo.WebhookParams{
			Flags:   64,
			Content: fmt.Sprintf("Task is now due on <t:%d:D>", d.Add(time.Hour*12).Unix()),
		})
	case "desc":
		name := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
		notes := data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
		go task.Update(AsanaClient, &asana.UpdateTaskRequest{
			TaskBase: asana.TaskBase{
				Name:  name.Value,
				Notes: notes.Value,
			},
		})
		orig, err := ds.ChannelMessage(ic.ChannelID, args[3])
		if err != nil {
			ds.FollowupMessageCreate(ic.Interaction, false, &discordgo.WebhookParams{
				Flags:   64,
				Content: "Task is updated but encountered a Discord error",
			})
			return
		}
		orig.Embeds[0].Fields[0] = &discordgo.MessageEmbedField{
			Name:  name.Value,
			Value: notes.Value,
		}
		ds.ChannelMessageEditEmbeds(ic.ChannelID, args[3], orig.Embeds)
		ds.InteractionResponseDelete(ic.Interaction)
	}
}

func handleAsanaSubmit(ds *discordgo.Session, ic *discordgo.InteractionCreate) (*discordgo.InteractionResponseData, error) {
	go func() {
		data := ic.ModalSubmitData()
		raw := strings.Split(data.CustomID, ",")
		if raw[0] == "asana" {
			handleAsanaUpdate(ds, ic)
			return
		}
		project := raw[0]
		assignee := raw[1]

		name := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
		notes := data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)

		task, err := AsanaClient.CreateTask(&asana.CreateTaskRequest{
			TaskBase: asana.TaskBase{
				Name:      name.Value,
				Notes:     notes.Value,
				Completed: asana.Bool(false),
			},
			Assignee:  assignee,
			Followers: []string{},
			Workspace: WorkspaceID,
			Projects:  []string{project},
			Tags:      []string{},
		})
		if err != nil {
			ds.InteractionResponseEdit(ic.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{
					{
						Title:       "Error",
						Description: err.Error(),
						Color:       0xEE2211,
					},
				},
			})
			return
		}

		if len(task.Name) > 256 {
			task.Name = task.Name[:250] + "..."
		}
		if len(notes.Value) > 1024 {
			notes.Value = notes.Value[:1020] + "..."
		}

		p := ProjectMap[project]
		sections, _, err := p.Sections(AsanaClient, &asana.Options{Limit: 25})
		if err != nil {
			ds.InteractionResponseEdit(ic.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{
					{
						Title:       "Error",
						Description: err.Error(),
						Color:       0xEE2211,
					},
				},
			})
			return
		}
		var sectionOptions []discordgo.SelectMenuOption
		for _, s := range sections {
			sectionOptions = append(sectionOptions, discordgo.SelectMenuOption{
				Label: s.Name,
				Value: s.ID,
			})
		}

		ds.InteractionResponseEdit(ic.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Color: 0xF06A6A,
					URL:   fmt.Sprintf("https://app.asana.com/0/%s/%s", project, task.ID),
					Title: "New Task Created in " + p.Name,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  task.Name,
							Value: notes.Value,
						},
						{
							Name:   "Assignee",
							Value:  UserMap[assignee].Name,
							Inline: true,
						},
					},
				},
			},
			Components: &[]discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							MenuType:    discordgo.StringSelectMenu,
							CustomID:    "asana,section," + task.ID,
							Placeholder: "Set a section for the task",
							Options:     sectionOptions,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Task Deadline",
							Style:    discordgo.PrimaryButton,
							CustomID: "asana,deadline," + task.ID,
							Emoji: discordgo.ComponentEmoji{
								Name: "⏰",
							},
						},
						discordgo.Button{
							Label:    "Task Description",
							Style:    discordgo.PrimaryButton,
							CustomID: "asana,desc," + task.ID,
							Emoji: discordgo.ComponentEmoji{
								Name: "🗒️",
							},
						},
						discordgo.Button{
							Label: "Open in Asana",
							Style: discordgo.LinkButton,
							URL:   fmt.Sprintf("https://app.asana.com/0/%s/%s", project, task.ID),
						},
					},
				},
			},
		})
	}()
	return nil, nil
}

func handleAsanaComponent(ds *discordgo.Session, ic *discordgo.InteractionCreate, args []string) (*discordgo.InteractionResponseData, error) {
	data := ic.MessageComponentData()
	switch args[1] {
	case "section":
		return handleSection(ds, ic, args)
	case "deadline":
		return &discordgo.InteractionResponseData{
			Title:    "Update Task Deadline",
			CustomID: data.CustomID + "," + ic.Message.ID,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.TextInput{
							Label:       "Task Title (YYYY-MM-DD)",
							Style:       discordgo.TextInputShort,
							CustomID:    "date",
							Value:       "",
							Required:    true,
							Placeholder: "YYYY-MM-DD",
						},
					},
				},
			},
		}, nil
	case "desc":
		task := &asana.Task{ID: args[2]}
		if err := task.Fetch(AsanaClient); err != nil {
			return nil, err
		}

		return &discordgo.InteractionResponseData{
			Title:    "Update Task",
			CustomID: data.CustomID + "," + ic.Message.ID,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.TextInput{
							Label:       "Task Title",
							Style:       discordgo.TextInputShort,
							CustomID:    "title",
							Value:       task.Name,
							Required:    true,
							Placeholder: "Try Pineapple Pizza",
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.TextInput{
							Label:       "Task Description",
							Style:       discordgo.TextInputParagraph,
							CustomID:    "desc",
							Value:       task.Notes,
							Placeholder: "Praise Pineapple Pizza",
						},
					},
				},
			},
		}, nil
	}
	return nil, nil
}

func handleSection(ds *discordgo.Session, ic *discordgo.InteractionCreate, args []string) (*discordgo.InteractionResponseData, error) {
	go func() {
		data := ic.MessageComponentData()

		task := &asana.Task{ID: args[2]}
		if err := task.Fetch(AsanaClient); err != nil {
			ds.FollowupMessageCreate(ic.Interaction, false, &discordgo.WebhookParams{
				Flags: 64,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Error",
						Description: err.Error(),
						Color:       0xEE2211,
					},
				},
			})
			return
		}

		section := data.Values[0]

		if err := task.AddProject(AsanaClient, &asana.AddProjectRequest{
			Project: task.Projects[0].ID,
			Section: section,
		}); err != nil {
			ds.FollowupMessageCreate(ic.Interaction, false, &discordgo.WebhookParams{
				Flags: 64,
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Error",
						Description: err.Error(),
						Color:       0xEE2211,
					},
				},
			})
			return
		}

		orig, err := ds.InteractionResponse(ic.Interaction)
		if err != nil {
			ds.FollowupMessageCreate(ic.Interaction, false, &discordgo.WebhookParams{
				Flags:   64,
				Content: "Task is updated but encountered a Discord error",
			})
		}

		sm := orig.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.SelectMenu)
		for _, s := range sm.Options {
			if s.Value == section {
				s.Default = true
				sm.Placeholder = s.Label
			}
		}

		ds.InteractionResponseEdit(ic.Interaction, &discordgo.WebhookEdit{
			Components: &orig.Components,
			Embeds:     &orig.Embeds,
		})
	}()
	return nil, nil
}
