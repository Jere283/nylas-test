package nylas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Structre that represents the Schema of the Nylas Get Message Response
type MessageResponse struct {
	RequestID string  `json:"request_id"`
	Data      Message `json:"data"`
}

type Message struct {
	Bcc         []NameEmailPair `json:"bcc"`
	Body        string          `json:"body"`
	Cc          []NameEmailPair `json:"cc"`
	Date        int             `json:"date"`
	Attachments []struct {
		ContentType string `json:"date"`
		FileName    string `json:"filename"`
		GrantId     string `json:"grant_id"`
		Id          string `json:"id"`
		IsInine     bool   `json:"is_inline"`
		Size        int    `json:"size"`
	} `json:"attachments"`
	Folders    []string        `json:"folders"`
	From       []NameEmailPair `json:"from"`
	GrantID    string          `json:"grant_id"`
	ID         string          `json:"id"`
	Object     string          `json:"object"`
	ReplyTo    []NameEmailPair `json:"reply_to"`
	ScheduleID string          `json:"schedule_id"`
	Snippet    string          `json:"snippet"`
	Starred    bool            `json:"starred"`
	Subject    string          `json:"subject"`
	ThreadID   string          `json:"thread_id"`
	To         []NameEmailPair `json:"to"`
	Unread     bool            `json:"unread"`
	UseDraft   bool            `json:"use_draft"`
}

type NameEmailPair struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetEmail(messageId, grantId string) MessageResponse {
	requestURL := fmt.Sprintf("https://api.us.nylas.com/v3/grants/%s/messages/%s", grantId, messageId)
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

	var message MessageResponse
	if err := json.NewDecoder(res.Body).Decode(&message); err != nil {
		log.Fatal(err)
	}

	decodedBody := strings.ReplaceAll(message.Data.Body, `\u003c`, "<")
	decodedBody = strings.ReplaceAll(decodedBody, `\u003e`, ">")
	message.Data.Body = decodedBody

	return message
}
