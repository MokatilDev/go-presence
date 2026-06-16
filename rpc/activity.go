package rpc

type (
	ActivityType      int
	StatusDisplayType int
)

const (
	Playing ActivityType = iota
	Streaming
	Listening
	Watching
	Custom
	Competing
)

const (
	Name StatusDisplayType = iota
	State
	Details
)

type Activity struct {
	Type          ActivityType        `json:"type"`
	StatusDisplay StatusDisplayType   `json:"status_display_type"`
	Details       *string             `json:"details,omitempty"`
	DetailsURL    *string             `json:"details_url,omitempty"`
	State         *string             `json:"state,omitempty"`
	StateURL      *string             `json:"state_url,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	Buttons       *[]ActivityButton   `json:"buttons,omitempty"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	LargeURL   string `json:"large_url"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
	SmallURL   string `json:"small_url"`
}

type ActivityTimestamps struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

type ActivityButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}
