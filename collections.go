package twtr

// Collection represent single collection structure.
// https://dev.twitter.com/rest/collections
type Collection struct {
	Response map[string]string `json:"response"`
	Object   CollectionObject  `json:"object"`
}

// Collections represent collections structure.
// https://dev.twitter.com/rest/collections
type Collections struct {
	Response struct {
		Results []map[string]string `json:"results"`
		Cursors struct {
			NextCursor     string `json:"next_cursor"`
			PreviousCursor string `json:"previous_cursor"`
		} `json:"cursors"`
	} `json:"response"`
	Object CollectionObject `json:"object"`
}

// CollectionEntry represent collection entry structure.
// https://dev.twitter.com/rest/collections
type CollectionEntry struct {
	Response struct {
		Timeline []struct {
			Tweet struct {
				ID        int64 `json:"id"`
				SortIndex int64 `json:"sort_index"`
			} `json:"tweet"`
			FeatureContext string `json:"feature_context"`
		} `json:"timeline"`
	} `json:"response"`
	Object CollectionObject `json:"object"`
}

// CollectionObject contains Twitter object used by collections.
// https://dev.twitter.com/rest/collections
type CollectionObject struct {
	Users     map[int64]User                `json:"users"`
	Tweets    map[int64]Tweet               `json:"tweets"`
	Timelines map[string]CollectionTimeline `json:"timelines"`
}

// CollectionTimeline represent timeline structure.
// https://dev.twitter.com/rest/collections
type CollectionTimeline struct {
	Name                 string `json:"name"`
	UserID               int64  `json:"user_id"`
	CollectionURL        string `json:"collection_url"`
	CustomTimelineURL    string `json:"custom_timeline_url"`
	Description          string `json:"description"`
	URL                  string `json:"url"`
	Visibility           string `json:"visibility"`
	TimelineOrder        string `json:"timeline_order"`
	CollectionType       string `json:"order"`
	CustomCollectionType string `json:"custom_collection_type"`
}
