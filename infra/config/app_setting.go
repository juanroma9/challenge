package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"unicode/utf8"
)

// AppSettings represents the struct of the YAML configuration file.
type AppSettings struct {
	AppConfig YAMLConfigFile `yaml:"appConfig"`
}

// Constructor
func NewAppSettings(configFileName string) *AppSettings {

	data, errIO := os.ReadFile(configFileName)
	if errIO != nil {
		fmt.Sprintf("failed to read file %s: %s", configFileName, errIO.Error())
		return nil
	}

	fmt.Println(string(data))
	instance := AppSettings{}

	err := yaml.Unmarshal(data, &instance)
	if err != nil {
		return nil
	}

	return &instance
}

type YAMLConfigFile struct {
	FileUploadConfig FileUploadConfig `yaml:"fileUploadConfig"`
	DBMySQLConfig    DBMySQLConfig    `yaml:"dbMySQLConfig"`
	EndPointConfig   EndPoints        `yaml:"endpointConfig"`
}

type FileUploadConfig struct {
	ContentType []string `yaml:"contentType"`
	Separator   []string `yaml:"separator"`
	Encoding    string   `yaml:"encoding"`
}

type DBMySQLConfig struct {
	DriverName     string `yaml:"driverName"`
	DataSourceName string `yaml:"dataSourceName"`
}

type EndPoints struct {
	ApiItems    string `yaml:"apiItems"`
	ApiSeller   string `yaml:"apiSeller"`
	ApiCategory string `yaml:"apiCategory"`
	ApiCurrency string `yaml:"apiCurrency"`
}

func (as AppSettings) GetMySQLConfig() *DBMySQLConfig {
	return &as.AppConfig.DBMySQLConfig
}

// Helper function to check if a value is in the allowed list
func (as AppSettings) IsAllowedSeparator(value string) (string, error) {

	for _, item := range as.AppConfig.FileUploadConfig.Separator {
		if strings.Contains(value, item) {
			return item, nil
		}
	}

	return "", errors.New("The following value is not allowed: " + value)
}

func (as AppSettings) IsAllowedEncoding(arrayByte []byte) (bool, error) {

	if as.AppConfig.FileUploadConfig.Encoding == "utf-8" {
		if utf8.Valid(arrayByte) {
			return true, nil
		}
	}

	return false, errors.New("The file content is not in an allowed encoding")
}

func (as AppSettings) IsAllowedContentType(value string) error {

	for _, item := range as.AppConfig.FileUploadConfig.ContentType {
		if value == item {
			return nil
		}
	}

	return errors.New("The following value is not allowed: " + value)
}
