package storylink

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const LinkToAllStories = `https://i.instagram.com/api/v1/feed/reels_tray/`

type TrayUser struct {
	Username string `json:"username"`
}

type TrayItemVideoVersion struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

type TrayItemImageVersion2 struct {
	Candidates []struct {
		Url string `json:"url"`
	} `json:"candidates"`
}

type TrayItem struct {
	TakenAt        int64                  `json:"taken_at"`
	ImageVersions2 TrayItemImageVersion2  `json:"image_versions2"`
	VideoVersions  []TrayItemVideoVersion `json:"video_versions"`
	//HasAudio        bool                    `json:"has_audio"`
}

type Tray struct {
	Id              int64      `json:"id"`
	LatestReelMedia int64      `json:"latest_reel_media"`
	User            TrayUser   `json:"user"`
	Items           []TrayItem `json:"items"`
}

type RawReelsTray struct {
	Trays []Tray `json:"tray"`
}

func GetReelsTray() (trays []Tray, err error) {
	req, err := http.NewRequest("GET", LinkToAllStories, nil)
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
	t := RawReelsTray{}
	if err = decode.Decode(&t); err != nil {
		return
	}
	trays = t.Trays
	return
}
