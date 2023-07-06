package storydw

import (
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

func StripQueryString(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	return u.String()
}

func FileExt(url string) string {
	filename := path.Base(StripQueryString(url))
	return path.Ext(filename)
}

func FormatTimestamp() string {
	// r := time.Unix(timestamp, 0)
	// fmt.Println(r)
	t := time.Now()
	p := strconv.Itoa(t.Day()) + "." + strconv.Itoa(int(t.Month())) + "." + strconv.Itoa(int(t.Year()))
	return p + " " + t.Format("15:04:05")
	// return t.Format(time.RFC3339)
}

func CreateIfNotExist(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}

func BuildPath(username, url string, counter int) string {
	dirname := path.Join("stories", username)
	CreateIfNotExist(dirname)
	ext := FileExt(url)
	ts := FormatTimestamp()
	filename := username + " " + ts[0:10] + " [" + strconv.Itoa(counter) + "]" + ext
	p := path.Join(dirname, filename)
	return p
}
