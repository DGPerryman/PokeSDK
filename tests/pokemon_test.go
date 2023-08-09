package pokesdk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pokesdk "github.com/DGPerryman/PokeSDK"
)

func TestGetPokemonByName(t *testing.T) {
	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect no error
	pokemon, err := sdk.GetPokemonByName("pikachu")
	assert.NoError(t, err)
	assert.NotNil(t, pokemon)

	// Check the details are returned correctly
	assert.Equal(t, 25, pokemon.ID)
	assert.Equal(t, "pikachu", pokemon.Name)
	assert.Equal(t, 112, pokemon.BaseExperience)
	assert.Equal(t, 4, pokemon.Height)
	assert.Equal(t, true, pokemon.IsDefault)
	assert.Equal(t, 35, pokemon.Order)
	assert.Equal(t, 60, pokemon.Weight)

	// Check stats are populated
	assert.Equal(t, 6, len(pokemon.Stats))
	assert.Equal(t, "attack", pokemon.Stats[1].Stat.Name)
	assert.Equal(t, 2, pokemon.Stats[1].Stat.ID)
	assert.Equal(t, 2, pokemon.Stats[1].Stat.GameIndex)
	assert.Equal(t, false, pokemon.Stats[1].Stat.IsBattleOnly)

	// Check affected nature within stat is populated
	assert.Equal(t, 2, pokemon.Stats[1].Stat.AffectingNatures.Decrease[0].ID)
	assert.Equal(t, "bold", pokemon.Stats[1].Stat.AffectingNatures.Decrease[0].Name)
}

func TestGetPokemonByID(t *testing.T) {

	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect no error
	pokemon, err := sdk.GetPokemonByID(25)
	assert.NoError(t, err)
	assert.NotNil(t, pokemon)

	// Check the details are returned correctly
	assert.Equal(t, 25, pokemon.ID)
	assert.Equal(t, "pikachu", pokemon.Name)
	assert.Equal(t, 112, pokemon.BaseExperience)
	assert.Equal(t, 4, pokemon.Height)
	assert.Equal(t, true, pokemon.IsDefault)
	assert.Equal(t, 35, pokemon.Order)
	assert.Equal(t, 60, pokemon.Weight)

	// Check stats are populated
	assert.Equal(t, 6, len(pokemon.Stats))
	assert.Equal(t, "attack", pokemon.Stats[1].Stat.Name)
	assert.Equal(t, 2, pokemon.Stats[1].Stat.ID)
	assert.Equal(t, 2, pokemon.Stats[1].Stat.GameIndex)
	assert.Equal(t, false, pokemon.Stats[1].Stat.IsBattleOnly)

	// Check affected nature within stat is populated
	assert.Equal(t, 2, pokemon.Stats[1].Stat.AffectingNatures.Decrease[0].ID)
	assert.Equal(t, "bold", pokemon.Stats[1].Stat.AffectingNatures.Decrease[0].Name)
}

func TestPokemonNotFound(t *testing.T) {

	// Create SDK
	sdk := pokesdk.NewSdk()

	// Call and expect an error
	pokemon, err := sdk.GetPokemonByName("nonexistant")
	assert.Error(t, err)
	assert.Nil(t, pokemon)

	// Check the error details are returned correctly
	assert.ErrorIs(t, err, pokesdk.ErrNotFound)
}
