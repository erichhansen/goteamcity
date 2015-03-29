## goteamcity

A [go](http://golang.org) package for determining the status of a [TeamCity](https://www.jetbrains.com/teamcity/) continuous intgration server. 

###Setup
You can install this package with `go get github.com/erichhansen/goteamcity`

You will need a 'conf.json' file in the directory where you are running this app. It will have the TeamCity url as well as username and password to use to connect to TeamCity. See [exampleconf.json](exampleconf.json)

### Response Values
- **Success** All projects report success
- **Failure** One or more projects report failure and are not being investigated. 
- **Investigating** One or more projects are failing, but all failing projects are being investigated. 
