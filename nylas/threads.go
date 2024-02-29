package nylas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Structre that represents the Schema of the Nylas Get Thread Response
type ThreadResponse struct {
	RequestID string `json:"request_id"`
	Data      struct {
		GrantID                   string          `json:"grant_id"`
		ID                        string          `json:"id"`
		Object                    string          `json:"object"`
		LatestDraftOrMessage      Message         `json:"latest_draft_or_message"`
		HasAttachments            bool            `json:"has_attachments"`
		HasDrafts                 bool            `json:"has_drafts"`
		EarliestMessageDate       int             `json:"earliest_message_date"`
		LatestMessageReceivedDate int             `json:"latest_message_received_date"`
		LatestMessageSentDate     int             `json:"latest_message_sent_date"`
		Participants              []NameEmailPair `json:"participants"`
		Snippet                   string          `json:"snippet"`
		Starred                   bool            `json:"starred"`
		Subject                   string          `json:"subject"`
		Unread                    bool            `json:"unread"`
		MessageIDs                []string        `json:"message_ids"`
		DraftIDs                  []string        `json:"draft_ids"`
		Folders                   []string        `json:"folders"`
	} `json:"data"`
}

func GetThread(messageId, grantId string) ThreadResponse {
	requestURL := fmt.Sprintf("https://api.us.nylas.com/v3/grants/%s/threads/%s", grantId, messageId)
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

	var threadResponse ThreadResponse
	if err := json.NewDecoder(res.Body).Decode(&threadResponse); err != nil {
		log.Fatal(err)
	}

	decodedBody := strings.ReplaceAll(threadResponse.Data.LatestDraftOrMessage.Body, `\u003c`, "<")
	decodedBody = strings.ReplaceAll(decodedBody, `\u003e`, ">")
	threadResponse.Data.LatestDraftOrMessage.Body = decodedBody

	return threadResponse
}
