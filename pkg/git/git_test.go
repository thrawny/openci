package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseGitURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name      string
		args      args
		expected  Remote
		expectErr bool
	}{
		{
			name: "git@",
			args: args{
				url: "git@github.com:thrawny/openci.git",
			},
			expected: Remote{
				Domain:  "github.com",
				Org:     "thrawny",
				Project: "openci",
			},
			expectErr: false,
		},
		{
			name: "https",
			args: args{
				url: "https://github.com/some-user/my-repo.git",
			},
			expected: Remote{
				Domain:  "github.com",
				Org:     "some-user",
				Project: "my-repo",
			},
			expectErr: false,
		},
		{
			name: "https without .git",
			args: args{
				url: "https://github.com/some-user/my-repo",
			},
			expected: Remote{
				Domain:  "github.com",
				Org:     "some-user",
				Project: "my-repo",
			},
			expectErr: false,
		},
		{
			name: "failed to parse",
			args: args{
				url: "https://github.com/some-user",
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ParseGitURL(tt.args.url)
			if tt.expectErr {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
