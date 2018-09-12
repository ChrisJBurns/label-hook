package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	configPath, err := GetConfigurationPath()
	if err != nil {
		log.Print(err)
		return
	}

	cfg, err = config.LoadConfiguration(configPath)
	if err != nil {
		log.Print(err)
		return
	}

	if err := ValidateConfig(cfg); err != nil {
		log.Print(err)
		return
	}

	CreateClient()

	router := mux.NewRouter()
	router.HandleFunc("/", ProcessPREvent).Methods("POST")

	log.Printf("Started label-hook on %s:%s", cfg.Host, cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Host+":"+cfg.Port, router))
}

// ValidateConfig checks all properties loaded from the config and validates them to
// check if they are empty. If so, it will error on the property that is empty as they're
// all mandatory.
func ValidateConfig(cfg config.Config) error {
	if cfg == (config.Config{}) {
		return errors.New("Config is empty")
	}

	if cfg.Organisation == "" {
		return errors.New("No organisation has been configured")
	}

	if cfg.Host == "" {
		return errors.New("No host has been configured")
	}

	if cfg.Port == "" {
		return errors.New("No port has been configured")
	}

	if cfg.AccessToken == "" {
		return errors.New("No access token has been configured")
	}

	return nil
}

// GetConfigurationPath gets the config file path that is passed in as an argument to the application.
// If there is anything other than the config file path provided as an argument, throw an error. This
// applies to if there are no parameters provided and/or if there are more than one.
func GetConfigurationPath() (string, error) {
	if len(os.Args) != 2 {
		return "", errors.New("Please specify the config location as an argument and only that")
	}

	return os.Args[1], nil
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
