package config

import (
	"github.com/OpenTreeMap/otm-ecoservice/eco"
)

type Config struct {
	Database eco.DBInfo
	Data     struct {
		Path string
	}
	Server struct {
		Host string
		Port string
	}
}


func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
