package leapapi

// /v2/health format
type Health struct {
	Version        string                 `json:"version"`
	VersionHash    string                 `json:"version_hash"`
	Health         []Service              `json:"health"`
	Features       map[string]interface{} `json:"features"`
	QueryTime      float32                `json:"query_time_ms"`
	HTTPStatusCode int
}
