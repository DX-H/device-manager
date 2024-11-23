package svc

import (
	"finishy1995/device-manager/legacy/http/internal/config"
	"finishy1995/device-manager/legacy/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                 config.Config
	DeviceMetadataModel    model.DeviceMetadataModel
	DeviceCameraDataModel  model.DeviceCameraDataModel
	DeviceSweeperDataModel model.DeviceSweeperDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:                 c,
		DeviceMetadataModel:    model.NewDeviceMetadataModel(conn),
		DeviceCameraDataModel:  model.NewDeviceCameraDataModel(conn),
		DeviceSweeperDataModel: model.NewDeviceSweeperDataModel(conn),
	}
}
