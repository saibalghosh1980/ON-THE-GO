package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	bo "section4/businessobject"
)

func main()  {

	alex := bo.Person{FirstName: "Saibal", LastName: "Ghosh", Age: 42}
	fmt.Println(& alex)
	log.Info(alex.FullName)
	alexPointer := &alex
	log.Info(alexPointer.ToString())
	log.Info(alex.FullName)
	log.Info(alex.ToString2())
	log.Info(alex.FullName)


}
