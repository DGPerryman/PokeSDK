package pokesdk_test

import (
	"testing"

	pokesdk "github.com/DGPerryman/PokeSDK"
	"github.com/stretchr/testify/assert"
)

func TestGetStatByName(t *testing.T) {
	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect no error
	stat, err := sdk.GetStatByName("attack")
	assert.NoError(t, err)
	assert.NotNil(t, stat)

	// Check the details are returned correctly
	assert.Equal(t, 2, stat.ID)
	assert.Equal(t, "attack", stat.Name)
	assert.Equal(t, 2, stat.GameIndex)
	assert.Equal(t, false, stat.IsBattleOnly)

	// Check AffectingNatures are populated
	assert.Equal(t, 4, len(stat.AffectingNatures.Decrease))
	assert.Equal(t, 2, stat.AffectingNatures.Decrease[0].ID)
	assert.Equal(t, "bold", stat.AffectingNatures.Decrease[0].Name)

	// Check stats of AffectingNatures are populated
	assert.Equal(t, 3, stat.AffectingNatures.Decrease[0].IncreasedStat.ID)
	assert.Equal(t, "defense", stat.AffectingNatures.Decrease[0].IncreasedStat.Name)
	assert.Equal(t, 3, stat.AffectingNatures.Decrease[0].IncreasedStat.GameIndex)
	assert.Equal(t, false, stat.AffectingNatures.Decrease[0].IncreasedStat.IsBattleOnly)
}

func TestGetStatByID(t *testing.T) {

	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect no error
	stat, err := sdk.GetStatByID(2)
	assert.NoError(t, err)
	assert.NotNil(t, stat)

	// Check the details are returned correctly
	assert.Equal(t, 2, stat.ID)
	assert.Equal(t, "attack", stat.Name)
	assert.Equal(t, 2, stat.GameIndex)
	assert.Equal(t, false, stat.IsBattleOnly)

	// Check AffectingNatures are populated
	assert.Equal(t, 4, len(stat.AffectingNatures.Decrease))
	assert.Equal(t, 2, stat.AffectingNatures.Decrease[0].ID)
	assert.Equal(t, "bold", stat.AffectingNatures.Decrease[0].Name)

	// Check stats of AffectingNatures are populated
	assert.Equal(t, 3, stat.AffectingNatures.Decrease[0].IncreasedStat.ID)
	assert.Equal(t, "defense", stat.AffectingNatures.Decrease[0].IncreasedStat.Name)
	assert.Equal(t, 3, stat.AffectingNatures.Decrease[0].IncreasedStat.GameIndex)
	assert.Equal(t, false, stat.AffectingNatures.Decrease[0].IncreasedStat.IsBattleOnly)
}

func TestStatNotFound(t *testing.T) {

	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect an error
	stat, err := sdk.GetStatByName("nonexistant")
	assert.Error(t, err)
	assert.Nil(t, stat)

	// Check the error details are returned correctly
	assert.ErrorIs(t, err, pokesdk.ErrNotFound)
}
