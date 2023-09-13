package config

type Local struct {
	Path       string `mapstructure:"path" json:"path" yaml:"path"`                      // 本地文件访问路径
	StorePath  string `mapstructure:"store-path" json:"store-path" yaml:"store-path"`    // 本地文件存储路径
	UploadPath string `mapstructure:"upload-path" json:"upload-path" yaml:"upload-path"` // 文件存储上传路径
}
