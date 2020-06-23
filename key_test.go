package cswsh

import (
	"regexp"
	"testing"
)

func Test_computeAcceptKey(t *testing.T) {
	type args struct {
		challengeKey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"random", args{"sh7Xipy6UioJq++T1WHfig=="}, "3XSEFP8sORYSSbaF+JXIo/eYgxY="},
		{"random", args{"I6q5fZ9pM4eDHaylVUCikQ=="}, "xmzOkHnfiC8w5CHPq6wKlPmpJUs="},
		{"random", args{"8q+09ByuIZ+FDyOieCPx6w=="}, "bGg0xX+Ev/znz7PKBXkod00S9JA="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeAcceptKey(tt.args.challengeKey); got != tt.want {
				t.Errorf("computeAcceptKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateChallengeKey(t *testing.T) {
	re := regexp.MustCompile(`^[a-zA-Z\d/+=]{24}$`)
	tests := []struct {
		name string
		//want    string
		wantErr bool
	}{
		{"RegExp Base64", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateChallengeKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("generateChallengeKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !re.MatchString(got) {
				t.Errorf("generateChallengeKey() got = %v, want Base64", got)
			}
		})
	}
}
