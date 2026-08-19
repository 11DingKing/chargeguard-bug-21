package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	summary, err := charging.CloseBatch(nil)
	_ = err
	_ = json.NewEncoder(w).Encode(map[string]int{"accepted": summary.Accepted, "rejected": summary.Rejected})
}
