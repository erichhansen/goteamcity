 package goteamcity

import(
    "fmt"
    "encoding/json"
    "os"
)

type configuration struct {
    TeamCityUrl string
    TeamCityUsername string
    TeamCityPassword string
}

func getTeamCityConfig() configuration {
    file, err := os.Open("conf.json")
    if err != nil {
        fmt.Println("error:", err)
    }
    decoder := json.NewDecoder(file)
    config := configuration{}
    err = decoder.Decode(&config)
    if err != nil {
        fmt.Println("error:", err)
    }
    return config;
}