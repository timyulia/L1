package main

import (
	"fmt"
	"strings"
)

type Human struct {
	firstName string
	lastName  string
	age       int
}

type Action struct {
	kind     string //вид действия
	duration int    //продолжительность в минутах
	Human           //исполнитель
}

func (h *Human) fullName() string { //конкатенация имени и фамилии
	var b strings.Builder
	b.WriteString(h.firstName)
	b.WriteByte(' ')
	b.WriteString(h.lastName)
	return b.String()
}

func main() {
	h := Human{"Elton", "John", 75}
	act := Action{"singing", 25, h}
	fmt.Printf("%s is %s for %d minutes\n", act.fullName(), act.kind, act.duration) //вызов метода Human из объекта типа Action
}
