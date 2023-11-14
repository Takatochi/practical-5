package store_test

import (
	"Pz5/app/model"
	mag "Pz5/app/model/magazine"
	"Pz5/pkg/store"
	"Pz5/pkg/store/dataJson"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestMagazineRepository_SaveAndLoadJSON(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := ioutil.TempFile("", "test_magazine.json")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Create a sample data to save
	NewMonthlyMagazine := mag.NewMonthlyMagazine("Monthly Journal", "Publisher XYZ", "2023-11-10")
	// Створення екземпляра нерегулярного журналу за допомогою конструктора
	irregularMagazineInstance := mag.NewIrregularMagazine("Нерегулярний журнал", "Видавництво XYZ", "15.11.2023")

	collection := mag.NewCollectionFromMagazines(NewMonthlyMagazine, irregularMagazineInstance)
	magazinesToSave := collection.AddDefaultMagazines()
	// Create Store and MagazineRepository
	store := dataJson.New(tempFile.Name())

	// Save the sample data
	err = store.Collection().SaveJSON(magazinesToSave)
	assert.NoError(t, err)

	// Load the data back
	loadedCollection, err := store.Collection().LoadJSON()
	assert.NoError(t, err)

	// Verify that the loaded data matches the original data
	assert.Equal(t, len(loadedCollection.MonthlyMagazine), len(collection.MonthlyMagazine))
	assert.Equal(t, len(loadedCollection.IrregularMagazine), len(collection.IrregularMagazine))
	// Compare Monthly Magazines
	for i, expectedMag := range collection.MonthlyMagazine {
		assert.Equal(t, expectedMag, loadedCollection.MonthlyMagazine[i])
	}

	// Compare Irregular Magazines
	for i, expectedMag := range collection.IrregularMagazine {
		assert.Equal(t, expectedMag, loadedCollection.IrregularMagazine[i])
	}
}

func TestStore_Collection(t *testing.T) {
	// Create a sample Store
	dbFileName := "test_db.json"
	store := dataJson.New(dbFileName)

	// Get the MagazineRepository from the Store
	magRepo := store.Collection()

	// Ensure that the returned MagazineRepository is not nil
	assert.NotNil(t, magRepo)
}

func BenchmarkMagazineRepository_SaveJSON(b *testing.B) {
	tempFile, err := ioutil.TempFile("", "bench_magazine.json")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	magazinesToSave := generateTestMagazines(1000) // Assuming you have a function to generate test magazines
	mRepo := createTestMagazineRepository(tempFile.Name())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := mRepo.SaveJSON(magazinesToSave)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMagazineRepository_LoadJSON(b *testing.B) {
	tempFile, err := ioutil.TempFile("", "bench_magazine.json")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	magazinesToSave := generateTestMagazines(1000) // Assuming you have a function to generate test magazines
	mRepo := createTestMagazineRepository(tempFile.Name())
	err = mRepo.SaveJSON(magazinesToSave)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := mRepo.LoadJSON()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func generateTestMagazines(count int) []model.Magazines {
	var testMagazines []model.Magazines

	// Implement your logic to generate test magazines here
	for i := 0; i < count; i++ {
		// For simplicity, create instances of MonthlyMagazine and IrregularMagazine
		monthlyMagazine := mag.NewMonthlyMagazine("Monthly Journal", "Publisher XYZ", "2023-11-10")
		irregularMagazine := mag.NewIrregularMagazine("Irregular Journal", "Publisher ABC", "2023-11-15")

		// Append them to the slice
		testMagazines = append(testMagazines, monthlyMagazine, irregularMagazine)
	}

	return testMagazines
}

func createTestMagazineRepository(fileName string) store.MagazineRepository {
	store := dataJson.New(fileName)
	return store.Collection()
}
