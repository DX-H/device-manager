// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type GenerateDemoDeviceDataReq struct {
	DeviceNumber int   `form:"device_number"`
	StartTime    int64 `form:"start_time"`
	EndTime      int64 `form:"end_time"`
}

type GenerateDemoDeviceDataResp struct {
	Code int `json:"code"`
}

type GenerateDemoMetadataReq struct {
	DeviceNumber      int `form:"device_number"`
	DeviceParamNumber int `form:"device_param_number"`
}

type GenerateDemoMetadataResp struct {
	Code int `json:"code"`
}

type GetMetadataRequest struct {
	SN string `form:"sn"`
}

type GetMetadataResponse struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}

type GetMetricsRequest struct {
	SN        string `form:"sn"`
	StartTime int64  `form:"start_time"`
	EndTime   int64  `form:"end_time"`
}

type GetMetricsResponse struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}

type GetUpdateResultReq struct {
	JobId int `form:"job_id"`
}

type GetUpdateResultResp struct {
	End                        bool  `json:"end"`
	DeviceNumber               int   `json:"device_number"`
	DeviceParamNumber          int   `json:"device_param_number"`
	Thread                     int   `json:"thread"`
	Seconds                    int   `json:"seconds"`
	SuccessDeviceCount         int   `json:"success_device_count"`
	AverageLatencyMicroseconds int64 `json:"average_latency_microseconds"`
}

type UpdateMetadataReq struct {
	DeviceNumber      int `json:"device_number"`
	DeviceParamNumber int `json:"device_param_number"`
	Thread            int `json:"thread"`
	Seconds           int `json:"seconds"`
}

type UpdateMetadataRequest struct {
	SN     string            `json:"sn"`
	Params map[string]string `json:"params"`
}

type UpdateMetadataResp struct {
	JobId int `json:"job_id"`
}

type UpdateMetadataResponse struct {
	Code int `json:"code"`
}
