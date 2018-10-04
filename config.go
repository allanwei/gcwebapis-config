package config

import (
	"os"
	"path/filepath"

	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

//GetDatabaseConf ...
func GetDatabaseConf() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)
	var conn string
	err = config.Load(file.NewSource(
		file.WithPath(exPath + "/conf.json")))
	if err != nil {
		return "", err
	}
	err = config.Get("connections", "groundcastdb").Scan(&conn)
	if err != nil {
		return "", err
	}
	return conn, nil

}
