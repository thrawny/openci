package open

import (
	"errors"
	"github.com/skratchdot/open-golang/open"
	"github.com/thrawny/openci/pkg/git"
	"github.com/thrawny/openci/pkg/providers"
	"github.com/thrawny/openci/pkg/util"
	"os"
	"path"
)

func CiProvider(wd string) (providers.Provider, error) {
	if util.FileExists(path.Join(wd, ".circleci/config.yml")) {
		return providers.CirleCI{}, nil
	}
	if util.FileExists(path.Join(wd, ".travis.yml")) {
		return providers.TravisCI{}, nil
	}
	if util.FileExists(path.Join(wd, "wercker.yml")) {
		return providers.Wercker{}, nil
	}
	return nil, errors.New("could not detect a ci provider")
}

func Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	repoRoot, err := git.RepoRoot(wd)
	if err != nil {
		return err
	}
	url, err := git.RemoteURL(repoRoot)
	if err != nil {
		return err
	}
	ciProvider, err := CiProvider(repoRoot)
	if err != nil {
		return err
	}
	remote, err := git.ParseGitURL(url)
	if err != nil {
		return err
	}
	return open.Run(ciProvider.GetProjectURL(remote))
}
