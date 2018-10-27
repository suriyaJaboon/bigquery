package structs

type ObjectType struct {
	Database struct {
		PostgresServer []PostgresServer `json: "postgresserver"`
		RedisServer []RedisServer `json: "redisserver"`
		ElasticServer []ElasticServer `json: "elasticserver"`
	} `json: "database"`
	SettingMode SettingMode `json:"settingmode"`
	Logs Logs `json: "logs"`
}

type SettingMode struct {
	PrivateKeyPath string `json:"private_key_path"`
	PublicKeyPath string `json:"public_key_path"`
	JWTExpirationDelta int `json:"jwt_expiration_delta"`
}

type Logs struct {
	Success string `json: "success"`
	Error string `json: "error"`
	Application string `json:"application"`
}