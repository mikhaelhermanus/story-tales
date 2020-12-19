package logger

import (
	"context"
	"encoding/json"
	"fmt"
)

type device struct {
	Phone       string `json:"phone"`
	DeviceType  string `json:"device_type"`
	DeviceOS    string `json:"device_os"`
	DeviceBrand string `json:"device_brand"`
	DeviceModel string `json:"device_model"`
}

// Device ...
func Device(ctx context.Context, data []byte) context.Context {
	dev := device{}
	json.Unmarshal(data, &dev)

	return RecordDevice(ctx, dev.Phone, dev.DeviceType, dev.DeviceOS, dev.DeviceBrand, dev.DeviceModel)
}

// RecordDevice ...
func RecordDevice(ctx context.Context, msisdn, deviceType, deviceOS, brand, model string) context.Context {
	v, ok := ctx.Value(logKey).(*Data)
	if ok {
		v.UserCode = msisdn
		v.Device = fmt.Sprintf("%s %s, %s %s", deviceType, deviceOS, brand, model)

		ctx = context.WithValue(ctx, logKey, v)

		return ctx
	}

	return ctx
}
