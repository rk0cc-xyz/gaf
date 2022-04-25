package structure

import (
	"encoding/json"
	"strconv"
)

// Simplified structure repersenting a single repository data from GitHub repository API.
type GitHubRepositoryStructure struct {
	// Name of repository.
	Name string `json:"name,omitempty"`
	// URL link to repository's website.
	Site string `json:"site,omitempty"`
	// Does this repository is fork from other.
	IsFork bool `json:"is_fork"`
	// Does this repository set as archive state.
	IsArchive bool `json:"is_archive"`
	// Count of repository starred.
	Starred uint64 `json:"starred,string"`
	// Count of repository watched.
	Watched uint64 `json:"watched,string"`
	// Count of reposiotry has been forked.
	Forked uint64 `json:"forked,string"`
	// Count of issues currently opened in this repository.
	OpenedIssue uint64 `json:"opened_issue,string"`
	// Programming language uses in this repository.
	//
	// It possibility returns "Other" when GitHub API returns null in this field.
	Language string `json:"language,omitempty"`
	// License uses in this repository.
	//
	// It returns "None" if GitHub API return null in this field.
	License string `json:"license,omitempty"`
	// An array of string which indicate topic of this repository.
	//
	// It should be an empty array if no topic is defined.
	Topics []string `json:"topics"`
}

// Parse a single node of GitHub repository API data to Go's structure.
//
// The rest parameter must be came form `json.Decode` with `UseNumber()` invoked.
func ParseFromRESTMap(rest map[string]interface{}) (*GitHubRepositoryStructure, error) {
	// Star
	sc, scerr := jsonNumToUint64(rest["stargazers_count"].(json.Number))
	if scerr != nil {
		return nil, scerr
	}

	// Watch
	wc, wcerr := jsonNumToUint64(rest["watchers_count"].(json.Number))
	if wcerr != nil {
		return nil, wcerr
	}

	// Fork
	fc, fcerr := jsonNumToUint64(rest["forks_count"].(json.Number))
	if fcerr != nil {
		return nil, fcerr
	}

	// Open issues
	oic, oicerr := jsonNumToUint64(rest["open_issues_count"].(json.Number))
	if oicerr != nil {
		return nil, oicerr
	}

	var lang string
	var lic string

	// Language
	rlang := rest["language"]
	if rlang == nil {
		lang = "Other"
	} else {
		lang = rlang.(string)
	}

	// License
	rlic := rest["license"]
	if rlic == nil {
		lic = "None"
	} else {
		lic = rlic.(map[string]interface{})["name"].(string)
	}

	ogtopic := rest["topics"].([]interface{})
	stopic := make([]string, len(ogtopic))

	for idx, item := range ogtopic {
		stopic[idx] = item.(string)
	}

	return &GitHubRepositoryStructure{
		Name:        rest["name"].(string),
		Site:        rest["html_url"].(string),
		IsFork:      rest["fork"].(bool),
		IsArchive:   rest["archived"].(bool),
		Starred:     sc,
		Watched:     wc,
		Forked:      fc,
		OpenedIssue: oic,
		Language:    lang,
		License:     lic,
		Topics:      stopic,
	}, nil
}

func jsonNumToUint64(num json.Number) (uint64, error) {
	return strconv.ParseUint(string(num), 10, 64)
}
