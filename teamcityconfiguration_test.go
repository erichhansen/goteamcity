package goteamcity

import(
    "testing"
)

func TestGetConfigShouldReadContents(t *testing.T) {
  teamCityConfig := teamCityConfig{ConfigFilePath: "testfiles/testconf.json"}
  config := teamCityConfig.getConfig();

  if (config.TeamCityUsername != "guest") {
    t.Fatalf("Username should be 'guest'")
  }

  if (config.TeamCityPassword != "guest") {
    t.Fatalf("Password should be 'guest'")
  }

  if (config.TeamCityUrl != "http://teamcity:8111") {
    t.Fatalf("Url should be 'http://teamcity:8111'")
  }
}
