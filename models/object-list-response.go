package models

type ObjectListResponse struct {
	Code    string         `json:"code"`
	TS      int            `json:"ts"`
	MSG     string         `json:"msg"`
	TraceID string         `json:"traceId"`
	Data    ObjectListData `json:"data"`
}

type ObjectListData struct {
	DeviceGroupList []DeviceGroupList `json:"device_group_list"`
	DeviceList      []Device          `json:"device_list"`
}
