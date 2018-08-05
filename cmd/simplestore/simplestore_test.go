package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	port = 0
)

func init() {
	p := flag.Int("port", 0, "The port is the appengine app running on")
	flag.Parse()
	port = *p
}

func urlFor(path string) string {
	return fmt.Sprintf("http://localhost:%d/%s", port, path)
}

func Test404(t *testing.T) {
	t.Log("Test404")
	getResp, err := http.Get(urlFor("messages/isnotasha256hash"))
	if err != nil {
		t.Error(err)
		return
	}
	if getResp.StatusCode != 404 {
		t.Error("expected a 404 from a invalid url")
	}
}
func TestSimpleStore(t *testing.T) {
	t.Log("TestSimpleStore")
	data := struct {
		Foo string
	}{"bar"}
	b, err := json.Marshal(data)

	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("sending: [%s]\n", b)
	resp, err := http.Post(urlFor("messages"), "application/json", bytes.NewReader(b))
	if err != nil {
		t.Error(err)
		return
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	// to get shasum,
	// printf '{"Foo":"bar"}' | shasum -a 256

	t.Logf("resp: %s\n", respBody)

	res := struct {
		Digest string `json:"digest"`
	}{}

	err = json.Unmarshal(respBody, &res)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Digest != "e3efd8637bdd1c1a5ee4d5ca163fa6ba8ab279b189621894a2731aefd672cdff" {
		t.Error("Got back incorrect digest")
	}

	getResp, err := http.Get(urlFor("messages/e3efd8637bdd1c1a5ee4d5ca163fa6ba8ab279b189621894a2731aefd672cdff"))
	if err != nil {
		t.Error(err)
		return
	}
	getRespBody, err := ioutil.ReadAll(getResp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if fmt.Sprintf("%s", getRespBody) != `{"Foo":"bar"}` {
		t.Error("did not get original message back from endpoint!")
		return
	}
	t.Logf("getRespBody\n%s\n", getRespBody)
}
