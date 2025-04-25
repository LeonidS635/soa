package posts

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/helpers"
)

func (h *GateWayPostsHandlers) getUserId(w http.ResponseWriter, r *http.Request) (int32, error) {
	req, err := http.NewRequest("GET", h.userServiceURL+"/user_id", nil)
	if err != nil {
		http.Error(w, "Error creating request: "+err.Error(), http.StatusInternalServerError)
		return -1, err
	}

	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
		return -1, errors.New("error getting user id")
	}

	data, err := helpers.ReadBodyFromResponse(resp)
	if err != nil {
		http.Error(w, "get user id: error reading response body: "+err.Error(), http.StatusInternalServerError)
		return -1, err
	}

	responseStruct := struct {
		UserId int `json:"user_id"`
	}{}
	err = json.Unmarshal(data, &responseStruct)
	if err != nil {
		http.Error(w, "get user id: error unmarshalling response body: "+err.Error(), http.StatusInternalServerError)
	}

	return int32(responseStruct.UserId), nil
}
