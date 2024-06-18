package belajar_golang_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_Viper(t *testing.T) {
	var config = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	name := config.GetString("app.name")
	assert.Equal(t, "app-belajar", name)
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
}

func TestYaml(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "app-belajar", config.GetString("app.name"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
}

func TestENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "app-belajar", config.GetString("APP_NAME"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
}

func TestAutomaticENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "app-belajar", config.GetString("APP_NAME"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
}
