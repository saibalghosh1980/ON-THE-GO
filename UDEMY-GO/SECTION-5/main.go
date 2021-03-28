package main

import (
	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	bo "section5/businessobject"
)

func main(){
    var crowd map[string]bo.Person
    crowd=make(map[string]bo.Person)//This line is required to instantiate

	alex := bo.Person{
		FirstName: "Alex",
		LastName:  "Berry",
	}
	id := guuid.New().String()
	log.Info(id)
	crowd[id]=alex

	for idval,item := range crowd{
		log.Info(idval)
		log.Info(item.ToString())
		log.Info(item.FullName)
		crowd[idval] = item //If you do not do this the values will not be updated in the original map
		//log.Info(crowd[idval].FullName)
	}
	//delete(crowd, id)

	log.Info(crowd)

}
