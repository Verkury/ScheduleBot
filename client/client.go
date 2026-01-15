package client

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func StartClient() {
	baseUrl := "https://www.asu.ru/timetable/students/15/2129441959/"
	token, err := getToken(baseUrl)
	if err != nil {
		fmt.Println("Err:", err)
	}
	getSchedule(baseUrl, token)

}

func getToken(url string) (string, error) {
    doc, err := goquery.NewDocument(url)
    if err != nil {
        return "", err
    }
    
    var token string
    re := regexp.MustCompile(`'X-CS-ID',\s*'([^']+)'`)

    doc.Find("script").Each(func(i int, s *goquery.Selection) {
        if token != "" { return }
        
        matches := re.FindStringSubmatch(s.Text())
        if len(matches) > 1 {
            token = matches[1]
        }
    })

    if token == "" {
        return "", fmt.Errorf("token not found")
    }
    return token, nil
}

func getSchedule(url, token string) {

}