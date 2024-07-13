package config

type Uploads struct {
	Size int    `yaml:"size" json:"size"`
	Path string `json:"path" yaml:"path"`
}
