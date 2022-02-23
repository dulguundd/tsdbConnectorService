package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"tsdbConnectorService/domain/device"
	"tsdbConnectorService/dto"
)

type DeviceService interface {
	GetAllDevice() ([]dto.DeviceResponse, *errs.AppError)
	GetDevice(int) (*dto.DeviceResponse, *errs.AppError)
}

type DefaultDeviceService struct {
	repo device.DeviceRepository
}

func (s DefaultDeviceService) GetAllDevice() ([]dto.DeviceResponse, *errs.AppError) {
	d, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var response []dto.DeviceResponse
	for i := range d {
		response = append(response, d[i].ToDtoDevice())
	}
	return response, nil
}

func (s DefaultDeviceService) GetDevice(id int) (*dto.DeviceResponse, *errs.AppError) {
	d, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := d.ToDtoDevice()
	return &response, nil
}

func NewDeviceService(repository device.DeviceRepository) DefaultDeviceService {
	return DefaultDeviceService{(repository)}
}
