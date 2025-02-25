package util

import (
	"encoding/json"
	"github.com/Whiteblock/go-prettyjson"
	"os"
	"strings"
)

func Prettyp(s string) string {
	_, noPretty := os.LookupEnv("NO_PRETTY")
	if noPretty {
		return s
	}
	s = strings.Trim(s, "\n\t\r\v ")
	if s[0] == '"' {
		return s
	}
	var tmp interface{}
	err := json.Unmarshal([]byte(s), &tmp)
	if err != nil {
		return s
	}
	pps, _ := prettyjson.Marshal(tmp)
	return string(pps)
}

func Prettypi(i interface{}) string {
	_, noPretty := os.LookupEnv("NO_PRETTY")
	if noPretty {
		out, _ := json.Marshal(i)
		return string(out)
	}
	out, _ := prettyjson.Marshal(i)
	return string(out)
}
