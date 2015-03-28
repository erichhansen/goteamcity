package goteamcity

import(
    "encoding/json"
    "io"
    "log"
)

const Success string = "Success"
const Fail string = "Fail"
const Investigating string = "Investigating"

type project struct {
	Name string
	WebUrl string
	LastBuildTime string
	LastBuildLabel string
	LastBuildStatus string
}

type teamcityResponse struct {
	Project []project
}

//const projectRequestPath string = "/httpAuth/app/rest/cctray/projects.xml"

func parseResponse(response io.ReadCloser) string {
    decoder := json.NewDecoder(response)
    teamCityStatus := teamcityResponse{}
    err := decoder.Decode(&teamCityStatus)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    projectCount := len(teamCityStatus.Project)
    successCount := 0
    failureCount := 0
    investigateCount := 0

    for i := 0; i < projectCount; i++ {
        buildStatus := teamCityStatus.Project[i].LastBuildStatus
        if buildStatus == "Success" {
        	successCount++
        } else if buildStatus == "Failure" {
        	failureCount++
        }
    }
    
    if successCount > 0 && failureCount == 0 {
    	return Success
    } else if failureCount == investigateCount {
    	return Investigating
    } 

    return Fail
}