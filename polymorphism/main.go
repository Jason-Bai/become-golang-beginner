package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

func calculateNetIncome(ic []Income) {
	var netincome int

	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}

	fmt.Printf("Net income of organization = $%d\n", netincome)
}

func polymorphism() {
	p1 := FixedBilling{projectName: "p1", biddedAmount: 5000}
	p2 := FixedBilling{projectName: "p2", biddedAmount: 10000}
	p3 := TimeAndMaterial{projectName: "p3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{p1, p2, p3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}

func main() {
	fmt.Println("1. polymorphism")
	polymorphism()
}
