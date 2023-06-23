package configuration

import (
	"os"

	"github.com/kataras/iris/v12"

	"gopkg.in/yaml.v3"
)

// Configuration holds the necessary information
// for our server, including the Iris one.
type Configuration struct {
	ServerName string `yaml:"ServerName"`
	Env        string `yaml:"Env"`
	// The server's host, if empty, defaults to 0.0.0.0
	Host string `yaml:"Host"`
	// The server's port, e.g. 80
	Port int `yaml:"Port"`
	// If not empty runs under tls with this domain using letsencrypt.
	Domain string `yaml:"Domain"`
	// Enables write response and read request compression.
	EnableCompression bool `yaml:"EnableCompression"`
	// Defines the "Access-Control-Allow-Origin" header of the CORS middleware.
	// Many can be separated by comma.
	// Defaults to "*" (allow all).
	AllowOrigin string `yaml:"AllowOrigin"`
	// If not empty a request logger is registered,
	// note that this will cost a lot in performance, use it only for debug.
	RequestLog string `yaml:"RequestLog"`
	// The database source configuration.
	DataSource *DBConfiguration `yaml:"DataSource"`
	// Iris specific configuration.
	Iris iris.Configuration `yaml:"Iris"`
}

type MysqlConfiguration struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	DataBase string `yaml:"DataBase"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"Password"`
}

type RedisConfiguration struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"Password"`
}

type SqliteConfiguration struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"Password"`
}

type DBConfiguration struct {
	Mysql  *MysqlConfiguration  `yaml:"Mysql"`
	Redis  *RedisConfiguration  `yaml:"Redis"`
	Sqlite *SqliteConfiguration `yaml:"Sqlite"`
}

// BindFile binds the yaml file's contents to this Configuration.
func (c *Configuration) BindFile(filename string) error {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(contents, c)
}
