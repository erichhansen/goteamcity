 package goteamcity

import(
    "fmt"
    "encoding/json"
    "os"
)

type teamCityConfig struct {
  ConfigFilePath string
}

type configurer interface {
  getConfig() configuration
}

type configuration struct {
    TeamCityUrl string
    TeamCityUsername string
    TeamCityPassword string
}

func (tcConfig teamCityConfig) getConfig() configuration {
    file, err := os.Open(tcConfig.ConfigFilePath)
    if err != nil {
        fmt.Println("error:", err)
    }
    defer file.Close();

    decoder := json.NewDecoder(file)
    config := configuration{}
    err = decoder.Decode(&config)
    if err != nil {
        fmt.Println("error:", err)
    }
    return config;
}
