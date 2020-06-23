package cswsh

import (
	"testing"
)

func TestScan(t *testing.T) {
	type test struct {
		name    string
		urlWs   string
		want    bool
		wantErr bool
		Config
	}
	var origin = "http://hacker.com"
	var tests = []test{
		{"ws", "ws://echo.websocket.org", true, false, Config{Origin: origin, Verbose: true, Socket: false}},
		{"wss", "wss://echo.websocket.org", true, false, Config{Origin: origin, Verbose: true, Socket: false}},
		{"http", "http://echo.websocket.org", false, true, Config{Origin: origin, Verbose: true, Socket: false}},
		{"wss", "wss://websocket.org", false, false, Config{Origin: origin, Verbose: true, Socket: false}},
		{"fake", "wss://fake.org", false, true, Config{Origin: origin, Verbose: true, Socket: false}},
		{"fake", "wss://fake.org", false, true, Config{Origin: origin, Verbose: true, Socket: true}},
		{"socket", "wss://juice-shop.herokuapp.com/socket.io/", true, false, Config{Origin: origin, Verbose: true, Socket: true}},
		{"socket", "wss://juice-shop.herokuapp.com/", false, true, Config{Origin: origin, Verbose: true, Socket: true}},
		{"empty", "", false, true, Config{Origin: origin, Verbose: true, Socket: false}},
		{"malformed", "::", false, true, Config{Origin: origin, Verbose: true, Socket: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Scan(tt.urlWs, tt.Config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error: %v, wantErr: %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Scan() try: %v, got: %v, want: %v", tt.urlWs, got, tt.want)
			}
		})
	}
}
