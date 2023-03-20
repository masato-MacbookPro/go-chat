package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// JSON形式でレスポンスを返却する
func Write(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode response by error '%#v'", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
