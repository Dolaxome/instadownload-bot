package storylink

var config = map[string]string{
	"ds_user_id": " ",
	"sessionid":  " ",
	"csrftoken":  " ",
	"User-Agent": "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+",
}

func SetDS(s string) {
	config["ds_user_id"] = s
}
func SetSESS(s string) {
	config["sessionid"] = s
}
func SetCSRF(s string) {
	config["csrftoken"] = s
}
