package config


var (
	ListenAddr string
)

var defaultListenAddr = "0.0.0.0:8080"

func init() {
	if ListenAddr == "" {
		ListenAddr = defaultListenAddr
	}
}