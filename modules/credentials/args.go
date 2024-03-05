package credentials

import (
	ansi "DataTerm/modules/ANSI"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Credentials struct {
	User     string
	Password string
	Host     string
	Database string
	Port     int
}

var (
	domainNameRegEx, _ = regexp.Compile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9])).([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}.[a-zA-Z]{2,3})$`)
	ipRegEx, _         = regexp.Compile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
)

func ReadFlags() (Credentials, string) {
	creds := flag.String("host", "", fmt.Sprintf(ansi.WarningColor, "[REQUIRED] -host=user@host"))
	port := flag.Int("p", 3306, "database port number")
	driver := flag.String("driver", "mysql", "Driver for the database <link table>")
	database := flag.String("db", "", fmt.Sprintf(ansi.WarningColor, "[REQUIRED] -db=<database name>"))
	flag.Parse()
	//Check Creds
	isOk, err := ValidCreds(*creds)
	if err != nil {
		log.Fatalf(ansi.ErrorColor, "[INIT] "+err.Error())
	}
	if !isOk {
		log.Fatalf(ansi.ErrorColor, "[INIT] Fail to load credentials")
	}
	//Check database
	if *database == "" {
		log.Fatalf(ansi.ErrorColor, "[INIT] database name is required (see required flags -h)")
	}
	user, host := strings.Split(*creds, "@")[0], strings.Split(*creds, "@")[1]
	//Check Driver
	if !ValidDriver(*driver) {
		log.Fatalf(ansi.ErrorColor, "[INIT] \""+*driver+"\" driver is not supported")
	}
	return Credentials{User: user, Host: host, Password: "_", Database: *database, Port: *port}, *driver
}

func ValidCreds(creds string) (bool, error) {
	//Check length and format of credentials
	splitCreds := strings.Split(creds, "@")
	if creds == "" {
		return false, fmt.Errorf("host is required (see required flags -h)")
	}
	if len(splitCreds) != 2 {
		return false, fmt.Errorf("creds format must be user@host")
	}
	//Match hostname with a valid domain name or IP
	host := splitCreds[1]
	if !domainNameRegEx.MatchString(host) && !ipRegEx.MatchString(host) {
		return false, fmt.Errorf("can't match hostname with a domain name or an IP adress")
	}
	return true, nil
}

func ValidDriver(driver string) bool {
	availableDrivers := map[string]bool{
		"mysql":    true,
		"postgres": true,
		"odbc":     false,
	}
	return availableDrivers[driver]
}
