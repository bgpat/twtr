package twtr

type Cursor struct {
	PreviousCursor    int64  `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
}
