package device

import (
	"github.com/dulguundd/logError-lib/errs"
	"tsdbConnectorService/dto"
)

type Device struct {
	Device_id   int
	Device_name string
	Device_spec string
	Serial_id   string
	Status      int
}

func (d Device) statusAsText() string {
	statusAsText := "active"
	if d.Status == 0 {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (d Device) ToDtoDevice() dto.DeviceResponse {
	return dto.DeviceResponse{
		Device_id:   d.Device_id,
		Device_name: d.Device_name,
		Device_spec: d.Device_spec,
		Serial_id:   d.Serial_id,
		Status:      d.statusAsText(),
	}
}

type DeviceRepository interface {
	FindAll() ([]Device, *errs.AppError)
	ById(id int) (*Device, *errs.AppError)
}
