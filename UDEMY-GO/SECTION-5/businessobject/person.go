package businessobject

type Person struct {
	FirstName string
	LastName string
	FullName string
}

func (person *Person) ToString() string  {

	person.FullName="Mr. "+person.FirstName+" "+person.LastName
	return person.FullName
}
