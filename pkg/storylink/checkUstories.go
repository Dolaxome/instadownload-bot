package storylink

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const LinkToStories = `https://i.instagram.com/api/v1/feed/user/{{USERID}}/reel_media/`

func GetUStoriesTray(id int64) (tray Tray, err error) {
	url := strings.Replace(LinkToStories, "{{USERID}}", strconv.FormatInt(id, 10), 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: config["ds_user_id"]})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: config["sessionid"]})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: config["csrftoken"]})

	req.Header.Set("User-Agent", config["User-Agent"])

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("Code:" + strconv.Itoa(resp.StatusCode))
		return
	}

	decode := json.NewDecoder(resp.Body)
	err = decode.Decode(&tray)
	return
}
