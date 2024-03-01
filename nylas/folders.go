package nylas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type FolderResponse struct {
	RequestId  string `json:"request_id"`
	NextCursor string `json:"next_cursor"`
	Data       []struct {
		BackgroundColor string   `json:"background_color"`
		ChildCount      int      `json:"child_count,omitempty"`
		GrantId         string   `json:"grant_id,omitempty"`
		Id              string   `json:"id"`
		Name            string   `json:"name"`
		Object          string   `json:"object"`
		ParentId        string   `json:"parent_id,omitempty"`
		SystemFolder    bool     `json:"system_folder,omitempty"`
		TextColor       string   `json:"text_color,omitempty"`
		TotalCount      int      `json:"total_count,omitempty"`
		UnreadCount     int      `json:"unread_count,omitempty"`
		Attributes      []string `json:"attributes,omitempty"`
	} `json:"data"`
}

func GetFolders(grantId string) FolderResponse {
	requestURL := fmt.Sprintf("https://api.us.nylas.com/v3/grants/%s/folders/", grantId)
	apiKey := os.Getenv("API_KEY")

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Println(res.StatusCode)

	var folders FolderResponse
	if err := json.NewDecoder(res.Body).Decode(&folders); err != nil {
		log.Fatal(err)
	}

	return folders
}
