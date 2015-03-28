package goteamcity

import(
    "fmt"
    "encoding/json"
    "os"
    "net/http"
    "io/ioutil"
    "log"
)

const requestPath string = "/httpAuth/app/rest/cctray/projects.xml"

type configuration struct {
    TeamCityUrl string
    TeamCityUsername string
    TeamCityPassword string
}

func GetTeamCityStatus() string {
    config := getTeamCityConfig()
    url := config.TeamCityUrl + requestPath;

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Accept", "application/json")
    req.SetBasicAuth(config.TeamCityUsername, config.TeamCityPassword)

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error: %s", err)
    } else {
        defer resp.Body.Close()
        contents, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Error: %s", err)
        }
        fmt.Println(string(contents))
    }

    return "ToDo"
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