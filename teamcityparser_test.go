 package goteamcity

import(
    "fmt"
    "os"
    "testing"
    "io"
)

type stubInvestigationReader struct {
    Investigating bool
}

func (s stubInvestigationReader) IsInvestigating(name string) bool {
    return s.Investigating
}

func (s stubInvestigationReader) ReadInvestigation(resp io.ReadCloser) bool {
    return s.Investigating
}


func TestParseResponseShouldBeFail(t *testing.T) {
    file, err := os.Open("testfiles/teamcity1.json")
    if err != nil {
        fmt.Println("error:", err)
    }

    stub := stubInvestigationReader{}
    status := parseResponse(file, stub)
    
    if status != Fail {
    	t.Fatalf("Should be Fail state")
    }
}