package structs

type PostgresServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	DB string `json:"db"`
}

type RedisServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	DB int `json:"db"`
}

type ElasticServer struct {
	Protocol string `json:"protocol"`
	Host string `json:"host"`
	Port string `json:"port"`
}
