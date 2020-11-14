package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func Get(content string) string {
	var r Response
	var d map[string]interface{}
	resp, err := http.PostForm("https://api.ownthink.com/bot", url.Values{"spoken": {content}, "appid": {"52a651c8c4e3d12c2cf3c7cf265960bd"}, "userid": {"15917977183"}})
	if err != nil {
		log.Println(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	err = json.Unmarshal([]byte(buf.String()), &r)
	if err != nil {
		log.Println(err)
	}

	data, err := json.Marshal(r.Data["info"])
	if err != nil {

	}
	json.Unmarshal([]byte(data), &d)
	return fmt.Sprintln(d["text"])
}
