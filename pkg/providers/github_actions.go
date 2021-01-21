package providers

import (
	"fmt"
	"github.com/thrawny/openci/pkg/git"
	"net/url"
)

const githubBaseUrl = "https://github.com"

type GitHubActions struct{}

func (GitHubActions) GetProjectURL(remote git.Remote, branch string) string {
	query := url.QueryEscape(fmt.Sprintf("branch=%s", branch))
	return fmt.Sprintf("%s/%s/%s/actions?query=%s", githubBaseUrl, remote.Org, remote.Project, query)
}
