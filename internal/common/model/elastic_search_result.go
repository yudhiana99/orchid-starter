package model

type ElasticsearchResult struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      uint64 `json:"total"`
		Successful int    `json:"successful"`
		Skipped    int    `json:"skipped"`
		Failed     int    `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    uint64 `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source any     `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
