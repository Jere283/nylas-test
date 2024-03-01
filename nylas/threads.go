package nylas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type UpdateBody struct {
	Unread  bool     `json:"unread"`
	Starred bool     `json:"starred"`
	Folders []string `json:"folders"`
}

func GetThread(threadId, grantId string) ThreadResponse {
	requestURL := fmt.Sprintf("https://api.us.nylas.com/v3/grants/%s/threads/%s", grantId, threadId)
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

	var threadResponse ThreadResponse
	if err := json.NewDecoder(res.Body).Decode(&threadResponse); err != nil {
		log.Fatal(err)
	}

	decodedBody := strings.ReplaceAll(threadResponse.Data.LatestDraftOrMessage.Body, `\u003c`, "<")
	decodedBody = strings.ReplaceAll(decodedBody, `\u003e`, ">")
	threadResponse.Data.LatestDraftOrMessage.Body = decodedBody

	return threadResponse
}

func UpdateThread(threadId, grantId string, reqBody UpdateBody) {
	requestURL := fmt.Sprintf("https://api.us.nylas.com/v3/grants/%s/threads/%s", grantId, threadId)
	apiKey := os.Getenv("API_KEY")

	payload, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("PUT", requestURL, bytes.NewBuffer(payload))
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	log.Println(res.StatusCode)
}
