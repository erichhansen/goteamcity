 package goteamcity

import(
    "fmt"
    "os"
    "testing"
)

func TestParseInvestigationShouldReturnTrue(t *testing.T) {
	file, err := os.Open("testfiles/investigation1.json")
    if err != nil {
        fmt.Println("error:", err)
    }
    
    reader := teamCityInvestigationReader{}
    isInvg := reader.ReadInvestigation(file)

    if !isInvg {
    	t.Fatalf("Should be investigating")
    }
}


func TestParseInvestigationShouldReturnFalse(t *testing.T) {
	file, err := os.Open("testfiles/investigation2.json")
    if err != nil {
        fmt.Println("error:", err)
    }

    reader := teamCityInvestigationReader{}
    isInvg := reader.ReadInvestigation(file)

    if isInvg {
    	t.Fatalf("Should not be investigating")
    }
}

