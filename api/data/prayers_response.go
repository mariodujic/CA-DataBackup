package data

type Prayer struct {
	ItemID        string        `json:"itemId"`
	MusicURL      string        `json:"musicUrl"`
	Tag           string        `json:"tag"`
	Text          string        `json:"text"`
	Title         string        `json:"title"`
	Type          int           `json:"type"`
	SegmentedList []interface{} `json:"segmentedList"`
}

type PrayersResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Prayer `json:"data"`
}
