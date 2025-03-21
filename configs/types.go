package configs

type Configs struct {
	Platform platform `json:"platform" yaml:"platform"`
	Database database `json:"database" yaml:"database"`
	OAuth    oauth    `json:"oauth" yaml:"oauth"`
	Admin    user     `json:"admin" yaml:"admin"`
}

type database struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     int64  `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
}

type oauth map[string]interface{}
type platform struct {
	Port    uint64 `json:"port" yaml:"port"`
	BaseURL string `json:"base_url" yaml:"base_url"`
}
type user struct {
	Username string `json:"username" yaml:"username"`
	Email    string `json:"email" yaml:"email"`
	Level    uint8  `json:"level" yaml:"level"`
	Token    string `json:"token" yaml:"token"`
}
