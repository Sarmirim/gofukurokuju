package reddit

// Source - (best?) quality
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
	Subreddit              string
	Title                  string
	Author                 string
	Ups                    int
	Thumbnail              string
	Preview                Preview
	Url_overridden_by_dest string
	URL                    string
	Created_utc            float64
	// time.Unix(Created_utc float64, 0)
}

// Children - "Parent" for data
type Children struct {
	Kind string
	Data Data
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

// type Epoch int64

// func (t Epoch) MarshalJSON() ([]byte, error) {
//     strDate := time.Time(time.Unix(int64(t), 0)).Format(time.RFC3339)
//     out := []byte(`"` + strDate + `"`)
//     return out, nil
// }

// func (t *Epoch) UnmarshalJSON(b []byte) (err error) {
//     s := strings.Trim(string(b), "\"")
//     q, err := time.Parse(time.RFC3339, s)
//     if err != nil {
//         return err
//     }
//     *t = Epoch(q.Unix())
//     return nil
// }
