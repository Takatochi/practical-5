package magazine_test

import (
	"Pz5/app/model/magazine"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIrregularMagazine_MagazinesEqual(t *testing.T) {
	// Створити два нерегулярних журнали з однаковими даними
	mag1 := magazine.NewIrregularMagazine("Назва", "Видавець", "2023-11-10")
	mag2 := magazine.NewIrregularMagazine("Назва", "Видавець", "2023-11-10")

	// Перевірити, чи журнали однакові
	assert.True(t, mag1.MagazinesEqual(mag2))
}

func TestMonthlyMagazine_MagazinesEqual(t *testing.T) {
	// Створити два щомісячних журнали з однаковими даними
	mag1 := magazine.NewMonthlyMagazine("Назва", "Видавець", "2023-11-10")
	mag2 := magazine.NewMonthlyMagazine("Назва", "Видавець", "2023-11-10")

	// Перевірити, чи журнали однакові
	assert.True(t, mag1.MagazinesEqual(mag2))
}

func TestCollection_AddMagazines(t *testing.T) {
	// Створити колекцію і додати два журнали
	coll := magazine.Collection{}
	irregularMag := magazine.NewIrregularMagazine("Нерегулярний Журнал", "Видавець", "2023-11-10")
	monthlyMag := magazine.NewMonthlyMagazine("Щомісячний Журнал", "Видавець", "2023-11-10")

	coll.AddMagazines(irregularMag, monthlyMag)

	// Перевірити, чи колекція містить додані журнали
	assert.Equal(t, 1, len(coll.IrregularMagazine))
	assert.Equal(t, 1, len(coll.MonthlyMagazine))
}

func TestCollection_AddDefaultMagazines(t *testing.T) {
	// Створити колекцію і додати два журнали
	coll := magazine.Collection{}
	irregularMag := magazine.NewIrregularMagazine("Нерегулярний Журнал", "Видавець", "2023-11-10")
	monthlyMag := magazine.NewMonthlyMagazine("Щомісячний Журнал", "Видавець", "2023-11-10")

	coll.AddMagazines(irregularMag, monthlyMag)

	// Додати типові журнали
	defaultMagazines := coll.AddDefaultMagazines()

	// Перевірити, чи типові журнали відповідають доданим журналам
	assert.Equal(t, 2, len(defaultMagazines))
	assert.Equal(t, irregularMag, defaultMagazines[0])
	assert.Equal(t, monthlyMag, defaultMagazines[1])
}

func TestMagazine_PrintInfo(t *testing.T) {
	// Створити журнал і зафіксувати виведення на екран
	mag := magazine.NewIrregularMagazine("Нерегулярний Журнал", "Видавець", "2023-11-10")
	assert.NotPanics(t, func() {
		mag.PrintInfo()
	})
}

func TestMagazine_GetFrequency(t *testing.T) {
	// Створити журнал і отримати його частоту виходу
	mag := magazine.NewIrregularMagazine("Нерегулярний Журнал", "Видавець", "2023-11-10")
	frequency := mag.GetFrequency()

	// Перевірити, чи частота виходу відповідає очікуваній
	assert.Equal(t, "Раз на кілька місяців", frequency)
}
