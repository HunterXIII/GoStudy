package main

import (
	"fmt"
)

type PepeSchnele struct {
	Speed    int
	Charisma int
	Wisdom   int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{
		Speed:    speed,
		Charisma: charisma,
		Wisdom:   wisdom,
	}
}

func (p *PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func main() {
	pepe1 := NewPepeSchnele(5, 7, 6)
	pepe2 := NewPepeSchnele(8, 4, 5)

	fmt.Printf(
		"Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d\n",
		pepe1.Speed, pepe1.Charisma, pepe1.Wisdom, pepe1.GetRating(),
	)

	fmt.Printf(
		"Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d\n",
		pepe2.Speed, pepe2.Charisma, pepe2.Wisdom, pepe2.GetRating(),
	)

	if pepe1.GetRating() > pepe2.GetRating() {
		fmt.Println("У первого Пепе рейтинг выше")
	} else if pepe1.GetRating() < pepe2.GetRating() {
		fmt.Println("У второго Пепе рейтинг выше")
	} else {
		fmt.Println("Рейтинги равны")
	}
}
