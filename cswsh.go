package cswsh

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type polRes struct {
	Sid          string   `json:"sid"`
	Upgrades     []string `json:"upgrades"`
	PingInterval int      `json:"pingInterval"`
	PingTimeout  int      `json:"pingTimeout"`
}
// set configure fo scan
type Config struct {
	Socket  bool
	Verbose bool
	Origin  string
}

var (
	errMalformedURL = errors.New("malformed ws or wss URL")
	errGettingSID   = errors.New("error getting SID")
)
// main function for package
func Scan(urlWs string, c Config) (bool, error) {
	urlHTTP, err := url.Parse(urlWs)
	if err != nil {
		return false, errMalformedURL
	}
	switch urlHTTP.Scheme {
	case "ws":
		urlHTTP.Scheme = "http"
	case "wss":
		urlHTTP.Scheme = "https"
	default:
		return false, errMalformedURL
	}
	if c.Socket {
		err := getSID(urlHTTP)
		if err != nil {
			return false, errGettingSID
		}
	}
	challengeKey, err := generateChallengeKey()
	if err != nil {
		return false, err
	}
	req := &http.Request{
		Method:     "GET",
		URL:        urlHTTP,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       urlHTTP.Host,
	}
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Sec-WebSocket-Key", challengeKey)
	req.Header.Set("Sec-WebSocket-Version", "13")
	req.Header.Set("Origin", c.Origin)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	if c.Verbose {
		dumpReq, err := httputil.DumpRequest(req, false)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(dumpReq))
		dumpRes, err := httputil.DumpResponse(res, false)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(dumpRes))
	}
	return res.StatusCode == 101, nil
}

func getSID(urlHTTP *url.URL) error {
	query := make(url.Values)
	query.Add("EIO", "3")
	query.Add("transport", "polling")
	urlHTTP.RawQuery = query.Encode()
	r, err := http.Get(urlHTTP.String())
	if err != nil {
		return err
	}
	response := new(polRes)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body[4:len(body)-4], response); err != nil {
		return err
	}
	query.Set("transport", "websocket")
	query.Add("sid", response.Sid)
	urlHTTP.RawQuery = query.Encode()
	return nil
}
