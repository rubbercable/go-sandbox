package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.Readall(r.Body); err == nill {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
