package models

type PropertyListResponse struct {
	Code    string           `json:"code"`
	TS      int              `json:"ts"`
	MSG     string           `json:"msg"`
	TraceID string           `json:"traceId"`
	Data    PropertyListData `json:"data"`
}

type PropertyListData struct {
	PropertyList []Property `json:"property_list"`
}

type Property struct {
	PID   string `json:"pid"`
	Value string `json:"value"`
	TS    int    `json:"ts"`
}
