package redis

// Config holds redis config
type Config struct {
	Address  string
	Password string
	DB       int
}
