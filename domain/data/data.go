package data

import (
	"github.com/dulguundd/logError-lib/errs"
	"time"
	"tsdbConnectorService/dto"
)

type Data struct {
	Time      time.Time
	Temp      float64
	Device_id int
}

type DataRepository interface {
	GetTodayClearData() ([]Data, *errs.AppError)
	GetTodayClearDataOfDevice(id int) ([]Data, *errs.AppError)
	PostData(data Data) (*Data, *errs.AppError)
}

func (d Data) ToDtoData() dto.DataResponse {
	return dto.DataResponse{
		Time:      d.Time,
		Temp:      d.Temp,
		Device_id: d.Device_id,
	}
}
