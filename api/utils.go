package api

func DoesDomainExist(username, password, domainName string) (*bool, error) {
	allDomains, err := GetAllDomains(username, password)
	if err != nil {
		return nil, err
	}
	exists := false
	for _, domain := range allDomains {
		if domain == domainName {
			exists = true
		}
	}
	return &exists, nil
}

func DoesEmailAccountExists(username, password, domainName, emailUsername string) (*bool, error) {
	doesDomainExist, err := DoesDomainExist(username, password, domainName)
	if err != nil {
		return nil, err
	}
	if !*doesDomainExist {
		temp := false
		return &temp, nil
	}
	allEmailAccounts, err := GetEmailAccounts(username, password, domainName)
	exists := false
	for _, email := range allEmailAccounts {
		if email == emailUsername {
			exists = true
		}
	}
	return &exists, nil
}
