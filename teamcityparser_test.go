 package goteamcity

import(
    "fmt"
    "os"
    "testing"
)

func TestParseResponseShouldBeFail(t *testing.T) {

    file, err := os.Open("testfiles/teamcity1.json")
    if err != nil {
        fmt.Println("error:", err)
    }

    status := parseResponse(file)

    if status != Fail {
    	t.Fatalf("Should be Fail state")
    }
}