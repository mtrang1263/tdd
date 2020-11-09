package v1

import (
	"net/http"
	"time"
)

var SomeAnonymousStruct = struct {
	someField   string
	someCounter int
}{
	"THIS_IS_COOL",
	-1,
}

func GetCurrentDate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Only GET requests are supported"))
		return
	}

	timeFormat := time.RFC3339

	if f := r.URL.Query().Get("format"); f != "" {
		switch f {
		case time.RFC3339Nano:
			timeFormat = time.RFC3339Nano
		case time.RFC1123:
			timeFormat = time.RFC1123
		case "MARTIN_COOL_FORMAT":
			timeFormat = "02-2006-01" // DD-YYYY-MM
		}
	}

	w.Write([]byte(time.Now().Format(timeFormat)))
}
