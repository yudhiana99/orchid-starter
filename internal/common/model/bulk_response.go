package model

type BulkResponse struct {
	Batches           int                   `json:"batches"`
	Deleted           int                   `json:"deleted"`
	Failures          []bulkResponseFailure `json:"failures"`
	Noops             int                   `json:"noops"`
	RequestsPerSecond float32               `json:"requests_per_second"`
	Retries           struct {
		Bulk   int `json:"bulk"`
		Search int `json:"search"`
	} `json:"retries"`
	ThrottledMillis      int  `json:"throttled_millis"`
	ThrottledUntilMillis int  `json:"throttled_until_millis"`
	TimedOut             bool `json:"timed_out"`
	Took                 int  `json:"took"`
	Total                int  `json:"total"`
	Updated              int  `json:"updated"`
	VersionConflicts     int  `json:"version_conflicts"`
}

type bulkResponseFailure struct {
	Index  string `json:"index,omitempty"`
	Type   string `json:"type,omitempty"`
	Id     string `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
	Shard  int    `json:"shard,omitempty"`
	Node   int    `json:"node,omitempty"`
}
