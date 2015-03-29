 package goteamcity

import(
    "fmt"
    "os"
    "testing"
)

type stubProject struct {
    Name, LastBuildStatus string
    Investigating bool
}

func (stub stubProject) IsInvestigating() bool {
    return stub.Investigating
}

func (stub stubProject) LastStatus() string {
    return stub.LastBuildStatus
}

func TestParseResponseShouldBeFail(t *testing.T) {

    file, err := os.Open("testfiles/teamcity1.json")
    if err != nil {
        fmt.Println("error:", err)
    }

    stub := stubProject{Investigating: false}
    proj := []project{stub}
    tc := teamcityResponse{}
    status := parseResponse(file, tc)

    if status != Fail {
    	t.Fatalf("Should be Fail state")
    }
}