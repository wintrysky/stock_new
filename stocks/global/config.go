package global

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ReadKey(key string, obj interface{}) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig();err != nil {
		log.Error(err)
	}

	err := v.UnmarshalKey(key, &obj)
	if err != nil {
		log.Error(err)
	}
}
