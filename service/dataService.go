package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"time"
	"tsdbConnectorService/domain/data"
	"tsdbConnectorService/dto"
)

type DataService interface {
	GetTodayClearData() ([]dto.DataResponse, *errs.AppError)
	GetTodayClearDataOfDevice(id int) ([]dto.DataResponse, *errs.AppError)
	PostData(req dto.NewDataRequest) (*dto.DataResponse, *errs.AppError)
}

type DefaultDataService struct {
	repo data.DataRepository
}

func (s DefaultDataService) GetTodayClearData() ([]dto.DataResponse, *errs.AppError) {
	d, err := s.repo.GetTodayClearData()
	if err != nil {
		return nil, err
	}
	var response []dto.DataResponse
	for i := range d {
		response = append(response, d[i].ToDtoData())
	}
	return response, nil
}
func (s DefaultDataService) GetTodayClearDataOfDevice(id int) ([]dto.DataResponse, *errs.AppError) {
	d, err := s.repo.GetTodayClearDataOfDevice(id)
	if err != nil {
		return nil, err
	}
	var response []dto.DataResponse
	for i := range d {
		response = append(response, d[i].ToDtoData())
	}
	return response, nil
}

func (s DefaultDataService) PostData(req dto.NewDataRequest) (*dto.DataResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	data := data.Data{
		Time:      time.Now(),
		Temp:      req.Temp,
		Device_id: req.Device_id,
	}
	d, err := s.repo.PostData(data)
	if err != nil {
		return nil, err
	}
	response := d.ToDtoData()
	return &response, nil
}

func NewDataService(repository data.DataRepository) DefaultDataService {
	return DefaultDataService{(repository)}
}
