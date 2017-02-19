package twtr

type Collection struct {
	Response CollectionResponse `json:"response"`
	Object   CollectionObject   `json:"object"`
}

type CollectionResponse struct {
	Results []map[string]string `json:"results"`
	Cursors CollectionCursor    `json:"cursors"`
}

type CollectionObject struct {
	Users     map[int64]User                `json:"users"`
	Tweets    map[int64]Tweet               `json:"tweets"`
	Timelines map[string]CollectionTimeline `json:"timelines"`
}

type CollectionCursor struct {
	NextCursor     string `json:"next_cursor"`
	PreviousCursor string `json:"previous_cursor"`
}

type CollectionTimeline struct {
	Name           string `json:"name"`
	UserID         int64  `json:"user_id"`
	CollectionURL  string `json:"collection_url"`
	Description    string `json:"description"`
	URL            string `json:"url"`
	Visibility     string `json:"visibility"`
	TimelineOrder  string `json:"timeline_order"`
	CollectionType string `json:"order"`
}
