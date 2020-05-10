package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//https://forum.directadmin.com/threads/cmd_api_pop.9527/
func CreateEmailAccount(username, password, domainName, emailUsername, emailPassword string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://echo.mxrouting.net:2222/CMD_API_POP", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "create")
	q.Add("domain", domainName)
	q.Add("user", emailUsername)
	q.Add("quota", "0")
	q.Add("passwd", emailPassword)
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

//https://www.directadmin.com/features.php?id=1505
func GetEmailAccounts(username, password, domainName string) ([]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://echo.mxrouting.net:2222/CMD_API_POP", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("action", "list")
	q.Add("domain", domainName)
	req.URL.RawQuery = q.Encode()
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

//https://github.com/arian/DirectAdminApi/blob/1a18151/Source/DA/Emails.php#L123
func RemoveEmailAccount(username, password, domainName, emailUsername string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://echo.mxrouting.net:2222/CMD_API_POP", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "delete")
	q.Add("domain", domainName)
	q.Add("user", emailUsername)
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
