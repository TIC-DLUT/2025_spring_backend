package config

type Config struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	JWTPassword string `json:"JWTPassword"`
	BasePath    string `json:"BasePath"`
	ApiKey      string `json:"ApiKey"`
	Model       string `json:"Model"`
}
