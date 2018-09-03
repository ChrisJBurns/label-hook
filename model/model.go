package model

// Event is the parent structure of the request body
type Event struct {
	PRs PR `json:"pull_request"`
}

// PR is the data structure of the "pull_request" dataset within Event
type PR struct {
	Labels []Label `json:"labels"`
	Head   Head    `json:"head"`
}

// Label is the data structure for the "labels" dataset within PR
type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// Head is the data structure for the "head" dataset within PR
type Head struct {
	Sha  string `json:"sha"`
	Repo Repo   `json:"repo"`
}

// Repo is the data structure for the "repo" dataset within Head
type Repo struct {
	Name string `json:"name"`
}
