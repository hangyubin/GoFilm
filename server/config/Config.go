package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config 应用配置结构体
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Spider   SpiderConfig
	Upload   UploadConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Charset  string
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// SpiderConfig 爬虫配置
type SpiderConfig struct {
	MaxGoroutine int
}

// UploadConfig 上传配置
type UploadConfig struct {
	FilmPictureUploadDir string
	FilmPictureUrlPath   string
	FilmPictureAccess    string
}

// LoadConfig 加载配置
func LoadConfig() *Config {
	config := &Config{}

	// 服务器配置
	config.Server.Port = getEnv("LISTENER_PORT", "3601")

	// 数据库配置
	config.Database.Type = getEnv("DATABASE_TYPE", "mysql")
	config.Database.Host = getEnv("DATABASE_HOST", "127.0.0.1")
	config.Database.Port = getEnv("DATABASE_PORT", "3306")
	config.Database.User = getEnv("DATABASE_USER", "root")
	config.Database.Password = getEnv("DATABASE_PASSWORD", "root")
	config.Database.DBName = getEnv("DATABASE_NAME", "FilmSite")
	config.Database.Charset = getEnv("DATABASE_CHARSET", "utf8mb4")

	// Redis配置
	config.Redis.Addr = fmt.Sprintf("%s:%s", 
		getEnv("REDIS_HOST", "127.0.0.1"), 
		getEnv("REDIS_PORT", "6379"))
	config.Redis.Password = getEnv("REDIS_PASSWORD", "root")
	config.Redis.DB, _ = strconv.Atoi(getEnv("REDIS_DB", "0"))

	// 爬虫配置
	config.Spider.MaxGoroutine, _ = strconv.Atoi(getEnv("MAX_GOROUTINE", "10"))

	// 上传配置
	config.Upload.FilmPictureUploadDir = getEnv("FILM_PICTURE_UPLOAD_DIR", "./static/upload/gallery")
	config.Upload.FilmPictureUrlPath = getEnv("FILM_PICTURE_URL_PATH", "/upload/pic/poster/")
	config.Upload.FilmPictureAccess = getEnv("FILM_PICTURE_ACCESS", "/api/upload/pic/poster/")

	return config
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	switch c.Type {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			c.User, c.Password, c.Host, c.Port, c.DBName, c.Charset)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Password, c.DBName)
	default:
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			c.User, c.Password, c.Host, c.Port, c.DBName, c.Charset)
	}
}
