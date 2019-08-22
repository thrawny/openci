package providers

import "github.com/thrawny/openci/pkg/git"

type Provider interface {
	GetProjectURL(remote git.Remote, branch string) string
}
