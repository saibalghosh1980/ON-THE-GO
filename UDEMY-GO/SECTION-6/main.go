package main

import "github.com/sirupsen/logrus"

type human interface {
	sayHello() string
}

type male struct {
	name string
	sex string
}

func (m male) sayHello() string{
	return m.sex
}

func (m female) sayHello() string{
	return m.sex
}


type female struct {
	name string
	sex string
}

func printNameAndSex(h human){
	logrus.Info(h.sayHello())
}

func main(){


	saibal :=male{
		name: "saibal",
		sex:  "male",
	}

	arundhati := female{
		name: "arundhati",
		sex:  "female",
	}

	printNameAndSex(saibal)
	printNameAndSex(arundhati)


}
