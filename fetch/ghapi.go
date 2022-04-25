package fetch

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rk0cc-xyz/gaf/structure"
	"github.com/tomnomnom/linkheader"
)

const (
	// The fetch result will be ordered by repository name.
	REPO_FULL_NAME = "full_name"

	// The fetch result will be ordered by time of repository created.
	REPO_CREATE_TIME = "created"

	// The fetch result will be ordered by time when repository updated.
	REPO_UPDATE_TIME = "updated"

	// The fetch result will be ordered by time when repository pushed a new commitment.
	REPO_PUSH_TIME = "pushed"

	// Ascending order
	ASC_ORDER = "asc"

	// Descending order
	DESC_ORDER = "desc"

	fetch_url = "https://api.github.com/user/repos?visibility=public&affiliation=owner"

	ghpat_env = "GAF_GITHUB_TOKEN"
)

// An error which the token is undefined on environment variable `GAF_GITHUB_TOKEN`.
type GitHubPATMissingError struct{}

func (gpaterr GitHubPATMissingError) Error() string {
	return "githubpatmissingerror: Can not find token in `" + ghpat_env + "`."
}

// Preference for fetching GitHub repository according to https://docs.github.com/en/rest/repos/repos#list-repositories-for-the-authenticated-user
type GitHubRepositoryAPIFetchSetting struct {

	// Define how to order the repository.
	OrderBy string

	// Define `OrderBy`'s sorting.
	OrderSort string
}

// A structre uses for fetching GitHub repository content.
type GitHubRepositoryAPIFetch struct {
	// Token provided to getting data.
	token string

	// Preference for fetching GitHub repository data.
	Setting GitHubRepositoryAPIFetchSetting
}

// Construct new GitHubRepositoryAPIFetch.
//
// It requried GitHub Personal Access Token provided in environment variable "GAF_GITHUB_TOKEN".
func NewGitHubRepositoryAPIFetch() (*GitHubRepositoryAPIFetch, error) {
	t := os.Getenv(ghpat_env)

	if len(t) == 0 {
		return nil, GitHubPATMissingError{}
	}

	return &GitHubRepositoryAPIFetch{
		token: t,
		Setting: GitHubRepositoryAPIFetchSetting{
			OrderBy:   REPO_FULL_NAME,
			OrderSort: ASC_ORDER,
		},
	}, nil
}

// Fetch a single page's context of t GitHub API result.
//
// Parameter `page` is optional that to fetch content in GitHub API page. By default, it uses page 1.
//
// The return parameters repersents returned structure of the repository, a boolean repersent it has next page and the error.
func (graf GitHubRepositoryAPIFetch) FetchPage(page ...uint64) ([]structure.GitHubRepositoryStructure, *bool, error) {
	var target_page uint64

	if len(page) > 1 {
		return nil, nil, errors.New("page parameter should be either omitted or provide a single parameter")
	} else if len(page) == 0 {
		target_page = 1
	} else {
		target_page = page[0]
	}

	resprest, respresterr := graf.downloadREST(target_page)
	if respresterr != nil {
		return nil, nil, respresterr
	} else if resprest.StatusCode != http.StatusOK {
		return nil, nil, errors.New("responsed with error from GitHub API")
	}

	// Check does continue in next page
	has_next_page := false
	hls := resprest.Header.Get("link")
	if len(hls) > 0 {
		hl := linkheader.Parse(hls)

		for _, shl := range hl {
			if strings.ToLower(shl.Rel) == "next" {
				has_next_page = true
				break
			}
		}
	}

	respbody, respbodyerr := io.ReadAll(resprest.Body)
	if respbodyerr != nil {
		return nil, nil, respbodyerr
	}

	decoded, decodederr := decodeRestByteToMapArray(respbody)
	if decodederr != nil {
		return nil, nil, decodederr
	}

	pgn := make([]structure.GitHubRepositoryStructure, 0)
	for _, m := range decoded {
		parsed, parsederr := structure.ParseFromRESTMap(m)
		if parsederr != nil {
			return nil, nil, parsederr
		}
		pgn = append(pgn, *parsed)
	}

	resprest.Body.Close()

	return pgn, &has_next_page, nil
}

// Return a response with specify page.
func (graf GitHubRepositoryAPIFetch) downloadREST(page uint64) (*http.Response, error) {
	param := make(map[string]string)
	param["page"] = strconv.FormatUint(page, 10)
	param["sort"] = graf.Setting.OrderBy
	param["direction"] = graf.Setting.OrderSort

	cru := []string{fetch_url}
	for k, v := range param {
		cru = append(cru, strings.Join([]string{
			k,
			v,
		}, "="))
	}

	comp_req_url := strings.Join(cru, "&")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, reqerr := http.NewRequest("GET", comp_req_url, nil)
	if reqerr != nil {
		return nil, reqerr
	}

	req.Header.Set("User-Agent", "Mozilla 5.0")
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", graf.token)

	return client.Do(req)
}

// Decode byte to an array of GitHub repository API.
func decodeRestByteToMapArray(b []byte) ([]map[string]interface{}, error) {
	var grn []map[string]interface{}

	d := json.NewDecoder(bytes.NewBuffer(b))
	d.UseNumber()
	if derr := d.Decode(&grn); derr != nil {
		return nil, derr
	}

	return grn, nil
}
