package twtr

type DirectMessage struct {
	ID
	CreatedAt           Time     `json:"created_at"`
	Entities            Entities `json:"entities"`
	Recipient           User     `json:"recipient"`
	RecipientID         int64    `json:"recipient_id"`
	RecipientScreenName string   `json:"recipient_screen_name"`
	Sender              User     `json:"sender"`
	SenderID            int64    `json:"sender_id"`
	SenderScreenName    string   `json:"sender_screen_name"`
	Text                string   `json:"text"`
}
