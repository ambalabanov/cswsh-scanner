package cswsh

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"
)

const GUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

func computeAcceptKey(challengeKey string) string {
	h := sha1.New()
	h.Write([]byte(challengeKey + GUID))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func generateChallengeKey() (string, error) {
	p := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, p); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(p), nil
}
