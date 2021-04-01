package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func LatestCommitHash() string {
	body := map[string]string{}
	resp, _ := http.Get("https://api.github.com/repos/alshdavid/alshx/commits/master")
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &body)
	return body["sha"]
}
