package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// browser,company,country,email,job,name,phone
type Users struct {
	Browsers []string
	Name     string
	Email    string
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	/*
		!!! !!! !!!
		обратите внимание - в задании обязательно нужен отчет
		делать его лучше в самом начале, когда вы видите уже узкие места, но еще не оптимизировалм их
		так же обратите внимание на команду в параметром -http
		перечитайте еще раз задание
		!!! !!! !!!
	*/
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileContents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile("@")
	seenBrowsers := make(map[string]bool)
	foundUsers := ""

	lines := strings.Split(string(fileContents), "\n")
	users := make([]Users, 0, len(lines))
	for _, line := range lines {
		user := Users{}
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	for i, user_brows := range users {
		isAndroid := false
		isMSIE := false

		for _, browser := range user_brows.Browsers {
			if match, err := regexp.MatchString("Android", browser); match && err == nil {
				isAndroid = true
				if seenBrowsers[browser] != true {
					seenBrowsers[browser] = true
				}
			}
			if match, err2 := regexp.MatchString("MSIE", browser); match && err2 == nil {
				isMSIE = true
				if seenBrowsers[browser] != true {
					seenBrowsers[browser] = true
				}
			}

		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := r.ReplaceAllString(user_brows.Email, " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user_brows.Name, email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
