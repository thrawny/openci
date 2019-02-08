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
			},
			expected: "https://circleci.com/gh/foo/baz",
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
			},
			expected: "https://travis-ci.org/foo/baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.args.provider.GetProjectURL(tt.args.remote))
		})
	}
}
