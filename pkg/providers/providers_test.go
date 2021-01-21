package providers

import (
	"github.com/stretchr/testify/assert"
	"github.com/thrawny/openci/pkg/git"
	"testing"
)

func TestProviders(t *testing.T) {
	type args struct {
		remote   git.Remote
		provider Provider
		branch   string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "circleci",
			args: args{
				remote: git.Remote{
					Domain:  "github.com",
					Org:     "foo",
					Project: "baz",
				},
				provider: CirleCI{},
				branch:   "qux",
			},
			expected: "https://circleci.com/gh/foo/workflows/baz/tree/qux",
		},
		{
			name: "wercker",
			args: args{
				remote: git.Remote{
					Domain:  "github.com",
					Org:     "foo",
					Project: "baz",
				},
				provider: Wercker{},
				branch:   "qux",
			},
			expected: "https://app.wercker.com/foo/baz",
		},
		{
			name: "travis",
			args: args{
				remote: git.Remote{
					Domain:  "github.com",
					Org:     "foo",
					Project: "baz",
				},
				provider: TravisCI{},
				branch:   "qux",
			},
			expected: "https://travis-ci.org/foo/baz",
		},
		{
			name: "github actions",
			args: args{
				remote: git.Remote{
					Domain:  "github.com",
					Org:     "foo",
					Project: "baz",
				},
				provider: GitHubActions{},
				branch:   "qux",
			},
			expected: "https://github.com/foo/baz/actions?query=branch%3Dqux",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.args.provider.GetProjectURL(tt.args.remote, tt.args.branch))
		})
	}
}
