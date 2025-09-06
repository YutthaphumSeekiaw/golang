package externalapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PushNotiRequest struct {
	Message string `json:"message"`
}

type PushNotiResponse struct {
	Status string `json:"status"`
}

func CallPushNoti(apiURL, message string) (*PushNotiResponse, error) {
	body, _ := json.Marshal(PushNotiRequest{Message: message})
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result PushNotiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
