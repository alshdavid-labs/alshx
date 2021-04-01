package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LatestCommitHash() string {
	body := map[string]string{}
	resp, err := http.Get("https://api.github.com/repos/alshdavid/alshx/commits/master")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(resp)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &body)
	return body["sha"]
}
