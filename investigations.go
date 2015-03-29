package goteamcity

import(
    "fmt"
    "encoding/json"
    "net/http"
    "io"
    "log"
)

type investigationsResponse struct {
    Investigation []investigation
}

type investigation struct {
    // there is a lot more to this but I only care about this
    State string
}

func isInvestigating(name string) bool {
	config := getTeamCityConfig()
    url := config.TeamCityUrl + fmt.Sprintf(investigationsPath, name);
	client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("Accept", "application/json")
    req.SetBasicAuth(config.TeamCityUsername, config.TeamCityPassword)

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error: %s", err)
    } 
	return parseInvestigation(resp.Body);
}

func parseInvestigation(resp io.ReadCloser) bool {
	decoder := json.NewDecoder(resp)
    response := investigationsResponse{}
    err := decoder.Decode(&response)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }

    // should be at most 1 
    investigationCount := len(response.Investigation)
    if investigationCount == 0 {
        return false
    } else if investigationCount == 1 {
        state := response.Investigation[0].State
        if state == "TAKEN" {
            return true
        }
        return false
    }

    log.Fatalf("Error: Wow, much investigation, such bad.")
    return false
}