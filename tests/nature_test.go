package pokesdk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pokesdk "github.com/DGPerryman/PokeSDK"
)

func TestGetNatureByName(t *testing.T) {

	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect no error
	nature, err := sdk.GetNatureByName("bold")
	assert.NoError(t, err)
	assert.NotNil(t, nature)

	// Check the details are returned correctly
	assert.Equal(t, 2, nature.ID)
	assert.Equal(t, "bold", nature.Name)

	// Check stats of the nature are populated
	assert.Equal(t, 3, nature.IncreasedStat.ID)
	assert.Equal(t, "defense", nature.IncreasedStat.Name)
	assert.Equal(t, 3, nature.IncreasedStat.GameIndex)
	assert.Equal(t, false, nature.IncreasedStat.IsBattleOnly)

	// Check AffectingNatures of the stats are populated
	assert.Equal(t, 6, nature.IncreasedStat.AffectingNatures.Decrease[0].ID)
	assert.Equal(t, "lonely", nature.IncreasedStat.AffectingNatures.Decrease[0].Name)
}

func TestGetNatureByID(t *testing.T) {

	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect no error
	nature, err := sdk.GetNatureByID(2)
	assert.NoError(t, err)
	assert.NotNil(t, nature)

	// Check the details are returned correctly
	assert.Equal(t, 2, nature.ID)
	assert.Equal(t, "bold", nature.Name)

	// Check stats of the nature are populated
	assert.Equal(t, 3, nature.IncreasedStat.ID)
	assert.Equal(t, "defense", nature.IncreasedStat.Name)
	assert.Equal(t, 3, nature.IncreasedStat.GameIndex)
	assert.Equal(t, false, nature.IncreasedStat.IsBattleOnly)

	// Check AffectingNatures of the stats are populated
	assert.Equal(t, 6, nature.IncreasedStat.AffectingNatures.Decrease[0].ID)
	assert.Equal(t, "lonely", nature.IncreasedStat.AffectingNatures.Decrease[0].Name)
}

func TestNatureNotFound(t *testing.T) {

	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect an error
	nature, err := sdk.GetNatureByName("nonexistant")
	assert.Error(t, err)
	assert.Nil(t, nature)

	// Check the error details are returned correctly
	assert.ErrorIs(t, err, pokesdk.ErrNotFound)
}
