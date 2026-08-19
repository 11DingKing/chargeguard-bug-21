package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"errors"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	summary, err := charging.CloseBatch(nil)
	if errors.Is(err, charging.ErrEmptyClosureBatch) {
		http.Error(w, "closure batch cannot be empty", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if summary == nil {
		http.Error(w, "summary missing", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]int{"accepted": summary.Accepted, "rejected": summary.Rejected})
}
