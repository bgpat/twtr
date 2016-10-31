package twitter

type Cursor struct {
	PreviousCursor    int64  `json:"previous_cursor,omitempty"`
	PreviousCursorStr string `json:"previous_cursor_str,omitempty"`
	NextCursor        int64  `json:"next_cursor,omitempty"`
	NextCursorStr     string `json:"next_cursor_str,omitempty"`
}
