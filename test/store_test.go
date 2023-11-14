package store_test

import (
	mag "Pz5/app/model/magazine"
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
