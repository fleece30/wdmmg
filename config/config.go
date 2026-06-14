package config

type Conf struct {
	Server ServerConf
	DB     DBConf
}

type ServerConf struct {
	ServerAddress string `env:"SERVER_ADDRESS,required"`
}

type DBConf struct {
	DBName   string `env:"DB_NAME,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
}
