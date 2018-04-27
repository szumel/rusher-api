package macro

import (
	"fmt"
	"github.com/szumel/rusher-api/internal/macro"
	"github.com/szumel/rusher-api/internal/response"
	"net/http"
)

const (
	paramCode = "code"

	errParamRequired = "param %s required"
)

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}

	code := r.URL.Query().Get(paramCode)
	if code == "" {
		sendParamError(w, "code")
		return
	}

	macro, err := macro.Fetch(code)
	if err != nil {
		sendErrorMsg(w, err.Error())
		return
	}

	w.Write(macro)
}

func sendParamError(w http.ResponseWriter, param string) {
	sendErrorMsg(w, fmt.Sprintf(errParamRequired, param))
}

func sendErrorMsg(w http.ResponseWriter, msg string) {
	je, err := errorMsg(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write(je)
}

func errorMsg(msg string) ([]byte, error) {
	return response.Json(nil, msg)
}
