package git

import (
	"reflect"
	"testing"
)

func TestParseGitURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want Remote
	}{
		{
			name: "git@",
			args: args{
				url: "git@github.com:thrawny/openci.git",
			},
			want: Remote{
				Domain:  "github.com",
				Org:     "thrawny",
				Project: "openci",
			},
		},
		{
			name: "https",
			args: args{
				url: "https://github.com/some-user/my-repo.git",
			},
			want: Remote{
				Domain:  "github.com",
				Org:     "some-user",
				Project: "my-repo",
			},
		},
		{
			name: "https without .git",
			args: args{
				url: "https://github.com/some-user/my-repo",
			},
			want: Remote{
				Domain:  "github.com",
				Org:     "some-user",
				Project: "my-repo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseGitURL(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGitURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
