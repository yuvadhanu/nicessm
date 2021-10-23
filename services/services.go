package services

import (
	"nicessm-api-service/config"
	"nicessm-api-service/daos"
	"nicessm-api-service/redis"
	"nicessm-api-service/shared"
)

//Service : ""
type Service struct {
	Daos         *daos.Daos
	Shared       *shared.Shared
	Redis        *redis.RedisCli
	ConfigReader *config.ViperConfigReader
}

//IService : ""
type IService interface {
}

//GetService :""
func GetService(dao *daos.Daos, s *shared.Shared, Redis *redis.RedisCli, configReader *config.ViperConfigReader) *Service {
	return &Service{dao, s, Redis, configReader}
}
