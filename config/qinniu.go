package config

type QiNiu struct {
	Enable    bool    `yaml:"enable"` // 是否启用七牛
	AccessKey string  `yaml:"access_key"`
	SecretKey string  `yaml:"secret_key"`
	Bucket    string  `yaml:"bucket"` // 存储桶的名字
	CDN       string  `yaml:"cdn"`    // 访问图片地址的前缀
	Zone      string  `yaml:"zone"`   // 存储的地区
	Size      float64 `yaml:"size"`   // 存储的大小限制，单位是MB
}
