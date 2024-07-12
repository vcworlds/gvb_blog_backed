package config

type SiteInfo struct {
	CreatedAt     string `yaml:"created_at" json:"created_at" bind:"not null"`
	BeiAn         string `yaml:"bei_an" json:"bei_an" bind:"not null"`
	Title         string `yaml:"title" json:"title" bind:"not null"`
	WechatQrImage string `yaml:"wechat_qr_image" json:"wechat_qr_image" bind:"not null"`
	QqQrImage     string `yaml:"qq_qr_image" json:"qq_qr_image" bind:"not null"`
	QqNumber      string `yaml:"qq_number" json:"qq_number" bind:"not null"`
	Version       string `yaml:"version" json:"version" bind:"not null"`
	Email         string `yaml:"email" json:"email" bind:"not null"`
	Github        string `yaml:"github" json:"github" bind:"not null"`
	Gitee         string `yaml:"gitee" json:"gitee" bind:"not null"`
	Name          string `yaml:"name" json:"name" bind:"not null"`
	Job           string `yaml:"job" json:"job" bind:"not null"`
	Address       string `yaml:"address" json:"address" bind:"not null"`
	Slogan        string `yaml:"slogan" json:"slogan" bind:"not null"`
	SloganEn      string `yaml:"slogan_en" json:"slogan_en" bind:"not null"`
}
