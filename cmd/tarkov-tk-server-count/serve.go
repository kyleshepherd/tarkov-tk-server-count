package main

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/kyleshepherd/tarkov-tk-server-count/internal/api"
	"github.com/spf13/cobra"
)

func serveCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "serve",
		Short:   "Serves API",
		Example: "tarkov-tk-server-count serve",
		RunE:    serveRun,
	}
}

func serveRun(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	s, err := discordgo.New("Bot " + cfg.Discord.BotToken)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info().Msgf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = s.Open()
	if err != nil {
		log.Error().Err(err)
		return err
	}

	defer s.Close()

	api := api.NewAPI(ctx, s)

	log.Info().Msgf("up and running on port %s", api.Server.Addr)
	err = api.Server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
