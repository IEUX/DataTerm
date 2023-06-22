package credentials

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

var (
	domainNameRegEx string = "^((?!-))(xn--)?[a-z0-9][a-z0-9-_]{0,61}[a-z0-9]{0,1}\.(xn--)?([a-z0-9\-]{1,61}|[a-z0-9-]{1,30}\.[a-z]{2,})$"
	ipRegEx string  = "^((25[0-5]|(2[0-4]|1[0-9]|[1-9]|)[0-9])(\.(?!$)|$)){4}$"
)

func ReadFlags() {
	creds := flag.String("h", "-", "user@host")
	port := flag.Int("p", 3306, "database port number")
	flag.Parse()

	fmt.Println(*creds, *port)
}

func ValidCreds(creds string) (bool, error) {
	//Check length and format of credentials
	splitCreds := strings.Split(creds, "@")
	if len(splitCreds) != 2 {
		return false, fmt.Errorf("Creds format is not user@host")
	}
	
	//Match hostname with a valid domain name or IP
	host := splitCreds[1]
	domainNameMatch, err := regexp.MatchString(domainNameRegEx, host)
	if err != nil {
		return false, err
	}
	ipMatch, err := regexp.MatchString(ipRegEx, host)
	if err != nil {
		return false, err
	}
	if !domainNameMatch && !ipMatch {
		return false,fmt.Errorf("host format not valid")
	}
	return true, nil
}
