package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	QueryBaseFolder string `default:"query"`
	DataFolder      string `default:"../../data"`

	HostDB     string `default:"localhost" split_words:"true"`
	PortDB     string `default:"3306" split_words:"true"`
	UserDB     string `default:"user" split_words:"true"`
	PasswordDB string `default:"password" split_words:"true"`

	TransactionTableName       string `default:"transacoes" split_words:"true"`
	TransactionTableSchemaName string `default:"points" split_words:"true"`

	TransactionProductTableName       string `default:"transacao_produto" split_words:"true"`
	TransactionProductTableSchemaName string `default:"points" split_words:"true"`

	CustomerTableName       string `default:"clientes" split_words:"true"`
	CustomerTableSchemaName string `default:"points" split_words:"true"`

	ProductTableName       string `default:"produtos" split_words:"true"`
	ProductTableSchemaName string `default:"points" split_words:"true"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	err := envconfig.Process("", config)
	return config, err
}
