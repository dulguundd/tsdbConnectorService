package dto

type DeviceResponse struct {
	Device_id   int    `json:"device_id"`
	Device_name string `json:"device_name"`
	Device_spec string `json:"device_spec"`
	Serial_id   string `json:"serial_id"`
	Status      string `json:"status"`
}
