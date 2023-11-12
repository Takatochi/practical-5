package app

import (
	"Pz5/app/model"
	magazzine "Pz5/pkg/magazine"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func Run() error {

	// Створення екземпляра MonthlyMagazine за допомогою експортованого конструктора
	NewMonthlyMagazine := model.NewMonthlyMagazine("Monthly Journal", "Publisher XYZ", "2023-11-10")

	// Створення екземпляра нерегулярного журналу за допомогою конструктора
	irregularMagazineInstance := model.NewIrregularMagazine("Нерегулярний журнал", "Видавництво XYZ", "15.11.2023")
	//coll := model.Collection{IrregularMagazine: &irregularMagazineInstance, MonthlyMagazine: &NewMonthlyMagazine}
	// Використання інтерфейсу для обробки нерегулярного журналу
	err := processPeriodical(irregularMagazineInstance)
	if err != nil {
		return err
	}
	err = processPeriodical(NewMonthlyMagazine)
	if err != nil {
		return err
	}

	var irregularMagazineInstanceSlice []magazzine.Magazines
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewIrregularMagazine("Нерегулярний журнал 2", "Видавництво ZXC ", "15.11.2023"))
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewIrregularMagazine(" журнал", "Видавництво XZB", "15.11.2022"))
	irregularMagazineInstanceSlice = append(irregularMagazineInstanceSlice, model.NewMonthlyMagazine("Нерегулярний журнал", "Видавництво XYZ", "15.11.2023"))
	//irregularMagazineInstanceSlice[0].PrintInfo()

	err = processSlicePeriodical(irregularMagazineInstanceSlice)
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
	// Збереження користувачів у JSON-файл

	err = SaveUsers(irregularMagazineInstanceSlice, "data.json")

	if err != nil {
		log.Fatal(err)
	}
	df, err := LoadUsers("data.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(df[0].Description)
	//err = processSlicePeriodical(df)
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

// SaveUsers функція для збереження користувачів у JSON-файл
func SaveUsers(df []magazzine.Magazines, filename string) error {
	data, err := json.MarshalIndent(df, " ", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("magazines saved to %s\n", filename)
	return nil
}

// LoadUsers функція для завантаження користувачів з JSON-файлу
func LoadUsers(filename string) ([]model.Magazines, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var df []model.Magazines
	err = json.Unmarshal(data, &df)
	if err != nil {
		return nil, err
	}

	return df, nil
}
