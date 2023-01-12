package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

type API struct {
	Server  *http.Server
	Discord *discordgo.Session
}

func NewAPI(ctx context.Context, discord *discordgo.Session) *API {
	api := &API{
		Discord: discord,
	}

	fmt.Printf("%d", len(api.Discord.State.Guilds))

	router := mux.NewRouter()

	router.HandleFunc("/count", api.GetCount).Methods(http.MethodGet)

	handler := cors.Default().Handler(router)

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":5001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	api.Server = srv

	return api
}

type Response struct {
	Count int `json:"count"`
}

func (a *API) GetCount(w http.ResponseWriter, r *http.Request) {
	res := &Response{
		Count: len(a.Discord.State.Guilds),
	}

	rJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Error().Err(err).Msg(err.Error())
		return
	}

	_, err = w.Write(rJSON)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Error().Err(err).Msg(err.Error())
		return
	}
}
