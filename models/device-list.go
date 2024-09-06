package models

type DeviceGroupList struct {
	DeviceList       []DeviceItem `json:"device_list"`
	GroupID          int          `json:"group_id"`
	GroupName        string       `json:"group_name"`
	GroupTypeID      int          `json:"group_type_id"`
	GroupTypeLogoURL string       `json:"group_type_logo_url"`
	LogoURL          string       `json:"logo_url"`
}

type DeviceItem struct {
	DeviceMac    string `json:"device_mac"`
	ProductModel string `json:"product_model"`
}

type Device struct {
	BindingTS           int          `json:"binding_ts"`
	BindingUserNickname string       `json:"binding_user_nickname"`
	ConnState           int          `json:"conn_state"`
	ConnStateTS         int          `json:"conn_state_ts"`
	DeviceParams        DeviceParams `json:"device_params"`
	ENR                 string       `json:"enr"`
	FirmwareVer         string       `json:"firmware_ver"`
	HardwareVer         string       `json:"hardware_ver"`
	MAC                 string       `json:"mac"`
	Nickname            string       `json:"nickname"`
	ProductModel        string       `json:"product_model"`
	ProductType         string       `json:"product_type"`
	PushSwitch          int          `json:"push_switch"`
	TimezoneName        string       `json:"timezone_name"`
}

type DeviceParams struct {
	AccessorySwitch       int    `json:"accessory_switch"`
	AINotificationV2      int    `json:"ai_notification_v2"`
	AudioAlarmSwitch      int    `json:"audio_alarm_switch"`
	BatteryChargingStatus string `json:"battery_charging_status"`
	IP                    string `json:"ip"`
	PublicIP              string `json:"public_ip"`
	SSID                  string `json:"ssid"`
}
