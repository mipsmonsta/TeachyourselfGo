package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	usernames := []string{
		"slimshaddy99",
		"Thebestusernameever",
		"ds123!!#$$",
	}

	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}$")
	an := regexp.MustCompile("[[:^alnum:]]") //posix character class [:alnum:] - not alphanumeric

	for _, username := range usernames {
		username = strings.TrimSpace(username)
		if len(username) > 12 {
			username = username[:12]
			fmt.Printf("trimmed username to %v\n", username)
		}
		
		if !re.MatchString(username){
			username = an.ReplaceAllString(username, "z")
			fmt.Printf("replaced non-alpha characters with z to %s\n", username)
		}
	}

}