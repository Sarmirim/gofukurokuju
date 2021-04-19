package reddit

// Reddit_video -
type Reddit_video struct {
	Bitrate_kbps string
	Fallback_url int
	Height       int
	Width        int
	Duration     int
	Is_gif       string
}

// Media -
type Media struct {
	Reddit_video Reddit_video
}

// Source - best? (atleast source) quality
type Source struct {
	URL    string
	Width  int
	Height int
}

// Images - different resolutions
type Images struct {
	Source      Source
	ID          string
	Variants    interface{}
	Resolutions interface{}
}

// Preview -
type Preview struct {
	Images  []Images
	Enabled bool
}

// Data - real data for post
type Data struct {
	Subreddit        string  `json:"subreddit"`
	Title            string  `json:"title"`
	Author           string  `json:"author"`
	Ups              int     `json:"ups"`
	Thumbnail        string  `json:"thumbnail"`
	Thumbnail_width  int     `json:"thumbnail_width"`
	Thumbnail_height int     `json:"thumbnail_height"`
	URL              string  `json:"url"`
	Created_utc      float64 `json:"created_utc"` // int64 doesn't work
	Is_video         bool    `json:"is_video"`
	// Permalink        string
	// Preview          Preview
	// Over_18          bool
	// Url_overridden_by_dest string
	// UTC                    time.Time
	// // time.Unix(Created_utc float64, 0)
}

// Children - "Parent" for data
type Children struct {
	Kind string
	Data *Data
}

// Data0 JSON root for post or comment sections
type Data0 struct {
	Modhash  string
	Dist     int
	Children []Children
	After    string
	Before   string
}

// Post JSON root
type Post struct {
	// Kind  int    `json:"kind"` ~= Kind string
	Kind string
	Data Data0
}
