package goteamcity

import(
    "encoding/json"
    "io"
    "log"
)

const Success string = "Success"
const Fail string = "Fail"
const Investigating string = "Investigating"
var investigationsPath string = "/httpAuth/app/rest/investigations?locator=buildType:(name:%s)"

type teamcityResponse struct {
	Project []investigater
}

type project struct {
    Name, WebUrl, LastBuildTime, LastBuildLabel, LastBuildStatus string
}

func parseResponse(response io.ReadCloser, r teamcityResponse) string {
    decoder := json.NewDecoder(response)
    err := decoder.Decode(&r)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    projectCount := len(r.Project)
    successCount := 0
    failureCount := 0
    investigateCount := 0

    for i := 0; i < projectCount; i++ {
    	proj := r.Project[i]
        buildStatus := proj.LastStatus()
        if buildStatus == "Success" {
        	successCount++
        } else if buildStatus == "Failure" {
        	if proj.IsInvestigating() {
        		investigateCount++
        	}
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



