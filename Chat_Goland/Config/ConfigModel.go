package Config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

// Config 定義應用程式配置結構
type Config struct {
	Redis RedisConfig `yaml:"redis"`
	MySql MySqlConfig `yaml:"mysql"`
	Jwt   JwtConfig   `yaml:"Jwt"`
}

// RedisConfig 定義 Redis 配置結構
type RedisConfig struct {
	Address      string        `yaml:"Address"`
	Password     string        `yaml:"password"`
	DB           int           `yaml:"db"`
	PoolSize     int           `yaml:"pool_size"`
	MinIdleConns int           `yaml:"min_idle_conns"`
	DialTimeout  time.Duration `yaml:"dial_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

// MySqlConfig 定義 Redis 配置結構
type MySqlConfig struct {
	UserName           string `yaml:"UserName"`
	Password           string `yaml:"Password"`
	Address            string `yaml:"Address"`
	Port               int    `yaml:"Port"`
	Database           string `yaml:"Database"`
	MaxLifetime        int    `yaml:"MaxLifetime"`
	MaxOpenConnections int    `yaml:"MaxOpenConnections"`
	MaxIdleConnections int    `yaml:"MaxIdleConnections"`
}

type JwtConfig struct {
	SecretKey string `yaml:"secret_key"`
}

// LoadConfig 讀取 YAML 配置
func LoadConfig() (*Config, error) {
	file, err := os.Open("Config/Config.yaml")
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
