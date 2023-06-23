package useragent

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestInternetExplorer(t *testing.T) {
	u := faker.New().UserAgent()
	tool.NotExpect(t, "", u.InternetExplorer())
}

func TestOpera(t *testing.T) {
	u := faker.New().UserAgent()
	tool.NotExpect(t, "", u.Opera())
}

func TestSafari(t *testing.T) {
	u := faker.New().UserAgent()
	tool.NotExpect(t, "", u.Safari())
}

func TestFirefox(t *testing.T) {
	u := faker.New().UserAgent()
	tool.NotExpect(t, "", u.Firefox())
}

func TestChrome(t *testing.T) {
	u := faker.New().UserAgent()
	tool.NotExpect(t, "", u.Chrome())
}

func TestUserAgent(t *testing.T) {
	u := faker.New().UserAgent()
	tool.Expect(t, true, len(u.UserAgent()) > 0)
}
