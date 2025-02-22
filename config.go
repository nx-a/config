package config

import "os"

type Config struct {
	data map[string]string
}

func New() *Config {
	cfg := new(Config)
	cfg.data = make(map[string]string, 1024)
	cfg.data["server.addr"] = def("SRV_ADDR", ":9090")
	cfg.data["db.showSql"] = def("DB_SHOW", "false")
	cfg.data["db.host"] = def("DB_HOST", "localhost")
	cfg.data["db.user"] = def("DB_USER", "postgres")
	cfg.data["db.password"] = def("DB_PASSWORD", "postgres")
	cfg.data["db.name"] = def("DB_NAME", "localhost")
	cfg.data["db.port"] = def("DB_PORT", "5432")
	cfg.data["db.schema"] = def("DB_SCHEMA", "public")
	cfg.data["db.ssl"] = def("DB_SSL", "disable")
	cfg.data["db.MaxIdleConns"] = def("DB_MAX_IDLE_COMM", "10")
	cfg.data["db.MaxOpenConns"] = def("DB_MAX_OPEN_COMM", "100")
	cfg.data["db.ConnMaxLifetime"] = def("DB_CONN_MAX_LIFE_TIME", "3600")
	return cfg
}
func (c *Config) Get(name string) string {
	return c.data[name]
}
func (c *Config) GetDef(name, def string) string {
	d, ok := c.data[name]
	if ok {
		return d
	}
	return def
}
func def(key string, _default string) string {
	res := os.Getenv(key)
	if res == "" {
		return _default
	}
	return res
}
