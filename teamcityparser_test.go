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

func TestParseResponseShouldBeInvestigating(t *testing.T) {
    file, err := os.Open("testfiles/teamcity1.json")
    if err != nil {
        fmt.Println("error:", err)
    }

    stub := stubInvestigationReader{Investigating: true}
    status := parseResponse(file, stub)

    if status != Investigating {
        t.Fatalf("Should be Investigating state")
    }
}

func TestParseResponseShouldBeSucess(t *testing.T) {
    file, err := os.Open("testfiles/teamcity2.json")
    if err != nil {
        fmt.Println("error:", err)
    }

    stub := stubInvestigationReader{Investigating: true}
    status := parseResponse(file, stub)

    if status != Success {
        t.Fatalf("Should be Success state")
    }
}

func TestParseNameShouldParse(t *testing.T) {
    projectName := "Example :: ProjectName"
    actualName := parseName(projectName)

    if actualName != "ProjectName" {
        t.Fatalf("Name did not parse correctly. Expected ProjectName got %s", actualName)
    }
}
