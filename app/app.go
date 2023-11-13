package app

import (
	"Pz5/app/model"
	magazzine "Pz5/pkg/magazine"
	"Pz5/pkg/store/dataJson"
	"fmt"
)

func Run() error {
	store := dataJson.New("data.json")
	// Створення екземпляра MonthlyMagazine за допомогою експортованого конструктора

	//NewMonthlyMagazine := model.NewMonthlyMagazine("Monthly Journal", "Publisher XYZ", "2023-11-10")
	//
	//// Створення екземпляра нерегулярного журналу за допомогою конструктора
	//irregularMagazineInstance := model.NewIrregularMagazine("Нерегулярний журнал", "Видавництво XYZ", "15.11.2023")
	//
	//// Використання інтерфейсу для обробки нерегулярного журналу
	//err := processPeriodical(irregularMagazineInstance)
	//if err != nil {
	//	return err
	//}
	//err = processPeriodical(NewMonthlyMagazine)
	//if err != nil {
	//	return err
	//}

	var irregularMagazineInstanceSlice []magazzine.Magazines
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewIrregularMagazine("Нерегулярний журнал 2", "Видавництво ZXC ", "15.11.2023"))
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewIrregularMagazine("Нерегулярний журнал 2", "Видавництво ZXC ", "15.11.2023"))
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewMonthlyMagazine("Нерегулярний журнал", "Видавництво XYZ", "15.11.2023"))
	//irregularMagazineInstanceSlice[0].PrintInfo()

	err := processSlicePeriodical(irregularMagazineInstanceSlice)
	if err != nil {
		return err
	}
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

	if err != nil {
		return err
	}

	// Припустимо, що у вас є екземпляр MonthlyMagazine та IrregularMagazine
	//monthlyMagazine := model.NewMonthlyMagazine("Monthly Journal", "Publisher XYZ", "2023-11-10")
	//irregularMagazineInstances := model.NewIrregularMagazine("Нерегулярний журнал", "Видавництво XYZ", "15.11.2023")

	// Створення колекції з цих екземплярів
	//collection := model.NewCollectionFromMagazines(monthlyMagazine)
	//collection.AddMagazines(irregularMagazineInstances)
	////// Виведення елементів IrregularMagazine
	//fmt.Println("df")
	//err = processPeriodical(collection)

	//err = processSlicePeriodical(df)
	if err != nil {
		return err
	}
	// Збереження користувачів у JSON-файл
	err = store.Collection().SaveJSON(irregularMagazineInstanceSlice)
	if err != nil {
		return err
	}
	load, err := store.Collection().LoadJSON()

	err = processSlicePeriodicalComparer(load.AddDefaultMagazines())
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	return nil
}

func processSlicePeriodical(magazines []magazzine.Magazines) error {
	for _, mag := range magazines {
		mag.PrintInfo()
	}
	return nil
}

func processSlicePeriodicalComparer(magazines []magazzine.Magazines) error {
	n := len(magazines)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if magazines[i].MagazinesEqual(magazines[j]) {
				fmt.Println("Журнали рівні\n")
				err := processPeriodical(magazines[i])
				if err != nil {
					return err
				}
				err = processPeriodical(magazines[j])
				if err != nil {
					return err
				}

			} else {
				fmt.Println("Журнали не рівні\n")
				err := processPeriodical(magazines[i])
				if err != nil {
					return err
				}
				err = processPeriodical(magazines[j])
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
func processPeriodical(p magazzine.Periodicals) error {
	fmt.Printf("Обробка видання - Частота виходу: %s\n", p.GetFrequency())
	p.PrintInfo()
	return nil
}
