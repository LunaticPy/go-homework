package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	Calendar Calendar `json:"calendar"`
	Server   Server   `json:"server"`
}

type Calendar struct{}

type duration time.Duration

func (d *duration) UnmarshalText(text []byte) error {
	temp, err := time.ParseDuration(string(text))
	*d = duration(temp)
	return err
}

type Server struct {
	Addr         string   `json:"Addr"`
	ReadTimeout  duration `json:"ReadTimeout"`
	WriteTimeout duration `json:"WriteTimeout"`
}

func GetConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
