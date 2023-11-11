package app

import (
	"Pz5/app/model"
	"Pz5/pkg/store"
	"fmt"
)

func Run() error {

	// Створення екземпляра MonthlyMagazine за допомогою експортованого конструктора
	NewMonthlyMagazine := model.NewMonthlyMagazine("Monthly Journal", "Publisher XYZ", "2023-11-10")

	// Створення екземпляра нерегулярного журналу за допомогою конструктора
	irregularMagazineInstance := model.NewIrregularMagazine("Нерегулярний журнал", "Видавництво XYZ", "15.11.2023")

	// Використання інтерфейсу для обробки нерегулярного журналу
	err := processPeriodical(irregularMagazineInstance)

	err = processPeriodical(NewMonthlyMagazine)
	if err != nil {
		return err
	}
	var irregularMagazineInstanceSlice []store.Magazine
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewIrregularMagazine("Нерегулярний журнал 2", "Видавництво ZXC ", "15.11.2023"))
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewIrregularMagazine(" журнал", "Видавництво XZB", "15.11.2022"))
	//irregularMagazineInstanceSlice[0].PrintInfo()

	err = processSlicePeriodical(irregularMagazineInstanceSlice)
	if err != nil {
		return err
	}
	return nil
}

func processPeriodical(p store.Periodicals) error {
	fmt.Printf("Обробка видання - Частота виходу: %s\n", p.GetFrequency())
	p.PrintInfo()

	return nil
}

func processSlicePeriodical(magazines []store.Magazine) error {
	for _, mag := range magazines {
		mag.PrintInfo()

	}
	return nil
}
