package posts

import (
	"encoding/binary"
	"net/http"
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
	req.ContentLength = r.ContentLength

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var userId int32
	err = binary.Read(resp.Body, binary.LittleEndian, &userId)
	if err != nil {
		http.Error(w, "Error reading body: "+err.Error(), http.StatusInternalServerError)
		return -1, err
	}

	return userId, nil
}
