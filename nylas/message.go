package nylas

// Structre that represents the Schema of the Nylas Get Message Response
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
