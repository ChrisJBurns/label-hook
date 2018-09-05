package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/ChrisJBurns/label-hook/config"
	"github.com/ChrisJBurns/label-hook/model"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

var client *github.Client
var ctx context.Context
var cfg config.Config

func main() {
	cfg = config.LoadConfiguration()
	if cfg == (config.Config{}) {
		log.Println("Configuration loading failed: returned empty")
		return
	}

	CreateClient()

	router := mux.NewRouter()
	router.HandleFunc("/", ProcessPREvent).Methods("POST")

	if cfg.Port == "" {
		log.Println("Configured port is empty")
		return
	}
	if cfg.Host == "" {
		log.Println("Configured host is empty")
		return
	}
	log.Printf("label-hook started and listening on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Host+":"+cfg.Port, router))
}

// ProcessPREvent does processing around PR
func ProcessPREvent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error when reading information from request body")
		panic(err)
	}

	var bodyMap = model.Event{}
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		log.Println("error when unmarshalling body into data object")
		log.Fatal(err)
	}

	labels := bodyMap.PRs.Labels

	if len(labels) > 0 {
		log.Println("sending success status")
		client.Repositories.CreateStatus(ctx,
			cfg.Organisation,
			bodyMap.PRs.Head.Repo.Name,
			bodyMap.PRs.Head.Sha,
			CreateRepoStatus("success"))

		log.Println("success status sent")
		w.Write([]byte(strconv.Itoa(1)))
		return
	}

	log.Println("sending failure status")
	client.Repositories.CreateStatus(ctx,
		cfg.Organisation,
		bodyMap.PRs.Head.Repo.Name,
		bodyMap.PRs.Head.Sha,
		CreateRepoStatus("failure"))

	log.Println("failure status sent")

	w.Write([]byte(strconv.Itoa(0)))
	return
}

// CreateRepoStatus creates a repo status object to send to GitHub status API.
// It will either send a "success" or "failure" status code
func CreateRepoStatus(status string) *github.RepoStatus {
	reviewLabels := "review/labels"
	description := "Please add either; bug, new feature or an improvement label"
	return &github.RepoStatus{
		State:       &status,
		Context:     &reviewLabels,
		Description: &description,
	}
}

// CreateClient creates the GitHub client for the accessing of the
// GitHub APIs using OAuth (Personal Access Token)
func CreateClient() {
	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
}
