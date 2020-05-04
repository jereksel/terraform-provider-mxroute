package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// GetAllDomains returns all domains for given account
func GetAllDomains(username string, password string) ([]string, error) {
	client := &http.Client{}
	//https://www.directadmin.com/features.php?id=336
	req, err := http.NewRequest("GET", "https://echo.mxrouting.net:2222/CMD_API_SHOW_DOMAINS", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	s := string(bodyText)
	values, err := url.ParseQuery(s)
	if err != nil {
		return nil, err
	}
	return values["list[]"], nil
}

//TODO: Test invalid domainName
func CreateDomain(username string, password string, domainName string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://echo.mxrouting.net:2222/CMD_API_DOMAIN", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	//https://www.directadmin.com/features.php?id=498
	q.Add("action", "create")
	q.Add("domain", domainName)
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	s := string(bodyText)
	values, err := url.ParseQuery(s)
	if err != nil {
		return err
	}
	if values["error"][0] == "1" {
		return fmt.Errorf("error returned from DirectAdmin API. Text: %s, Details: %s", values["text"][0], values["details"][0])
	}
	return nil
}

func RemoveDomain(username string, password string, domainName string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://echo.mxrouting.net:2222/CMD_API_DOMAIN", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	//https://www.directadmin.com/features.php?id=498
	q.Add("delete", "anything")
	q.Add("confirmed", "anything")
	q.Add("select0", domainName)
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	s := string(bodyText)
	values, err := url.ParseQuery(s)
	if err != nil {
		return err
	}
	if values["error"][0] == "1" {
		return fmt.Errorf("error returned from DirectAdmin API. Text: %s, Details: %s", values["text"][0], values["details"][0])
	}
	return nil
}

func GetDomainDkim(username string, password string, domainName string) (string, error) {
	return "", nil
}

func ExtractDcim(str string) (string, error) {
	re := regexp.MustCompile(`"(.*?)"`)

	domainKeyBlock := false
	domainKey := ""

	for _, line := range strings.Split(str, "\n") {

		if domainKeyBlock && !strings.HasPrefix(line, "\t") {
			domainKeyBlock = false
		}

		if strings.HasPrefix(line, "x._domainkey") {
			domainKeyBlock = true
		}

		if domainKeyBlock {
			domainKey = domainKey + re.FindAllStringSubmatch(line, -1)[0][1]
		}

	}

	return domainKey, nil

}
