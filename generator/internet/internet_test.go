package internet

import (
	"net/mail"
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestUser(t *testing.T) {
	i := faker.New().Internet()

	user := i.User()
	tool.Expect(t, true, len(user) > 0)
	tool.Expect(t, false, strings.Contains(user, " "))
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func TestDomain(t *testing.T) {
	i := faker.New().Internet()

	domain := i.Domain()

	tool.Expect(t, true, len(domain) > 0)
	tool.Expect(t, true, strings.Index(domain, ".") > 0)

	split := strings.Split(domain, ".")
	tool.Expect(t, 2, len(split))
}

func TestEmail(t *testing.T) {
	i := faker.New().Internet()

	email := i.Email()
	split := strings.Split(email, "@")

	tool.Expect(t, 2, len(split))
	tool.Expect(t, true, isValidEmail(email))
}

func TestFreeEmail(t *testing.T) {
	i := faker.New().Internet()

	email := i.FreeEmail()
	split := strings.Split(email, "@")

	tool.Expect(t, 2, len(split))
	tool.Expect(t, true, isValidEmail(email))
}

func TestSafeEmail(t *testing.T) {
	i := faker.New().Internet()

	email := i.SafeEmail()
	split := strings.Split(email, "@")

	tool.Expect(t, 2, len(split))
	tool.Expect(t, true, isValidEmail(email))
}

func TestCompanyEmail(t *testing.T) {
	i := faker.New().Internet()

	email := i.CompanyEmail()
	split := strings.Split(email, "@")

	tool.Expect(t, 2, len(split))
	tool.Expect(t, false, strings.Contains(email, " "))
	tool.Expect(t, true, isValidEmail(email))
}

func TestPassword(t *testing.T) {
	i := faker.New().Internet()

	tool.Expect(t, true, len(i.Password()) >= 6)
}

func TestTLD(t *testing.T) {
	i := faker.New().Internet()

	tool.Expect(t, true, len(i.TLD()) > 0)
}

func TestSlug(t *testing.T) {
	i := faker.New().Internet()

	tool.Expect(t, true, len(i.Slug()) > 0)
}

func TestURL(t *testing.T) {
	i := faker.New().Internet()

	tool.Expect(t, true, len(i.URL()) > 0)
}

func TestIpv4(t *testing.T) {
	i := faker.New().Internet()

	ip := i.Ipv4()
	tool.Expect(t, true, len(ip) > 0)
	split := strings.Split(ip, ".")
	tool.Expect(t, 4, len(split))
}

func TestLocalIpv4(t *testing.T) {
	i := faker.New().Internet()

	ip := i.LocalIpv4()
	tool.Expect(t, true, len(ip) > 0)
	split := strings.Split(ip, ".")
	tool.Expect(t, 4, len(split))
}

func TestIpv6(t *testing.T) {
	i := faker.New().Internet()

	ip := i.Ipv6()
	tool.Expect(t, true, len(ip) > 0)
	tool.Expect(t, 39, len(ip))

	split := strings.Split(ip, ":")
	tool.Expect(t, 8, len(split))
}

func TestMacAddress(t *testing.T) {
	i := faker.New().Internet()

	tool.Expect(t, 17, len(i.MacAddress()))
}

func TestHTTPMethod(t *testing.T) {
	i := faker.New().Internet()

	tool.Expect(t, true, len(i.HTTPMethod()) > 0)
}

func TestQuery(t *testing.T) {
	i := faker.New().Internet()

	query := i.Query()
	tool.Expect(t, 0, strings.Index(query, "?"))
}

func TestStatusCode(t *testing.T) {
	i := faker.New().Internet()

	tool.NotExpect(t, 0, i.StatusCode())
}

func TestStatusCodeMessage(t *testing.T) {
	i := faker.New().Internet()

	tool.NotExpect(t, "", i.StatusCodeMessage())
}

func TestStatusCodeWithMessage(t *testing.T) {
	i := faker.New().Internet()

	tool.NotExpect(t, "", i.StatusCodeWithMessage())
}
