package main

import "fmt"

//адаптор измерителя роста в футах и дюймах к метрической системе

type getHeight interface { //интерфейс,который требуется системе
	getHeightCM() float64
}

type height struct { //измаритель роста в футах и дюймах
	feet   int
	inches int
}

func (h *height) getFeet() int { //выдает футы роста
	return h.feet
}
func (h *height) getInches() int { //выдает дюймы роста
	return h.inches
}

type adapter struct { //адаптер, который преобразует в см
	hgt *height
}

func (a adapter) getHeightCM() float64 { //адаптация - перевод из футов и дюймов в см
	foot := a.hgt.getFeet()
	inc := a.hgt.getInches()
	inc += foot * 12
	return float64(inc) * 2.54
}

func main() {
	hgt := height{6, 3}
	var adp getHeight = adapter{&hgt} //интерфейс, у которого есть метод получения роста
	fmt.Println(adp.getHeightCM())
}
