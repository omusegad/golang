package recievers

import "errors"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// add reciever function,
func (p *Person) GetAge() bool {
	if p.Age > 20 {
		return true
	}
	return false
}

// get person full name
func (p *Person) FullName() string {
	return p.FirstName + p.LastName
}

// update first name
func (p *Person) SetFirstName(f string) error {
	if len(f) < 3 {
		//panic("invalid first name")
		return errors.New("Invalid FirstName")
	}
	p.FirstName = f
	return nil
}


// Accept an object as a constructor 
func NewPerson(fn, ln string, ag int) Person{
	return Person{
		FirstName: fn,
		LastName: ln,
		Age: ag,
	}
}