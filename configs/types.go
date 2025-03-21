package configs

type Configs struct {
	Port     int      `json:"port" yaml:"port"`
	Database database `json:"database" yaml:"database"`
}

type database struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     int64  `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
}
