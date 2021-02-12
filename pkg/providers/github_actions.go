package providers

import (
	"fmt"
	"github.com/thrawny/openci/pkg/git"
	"net/url"
)

const githubBaseUrl = "https://github.com"

type GitHubActions struct{}

func (GitHubActions) GetProjectURL(remote git.Remote, branch string) string {
	return fmt.Sprintf("%s/%s/%s/actions?query=branch%%3A%s", githubBaseUrl, remote.Org, remote.Project, url.QueryEscape(branch))
}
