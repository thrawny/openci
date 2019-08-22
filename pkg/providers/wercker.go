package providers

import (
	"fmt"
	"github.com/thrawny/openci/pkg/git"
)

const werckerBaseURL = "https://app.wercker.com"

type Wercker struct{}

func (Wercker) GetProjectURL(remote git.Remote, branch string) string {
	return fmt.Sprintf("%s/%s/%s", werckerBaseURL, remote.Org, remote.Project)
}
