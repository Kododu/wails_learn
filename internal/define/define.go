package define

var ConfigName = "redis-client.json"

type Connection struct {
	Identify string `json:"identify"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Connections []*Connection `json:"connections"`
}
