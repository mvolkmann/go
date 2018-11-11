package encapsulation

import (
	"time"
)

// The "reference time" for time format strings
// is "Mon Jan 2 15:04:05 MST 2006".
// Format strings are built using these specific values
// and variations on them.
// For more detail, see the comments at https://golang.org/src/time/format.go.
// This is a date format string that is passed to the Time Format method.
const dateFormat = "January 2, 2006"

// Person describes a person.
type Person struct {
	Name     string    // exported
	birthday time.Time // encapsulated
}

// NewPerson creates and returns a new Person object.
// This is needed because the birthday field is not exported.
// If it were exported, code in other packages could
// directly create Person objects.
func NewPerson(name string, birthday time.Time) Person {
	return Person{name, birthday}
}

// Age calculates and returns the age of a Person.
func (p *Person) Age() int {
	todayY, todayM, todayD := time.Now().Date()
	birthdayY, birthdayM, birthdayD := p.birthday.Date()
	age := todayY - birthdayY
	// Subtract one if the birthday has not occurred yet this year.
	if todayM < birthdayM || (todayM == birthdayM && todayD < birthdayD) {
		age--
	}
	return age
}

// Birthday returns the value of the encapsulated field birthday.
func (p *Person) Birthday() time.Time {
	return p.birthday
}

// FormattedBirthday returns the formatted birthday of a Person.
func (p *Person) FormattedBirthday() string {
	return p.birthday.Format(dateFormat)
}
