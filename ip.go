package ip

import (
	"net/http"
	"io/ioutil"
	"bytes"
)

// http://icanhazip.com
// http://checkip.amazonaws.com
func GetIp() (string, error) {
	rsp, err := http.Get("http://icanhazip.com/")
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(buf)), nil
}
