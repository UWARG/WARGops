package bot

import (
	"log"
	"time"

	"bitbucket.org/mikehouston/asana-go"
	"github.com/UWARG/WARGops/bot/commands"
	"github.com/bwmarrin/discordgo"
)

func Start(token string, asanaToken string) error {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers
	session.AddHandler(commands.OnInteractionCommand)
	session.AddHandler(commands.OnAutocomplete)
	session.AddHandler(commands.OnModalSubmit)
	session.AddHandler(commands.OnMessageComponent)

	if err := session.Open(); err != nil {
		log.Fatalf("error opening connection to Discord: %v", err)
	}

	client := asana.NewClientWithAccessToken(asanaToken)
	commands.AsanaClient = client

	w := &asana.Workspace{
		ID: commands.WorkspaceID,
	}
	refreshAsana(client, w)
	ticker := time.NewTicker(time.Minute * 30)
	for range ticker.C {
		log.Println("refreshing asana...")
		refreshAsana(client, w)
	}

	return nil
}

func refreshAsana(client *asana.Client, w *asana.Workspace) {
	projects, _, err := w.Projects(client, &asana.Options{
		Fields: []string{"archived", "name"},
		Limit:  100,
	})
	if err != nil {
		log.Printf("could not load asana projects: %v", err)
	}
	commands.Projects = make([]*asana.Project, 0, len(projects))
	commands.ProjectNames = make([]string, 0, len(projects))
	for _, p := range projects {
		if *p.Archived {
			continue
		}
		commands.ProjectMap[p.ID] = p
		commands.Projects = append(commands.Projects, p)
		commands.ProjectNames = append(commands.ProjectNames, p.Name)
	}

	users, err := w.AllUsers(client, &asana.Options{
		Fields: []string{"name", "email"},
		Limit:  100,
	})
	if err != nil {
		log.Printf("could not load asana users: %v", err)
	}
	commands.Users = make([]*asana.User, 0, len(users))
	commands.UserNames = make([]string, 0, len(users))
	for _, u := range users {
		commands.UserMap[u.ID] = u
		commands.Users = append(commands.Users, u)
		commands.UserNames = append(commands.UserNames, u.Name)
	}

	log.Println("asana items reloaded")
}
