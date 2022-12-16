package gogen

import (
	"gogen-transaction/shared/util"
	"time"
)

type Runner interface {
	Run() error
}

type ApplicationData struct {
	AppName       string `json:"appName"`
	AppInstanceID string `json:"appInstanceID"`
	StartTime     string `json:"startTime"`
}

func NewApplicationData(appName string) ApplicationData {
	appInstanceID := util.GenerateID(4)
	return ApplicationData{
		AppName:       appName,
		AppInstanceID: appInstanceID,
		StartTime:     time.Now().Format("2006-01-02 15:04:05"),
	}
}
