package git

import (
	"fmt"
	"github.com/thrawny/openci/pkg/util"
	"os/exec"
	"path"
	"regexp"
	"strings"
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

func IsGitRepo(wd string) bool {
	return util.FileExists(path.Join(wd, ".git"))
}

func RemoteURL(wd string) (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	cmd.Dir = wd
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
