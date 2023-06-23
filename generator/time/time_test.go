package time

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestTimeUnix(t *testing.T) {
	tf := faker.New().Time()

	t1 := time.Now().Add(72 * time.Hour)
	t2 := tf.Unix(t1)

	tool.Expect(t, true, t1.Unix() > t2)
}

func TestTimeTime(t *testing.T) {
	tf := faker.New().Time()

	t1 := time.Now().Add(72 * time.Hour)
	t2 := tf.Time(t1)

	tool.Expect(t, true, t1.After(t2))
}

func TestISO8601(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, 23, len(tf.ISO8601(t1)))
}

func TestANSIC(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, 24, len(tf.ANSIC(t1)))
}

func TestUnixDate(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, 28, len(tf.UnixDate(t1)))
}

func TestRubyDate(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, 30, len(tf.RubyDate(t1)))
}

func TestRFC822(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, 19, len(tf.RFC822(t1)))
}

func TestRFC822Z(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, 21, len(tf.RFC822Z(t1)))
}

func TestRFC850(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	t2 := tf.RFC850(t1)
	tool.Expect(t, true, len(t2) > 0)
}

func TestRFC1123(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, 29, len(tf.RFC1123(t1)))
}

func TestRFC1123Z(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, 31, len(tf.RFC1123Z(t1)))
}

func TestRFC3339(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, true, len(tf.RFC3339(t1)) > 0)
}

func TestRFC3339Nano(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	tool.Expect(t, true, len(tf.RFC3339Nano(t1)) > 0)
}

func TestKitchen(t *testing.T) {
	tf := faker.New().Time()
	t1, _ := time.Parse("2006-01-02T15:04:05+000", "2017-01-02T15:04:05+000")
	t2 := tf.Kitchen(t1)
	tool.Expect(t, true, len(t2) >= 6)
	tool.Expect(t, true, len(t2) <= 7)
}

func TestTimeBetween(t *testing.T) {
	tf := faker.New().Time()

	t1 := time.Now()

	t2 := t1.Add(time.Duration(24) * time.Hour)

	t3 := tf.TimeBetween(t1, t2)

	tool.Expect(t, true, t1.Before(t3))
	tool.Expect(t, true, t2.After(t3))
}

func TestAmPm(t *testing.T) {
	tf := faker.New().Time()

	value := tf.AmPm()
	tool.Expect(t, true, (value == "am" || value == "pm"))
}

func TestDayOfMonth(t *testing.T) {
	tf := faker.New().Time()

	day := tf.DayOfMonth()
	tool.Expect(t, true, day > 0)
	tool.Expect(t, true, day <= 31)
}

func TestDayOfWeek(t *testing.T) {
	tf := faker.New().Time()

	day := tf.DayOfWeek()

	tool.Expect(t, true, day >= 0)
	tool.Expect(t, true, day <= 6)
	tool.Expect(t, "time.Weekday", fmt.Sprintf("%T", day))
}

func TestMonth(t *testing.T) {
	tf := faker.New().Time()

	month := tf.Month()
	tool.Expect(t, true, month >= 1)
	tool.Expect(t, true, month <= 12)
}

func TestMonthName(t *testing.T) {
	tf := faker.New().Time()

	name := tf.MonthName()
	tool.Expect(t, true, len(name) > 0)
}

func TestYear(t *testing.T) {
	tf := faker.New().Time()

	year := tf.Year()
	tool.Expect(t, true, year >= 1970)
	tool.Expect(t, 4, len(strconv.Itoa(year)))
}

func TestCentury(t *testing.T) {
	tf := faker.New().Time()

	century := tf.Century()
	tool.Expect(t, true, len(century) > 0)
}

func TestTimezone(t *testing.T) {
	tf := faker.New().Time()

	timezone := tf.Timezone()
	tool.Expect(t, true, len(timezone) > 0)
}
