package leapapi

import "time"

// Service struct from /v2/health
type Service struct {
	Name   string
	Status string
	Data   map[string]interface{}
	Time   time.Time
}

func (s *Service) UnmarshalJSON(b []byte) error {
	var r struct {
		N string                 `json:"service"`
		S string                 `json:"status"`
		D map[string]interface{} `json:"service_data"`
		T int64                  `json:"time"`
	}

	err := json.Unmarshal(b, &r)
	if err == nil {
		s.Name = r.N
		s.Status = r.S
		s.Data = r.D
		s.Time = fromTS(r.T)
	}
	return err
}
