package config

type Config struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	JWTPassword string `json:"JWTPassword"`
}
