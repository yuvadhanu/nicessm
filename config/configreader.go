package config

import (
	"log"

	"github.com/spf13/viper"
)

var ConfigReader *viperConfigReader

type Reader interface {
	GetString(key string) string
	GetInt(key string) int
}

type viperConfigReader struct {
	viper *viper.Viper
}

func (v viperConfigReader) GetString(key string) string {
	return v.viper.GetString(key)
}

func (v viperConfigReader) GetInt(key string) int {
	return v.viper.GetInt(key)
}

func init() {
	v := viper.New()
	v.SetConfigName("nicessmconfig")
	v.AddConfigPath("config")

	err := v.ReadInConfig()
	if err != nil {
		log.Panic("Not able to read configuration..", err.Error())
	}
	ConfigReader = &viperConfigReader{
		viper: v,
	}
}

// ViperConfigReader ..
type ViperConfigReader struct {
	viper *viper.Viper
}

// GetString ...
func (v ViperConfigReader) GetString(key string) string {
	return v.viper.GetString(key)
}

//GetInt ...
func (v ViperConfigReader) GetInt(key string) int {
	return v.viper.GetInt(key)
}

//Config ..
func Config() *ViperConfigReader {
	v := viper.New()
	v.SetConfigName("nicessmconfig")
	v.AddConfigPath("config")
	err := v.ReadInConfig()
	if err != nil {
		log.Panic("Not able to read configuration..", err.Error())
	}
	cv := &ViperConfigReader{
		viper: v,
	}
	return cv
}
