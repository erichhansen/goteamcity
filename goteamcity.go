package goteamcity

import(
    "fmt"
    "net/http"
    "log"
)

const projectRequestPath string = "/httpAuth/app/rest/cctray/projects.xml"

func GetTeamCityStatus() string {
    teamCityConfig := teamCityConfig{ConfigFilePath: "conf.json"}
    config := teamCityConfig.getConfig();
    url := config.TeamCityUrl + projectRequestPath;

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Accept", "application/json")
    req.SetBasicAuth(config.TeamCityUsername, config.TeamCityPassword)

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    defer resp.Body.Close()
    
    reader := teamCityInvestigationReader{}
    status := parseResponse(resp.Body, reader)
    fmt.Println(status)
    return status
}
