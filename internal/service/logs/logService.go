package logs

import (
	"dataApi/internal/model/logs"
	"github.com/pkg/errors"
)

type Param struct {
	Id   int    `json:"id"`
	Date string `json:"date"`
	Data string `json:"data"`
	Ip   string `json:"ip"`
}

func SaveLog(data Param) (out bool, err error) {
	out, err = logs.InsertLog(data.Ip)
	if err != nil {
		return out, errors.Wrap(err, "logService save failed")
	}
	return true, nil
}
