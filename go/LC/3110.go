package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	emailsLen := len(emails)
	uniqueEmails := make(map[string]bool)
	for i := 0; i < emailsLen; i++ {
		e := emails[i]
		eLen := len(e)
		j := 0
		var currLocal strings.Builder
		for j < eLen {
			fmt.Println(currLocal.String())
			currChar := string(e[j])
			if currChar == "@" {
				break
			}
			if currChar == "+" {
				break
			}
			if currChar == "." {
				j++
				currChar = string(e[j])
			}
			currLocal.WriteString(currChar)
			j++
		}
		if string(e[j]) == "+" {
			for j < eLen && string(e[j]) != "@" {
				j++
			}
		}
		currDomain := string(e[j:eLen])
		currMail := currLocal.String() + currDomain
		_, ok := uniqueEmails[currMail]
		if !ok {
			uniqueEmails[currMail] = true
		}
	}
	fmt.Println(uniqueEmails)
	fmt.Println(len(uniqueEmails))
}
