package businessobject

type Person struct{

	FirstName string
	LastName string
	Age int16
	FullName string

}

func (person *Person) ToString() string  {

	(*person).FullName=person.FirstName+" "+person.LastName
	return person.FullName

}

func (person *Person) ToString2() string  {

	person.FullName="Mr. "+person.FirstName+" "+person.LastName
	return person.FullName

}


