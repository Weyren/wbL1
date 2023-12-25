package main

import (
	"fmt"
)

//Айти компании из Сальвадора потребовался новый иностранный тимлид,
//но у них закончились баксы - все вложили в биток

// структура для адаптции - иностранный тимлид
type TeamLead struct{}

// функция работать(готов только за доллары), не реализует интерфейс Worker
func (t *TeamLead) workForDollars(salary int) {
	fmt.Println("I'm working for", salary, "dollars")
}

// интерфейс работника компании
type Worker interface {
	WorkForBTC(btc float64) //Налогвая не дремлет. Требует указывать зарплату в битках
}

type Junior struct{}

// типичный программист работает за биткоины(официальная валюта Сальвадора), реализует интерфейс Worker
func (w *Junior) WorkForBTC(btc float64) {
	fmt.Println("I'm working for", btc, "BTC")

}

// адаптер для тимлида
type TeamLeadAdapter struct {
	tl *TeamLead
}

// теперь тимлид работает в сальвадорской компании и реализует интерфейс Worker
func (tla *TeamLeadAdapter) WorkForBTC(btc float64) {
	tla.tl.workForDollars(convert(btc))
}

// компания продает биток и покупат баксы
func convert(btc float64) int {
	return int(btc * 41000)
}

func main() {
	teamlead := &TeamLeadAdapter{&TeamLead{}} //создали структуру адаптированного тимлида
	junior1 := &Junior{}
	junior2 := &Junior{}
	workers := []Worker{teamlead, junior1, junior2} //все работники компании

	for _, worker := range workers {
		if worker == teamlead {
			worker.WorkForBTC(0.2)
		} else {
			worker.WorkForBTC(0.1)
		}
	}
}
