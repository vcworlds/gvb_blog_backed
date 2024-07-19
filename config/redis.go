package config

import "fmt"

type Redis struct {
	Ip       string `yaml:"ip" json:"ip"`
	Port     int    `json:"addr" yaml:"addr"`
	Password string `yaml:"password" json:"password"`
	PoolSize int    `json:"pool_size" yaml:"pool_size"`
}

func (receiver Redis) Addr() string {
	return fmt.Sprintf("%s:%d", receiver.Ip, receiver.Port)
}
