package goteamcity

import(
    "fmt"
    "encoding/json"
    "os"
    "net/http"
)

type configuration struct {
    TeamCityUrl string
}

func GetTeamCityStatus() string {
    url := getTeamCityUrl()
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("error:", err)
    }

    fmt.Println(resp)
    return "Failed"
}
  
func getTeamCityUrl() string {
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
    fmt.Println(config.TeamCityUrl)
    return config.TeamCityUrl;
}