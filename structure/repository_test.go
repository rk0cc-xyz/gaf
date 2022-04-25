package structure_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/17media/structs"
	"github.com/rk0cc-xyz/gaf/structure"
)

func TestSingleRepoParse(t *testing.T) {
	const apistr = `
	{
		"id": 482776754,
		"node_id": "R_kgDOHMaWsg",
		"name": "tempcord",
		"full_name": "rk0cc/tempcord",
		"private": false,
		"owner": {
		  "login": "rk0cc",
		  "id": 70585816,
		  "node_id": "MDQ6VXNlcjcwNTg1ODE2",
		  "avatar_url": "https://avatars.githubusercontent.com/u/70585816?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/rk0cc",
		  "html_url": "https://github.com/rk0cc",
		  "followers_url": "https://api.github.com/users/rk0cc/followers",
		  "following_url": "https://api.github.com/users/rk0cc/following{/other_user}",
		  "gists_url": "https://api.github.com/users/rk0cc/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/rk0cc/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/rk0cc/subscriptions",
		  "organizations_url": "https://api.github.com/users/rk0cc/orgs",
		  "repos_url": "https://api.github.com/users/rk0cc/repos",
		  "events_url": "https://api.github.com/users/rk0cc/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/rk0cc/received_events",
		  "type": "User",
		  "site_admin": false
		},
		"html_url": "https://github.com/rk0cc/tempcord",
		"description": "Profiled body temperature recorder for warm-blood animals",
		"fork": false,
		"url": "https://api.github.com/repos/rk0cc/tempcord",
		"forks_url": "https://api.github.com/repos/rk0cc/tempcord/forks",
		"keys_url": "https://api.github.com/repos/rk0cc/tempcord/keys{/key_id}",
		"collaborators_url": "https://api.github.com/repos/rk0cc/tempcord/collaborators{/collaborator}",
		"teams_url": "https://api.github.com/repos/rk0cc/tempcord/teams",
		"hooks_url": "https://api.github.com/repos/rk0cc/tempcord/hooks",
		"issue_events_url": "https://api.github.com/repos/rk0cc/tempcord/issues/events{/number}",
		"events_url": "https://api.github.com/repos/rk0cc/tempcord/events",
		"assignees_url": "https://api.github.com/repos/rk0cc/tempcord/assignees{/user}",
		"branches_url": "https://api.github.com/repos/rk0cc/tempcord/branches{/branch}",
		"tags_url": "https://api.github.com/repos/rk0cc/tempcord/tags",
		"blobs_url": "https://api.github.com/repos/rk0cc/tempcord/git/blobs{/sha}",
		"git_tags_url": "https://api.github.com/repos/rk0cc/tempcord/git/tags{/sha}",
		"git_refs_url": "https://api.github.com/repos/rk0cc/tempcord/git/refs{/sha}",
		"trees_url": "https://api.github.com/repos/rk0cc/tempcord/git/trees{/sha}",
		"statuses_url": "https://api.github.com/repos/rk0cc/tempcord/statuses/{sha}",
		"languages_url": "https://api.github.com/repos/rk0cc/tempcord/languages",
		"stargazers_url": "https://api.github.com/repos/rk0cc/tempcord/stargazers",
		"contributors_url": "https://api.github.com/repos/rk0cc/tempcord/contributors",
		"subscribers_url": "https://api.github.com/repos/rk0cc/tempcord/subscribers",
		"subscription_url": "https://api.github.com/repos/rk0cc/tempcord/subscription",
		"commits_url": "https://api.github.com/repos/rk0cc/tempcord/commits{/sha}",
		"git_commits_url": "https://api.github.com/repos/rk0cc/tempcord/git/commits{/sha}",
		"comments_url": "https://api.github.com/repos/rk0cc/tempcord/comments{/number}",
		"issue_comment_url": "https://api.github.com/repos/rk0cc/tempcord/issues/comments{/number}",
		"contents_url": "https://api.github.com/repos/rk0cc/tempcord/contents/{+path}",
		"compare_url": "https://api.github.com/repos/rk0cc/tempcord/compare/{base}...{head}",
		"merges_url": "https://api.github.com/repos/rk0cc/tempcord/merges",
		"archive_url": "https://api.github.com/repos/rk0cc/tempcord/{archive_format}{/ref}",
		"downloads_url": "https://api.github.com/repos/rk0cc/tempcord/downloads",
		"issues_url": "https://api.github.com/repos/rk0cc/tempcord/issues{/number}",
		"pulls_url": "https://api.github.com/repos/rk0cc/tempcord/pulls{/number}",
		"milestones_url": "https://api.github.com/repos/rk0cc/tempcord/milestones{/number}",
		"notifications_url": "https://api.github.com/repos/rk0cc/tempcord/notifications{?since,all,participating}",
		"labels_url": "https://api.github.com/repos/rk0cc/tempcord/labels{/name}",
		"releases_url": "https://api.github.com/repos/rk0cc/tempcord/releases{/id}",
		"deployments_url": "https://api.github.com/repos/rk0cc/tempcord/deployments",
		"created_at": "2022-04-18T08:45:15Z",
		"updated_at": "2022-04-20T05:37:00Z",
		"pushed_at": "2022-04-20T05:36:27Z",
		"git_url": "git://github.com/rk0cc/tempcord.git",
		"ssh_url": "git@github.com:rk0cc/tempcord.git",
		"clone_url": "https://github.com/rk0cc/tempcord.git",
		"svn_url": "https://github.com/rk0cc/tempcord",
		"homepage": "",
		"size": 87,
		"stargazers_count": 0,
		"watchers_count": 0,
		"language": "Dart",
		"has_issues": true,
		"has_projects": true,
		"has_downloads": true,
		"has_wiki": true,
		"has_pages": false,
		"forks_count": 0,
		"mirror_url": null,
		"archived": false,
		"disabled": false,
		"open_issues_count": 0,
		"license": {
		  "key": "other",
		  "name": "Other",
		  "spdx_id": "NOASSERTION",
		  "url": null,
		  "node_id": "MDc6TGljZW5zZTA="
		},
		"allow_forking": true,
		"is_template": false,
		"topics": [
		  "archive",
		  "body-temperature",
		  "dart",
		  "flutter",
		  "health",
		  "health-check",
		  "healthcare",
		  "logging"
		],
		"visibility": "public",
		"forks": 0,
		"open_issues": 0,
		"watchers": 0,
		"default_branch": "main"
	}
	`

	mockMap := make(map[string]interface{})

	md := json.NewDecoder(bytes.NewBufferString(apistr))
	md.UseNumber()
	md.Decode(&mockMap)

	parsed, parsederr := structure.ParseFromRESTMap(mockMap)
	if parsederr != nil {
		t.Fatal(parsederr)
	}

	expected := structure.GitHubRepositoryStructure{
		Name:        "tempcord",
		Site:        "https://github.com/rk0cc/tempcord",
		IsFork:      false,
		IsArchive:   false,
		Starred:     0,
		Watched:     0,
		Forked:      0,
		OpenedIssue: 0,
		Language:    "Dart",
		License:     "Other",
		Topics: []string{
			"archive",
			"body-temperature",
			"dart",
			"flutter",
			"health",
			"health-check",
			"healthcare",
			"logging",
		},
	}

	expectedMap := structs.Map(expected)
	actualMap := structs.Map(parsed)

	for k, v := range expectedMap {
		if k == "Topics" {
			for idx, item := range expectedMap[k].([]string) {
				if actualMap[k].([]string)[idx] != item {
					t.Fail()
				}
			}
		} else {
			if actualMap[k] != v {
				t.Fail()
			}
		}
	}
}
