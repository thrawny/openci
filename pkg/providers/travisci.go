package providers

import (
	"fmt"
	"github.com/thrawny/openci/pkg/git"
)

const travisBaseURL = "https://travis-ci.org"

type TravisCI struct{}

func (TravisCI) GetProjectURL(remote git.Remote) string {
	return fmt.Sprintf("%s/%s/%s", travisBaseURL, remote.Org, remote.Project)
}
