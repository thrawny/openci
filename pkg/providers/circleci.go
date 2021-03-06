package providers

import (
	"fmt"
	"github.com/thrawny/openci/pkg/git"
)

const circleBaseURL = "https://circleci.com"

type CirleCI struct{}

func (CirleCI) GetProjectURL(remote git.Remote, branch string) string {
	gitProvider := "gh"
	if remote.Domain == "bitbucket.org" {
		gitProvider = "bb"
	}
	return fmt.Sprintf("%s/%s/%s/workflows/%s/tree/%s", circleBaseURL, gitProvider, remote.Org, remote.Project, branch)
}
