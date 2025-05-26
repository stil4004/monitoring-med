package logger

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
)

func Zap() (*zap.Logger, error) {
	jsonFile, err := os.Open("./config/zap_logger.json")
	if err != nil {
		return &zap.Logger{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cfg zap.Config
	json.Unmarshal([]byte(byteValue), &cfg)

	return zap.Must(cfg.Build()), nil
}
