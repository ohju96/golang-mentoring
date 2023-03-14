package toml

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Local struct {
		LocalServer string
		Dbms        string
		Db          string
		User        string
		Password    string
		Host        string
	}
}

func InitToml(path string) Config {

	config := new(Config)

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed toml setting: %v", err)
	}

	if _, err := toml.NewDecoder(file).Decode(config); err != nil {
		log.Fatalf("failed toml setting: %v", err)
	}

	return *config
}
