package useragent

// UserAgent is a faker struct for UserAgent
type UserAgent struct {
}

// InternetExplorer returns a fake internet explorer UserAgent
func (u UserAgent) InternetExplorer() string {
	return u.Faker.RandomStringElement(internetExplorerUserAgents)
}

// Chrome returns a fake chrome browser user agent for Internet
func (u UserAgent) Chrome() string {
	return u.Faker.RandomStringElement(chromeUserAgents)
}

// Firefox returns a fake firefox browser user agent for Internet
func (u UserAgent) Firefox() string {
	return u.Faker.RandomStringElement(firefoxUserAgents)
}

// Opera returns a fake opera browser user agent for Internet
func (u UserAgent) Opera() string {
	return u.Faker.RandomStringElement(operaUserAgents)
}

// Safari returns a fake safari browser user agent for Internet
func (u UserAgent) Safari() string {
	return u.Faker.RandomStringElement(safariUserAgents)
}

// UserAgent returns a fake browser user agent for Internet
func (u UserAgent) UserAgent() string {
	userAgents := []string{}
	userAgents = append(userAgents, chromeUserAgents...)
	userAgents = append(userAgents, firefoxUserAgents...)
	userAgents = append(userAgents, operaUserAgents...)
	userAgents = append(userAgents, safariUserAgents...)
	return u.Faker.RandomStringElement(userAgents)
}
