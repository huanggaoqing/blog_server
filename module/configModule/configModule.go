package configModule

type Module struct {
	Mode     string   `yaml:"mode"`
	DataBase DataBase `yaml:"data_base"`
	Server   Server   `yaml:"server"`
	DataPool DataPool `yaml:"data_pool"`
}

// DataBase 数据库配置
type DataBase struct {
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
}

// Server 服务配置
type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// DataPool 数据库连接池配置
type DataPool struct {
	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxOpenConns int `yaml:"max_open_conns"`
	MaxLifetime  int `yaml:"max_lifetime"`
}
