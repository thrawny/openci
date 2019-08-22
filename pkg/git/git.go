package git

import (
	"errors"
	"fmt"
	"github.com/thrawny/openci/pkg/util"
	"os/exec"
	"regexp"
)

var re = regexp.MustCompile(`(?m)^(https|git)(://|@)([^/:]+)[/:]([^/:]+)/([^\.]+)(\.git)?$`)

type Remote struct {
	Domain  string
	Org     string
	Project string
}

func ParseGitURL(url string) (Remote, error) {
	matches := re.FindStringSubmatch(url)
	if !(len(matches) >= 6) {
		return Remote{}, fmt.Errorf("could not parse git url: %s", url)
	}
	return Remote{
		Domain:  matches[3],
		Org:     matches[4],
		Project: matches[5],
	}, nil
}

func RepoRoot(wd string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = wd
	res, err := util.RunCmd(cmd)
	if err != nil {
		return "", errors.New(res)
	}
	return res, nil
}

func RemoteURL(wd string) (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	cmd.Dir = wd
	res, err := util.RunCmd(cmd)
	if err != nil {
		return "", errors.New(res)
	}
	return res, nil
}

func Branch(wd string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = wd
	res, err := util.RunCmd(cmd)
	if err != nil {
		return "", errors.New(res)
	}
	return res, nil
}
