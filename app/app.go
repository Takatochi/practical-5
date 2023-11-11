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
	if err != nil {
		return err
	}
	err = processPeriodical(NewMonthlyMagazine)
	if err != nil {
		return err
	}
	var irregularMagazineInstanceSlice []store.Magazines
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewIrregularMagazine("Нерегулярний журнал 2", "Видавництво ZXC ", "15.11.2023"))
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewIrregularMagazine(" журнал", "Видавництво XZB", "15.11.2022"))
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewMonthlyMagazine("Нерегулярний журнал", "Видавництво XYZ", "15.11.2023"))
	//irregularMagazineInstanceSlice[0].PrintInfo()

	//err = processSlicePeriodical(irregularMagazineInstanceSlice)
	//if err != nil {
	//	return err
	//}
	err = processSlicePeriodicalComparer(irregularMagazineInstanceSlice)
	if err != nil {
		return err
	}
	// Викликаємо метод MagazinesEqual для порівняння об'єктів
	//if irregularMagazineInstance.MagazinesEqual(irregularMagazineInstance) {
	//	fmt.Println("Журнали рівні")
	//} else {
	//	fmt.Println("Журнали не рівні")
	//}

	return nil
}

func processPeriodical(p store.Periodicals) error {
	fmt.Printf("Обробка видання - Частота виходу: %s\n", p.GetFrequency())
	p.PrintInfo()

	return nil
}

//func processSlicePeriodical(magazines []store.Magazines) error {
//	for _, mag := range magazines {
//		mag.PrintInfo()
//	}
//	return nil
//}

func processSlicePeriodicalComparer(magazines []store.Magazines) error {
	n := len(magazines)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if magazines[i].MagazinesEqual(magazines[j]) {
				fmt.Println("Журнали рівні\n")
				magazines[i].PrintInfo()
				magazines[j].PrintInfo()
			} else {
				fmt.Println("Журнали не рівні\n")
				magazines[i].PrintInfo()
				magazines[j].PrintInfo()
			}
		}
	}

	return nil
}
