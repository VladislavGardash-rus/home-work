package cfg

import "github.com/jinzhu/configor"

var _config = new(Cfg)

type Cfg struct {
	Logger             LoggerConf         `json:"logger"`
	CalendarHttpServer CalendarHttpServer `json:"calendarHttpServer"`
	Storage            Storage            `json:"storage"`
}

type CalendarHttpServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type LoggerConf struct {
	Level string `json:"level"`
}

type Storage struct {
	Type       string `json:"type"` //memory || database
	Connection string `json:"connection"`
}

func InitConfig(configFile string) error {
	err := configor.Load(_config, configFile)
	if err != nil {
		return err
	}

	return nil
}

func Config() *Cfg {
	return _config
}
