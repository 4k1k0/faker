package time

import (
	"time"
)

// Time is a faker struct for Time
type Time struct {
}

// Unix returns a fake unix for Time
func (t Time) Unix(max time.Time) int64 {
	return max.Unix() - t.Faker.Int64Between(0, max.Unix())
}

// Time returns a fake time for Time
func (t Time) Time(max time.Time) time.Time {
	return time.Unix(int64(t.Unix(max)), 0)
}

// TimeBetween returns a fake time between for Time
func (t Time) TimeBetween(min, max time.Time) time.Time {
	diff := time.Nanosecond * time.Duration(t.Faker.IntBetween(0, int(max.Sub(min).Nanoseconds())))
	duration := time.Nanosecond * diff
	return min.Add(duration)
}

// ISO8601 returns a fake time in ISO8601 format for Time
func (t Time) ISO8601(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format("2006-02-01T15:04:05+000")
}

// ANSIC returns a fake time in ANSIC format for Time
func (t Time) ANSIC(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.ANSIC)
}

// UnixDate returns a fake date in unix format for Time
func (t Time) UnixDate(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.UnixDate)
}

// RubyDate returns a fake date in ruby format for Time
func (t Time) RubyDate(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.RubyDate)
}

// RFC822 returns a fake time in RFC822 format for Time
func (t Time) RFC822(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.RFC822)
}

// RFC822Z returns a fake time in RFC822Z format for Time
func (t Time) RFC822Z(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.RFC822Z)
}

// RFC850 returns a fake time in RFC850 format for Time
func (t Time) RFC850(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.RFC850)
}

// RFC1123 returns a fake time in RFC1123 format for Time
func (t Time) RFC1123(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.RFC1123)
}

// RFC1123Z returns a fake time in RFC1123Z format for Time
func (t Time) RFC1123Z(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.RFC1123Z)
}

// RFC3339 returns a fake time in RFC3339 format for Time
func (t Time) RFC3339(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.RFC3339)
}

// RFC3339Nano returns a fake time in RFC3339Nano format for Time
func (t Time) RFC3339Nano(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.RFC3339Nano)
}

// Kitchen returns a fake time in Kitchen format for Time
func (t Time) Kitchen(max time.Time) string {
	t1 := t.Time(max)
	return t1.Format(time.Kitchen)
}

// AmPm returns a fake "am" or "pm" string for Time
func (t Time) AmPm() string {
	return t.Faker.RandomStringElement([]string{"am", "pm"})
}

// DayOfMonth returns a fake day of month for Time
func (t Time) DayOfMonth() int {
	return t.Faker.IntBetween(1, 31)
}

// DayOfWeek returns a fake day of week for Time
func (t Time) DayOfWeek() time.Weekday {
	days := []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}

	return days[t.Faker.IntBetween(0, len(days)-1)]
}

// Month returns a fake month for Time
func (t Time) Month() time.Month {
	months := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}

	return months[t.Faker.IntBetween(0, len(months)-1)]
}

// MonthName returns a fake month name for Time
func (t Time) MonthName() string {
	return t.Month().String()
}

// Year returns a fake year for Time
func (t Time) Year() int {
	return t.Faker.IntBetween(1970, time.Now().Year())
}

// Century returns a fake century for Time
func (t Time) Century() string {
	centuries := []string{
		"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
		"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI",
	}

	return t.Faker.RandomStringElement(centuries)
}

// Timezone returns a fake timezone for Time
func (t Time) Timezone() string {
	return t.Faker.RandomStringElement(timezones)
}
