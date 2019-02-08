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
	return nil, errors.New("could not open a ci provider")
}

func Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if !git.IsGitRepo(wd) {
		return errors.New("current directory is not a git repository")
	}
	url, err := git.RemoteURL(wd)
	if err != nil {
		return err
	}
	ciProvider, err := CiProvider(wd)
	if err != nil {
		return err
	}
	remote, err := git.ParseGitURL(url)
	if err != nil {
		return err
	}
	return open.Run(ciProvider.GetProjectURL(remote))
}
