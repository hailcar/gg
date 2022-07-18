package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// unix
// time
//
func Handler(w http.ResponseWriter, r *http.Request) {
	unixTime := r.URL.Query().Get("unix")
	if unixTime != "" {
		if i, err := strconv.ParseInt(unixTime, 10, 64); err == nil {
			timeString := time.Unix(i, 0).Format("2006-01-02 15:04:05")
			fmt.Fprintf(w, timeString)
		}
	}
	lTime := r.URL.Query().Get("time")
	if lTime != "" {
		timeString, err := time.Parse("2006-01-02 15:04:05", lTime)
		if err != nil {
			fmt.Fprintf(w, strconv.Itoa(int(timeString.Unix())))
		}
	}
	tokens := r.URL.Query().Get("token")
	if tokens != "" {
		tokenGroup := strings.Split(tokens, ".")
		if len(tokenGroup) == 3 {
			targetToken := tokenGroup[2]
			if decodeBytes, err := base64.StdEncoding.DecodeString(targetToken); err == nil {
				fmt.Fprintf(w, string(decodeBytes))
			}
		}
	}
}
