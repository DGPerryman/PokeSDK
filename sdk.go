package pokesdk

import (
	"strconv"
	"time"
)

// Sdk: used to access the PokeSDK
type Sdk struct {
	// client: used to make REST requests
	client *client
	// natures: caches natures by name
	naturesByName map[string]*Nature
	// stats: caches stats by name
	statsByName map[string]*Stat
}

// NewSdk: creates a new Sdk
func NewSdk() *Sdk {
	return &Sdk{
		client:        newClient(),
		naturesByName: map[string]*Nature{},
		statsByName:   map[string]*Stat{},
	}
}

// SetTimeout: sets the timeout used for REST requests
func (s *Sdk) SetTimeout(timeout time.Duration) {
	s.client.setTimeout(timeout)
}

// GetPokemonByID: gets Pokemon details by the given ID
func (s *Sdk) GetPokemonByID(id int) (*Pokemon, error) {
	return s.getPokemon(strconv.Itoa(id))
}

// GetPokemonByName: gets Pokemon details by the given name
func (s *Sdk) GetPokemonByName(name string) (*Pokemon, error) {
	return s.getPokemon(name)
}

// getPokemon: gets Pokemon details by the given name or id
func (s *Sdk) getPokemon(nameOrID string) (*Pokemon, error) {
	// Fetch the pokemon from the endpoint
	ePokemon, err := s.client.getPokemon(nameOrID)
	if err != nil {
		return nil, err
	}

	// Convert the pokemon
	pokemon := convertPokemon(ePokemon)

	// Populate pokemon stats
	for i, eStat := range ePokemon.Stats {
		stat, err := s.GetStatByName(eStat.Stat.Name)
		if err != nil {
			return nil, err
		}
		pokemon.Stats[i].Stat = stat
	}

	return pokemon, nil
}

// GetNatureByID: gets nature details by the given ID
func (s *Sdk) GetNatureByID(id int) (*Nature, error) {
	return s.getNature(strconv.Itoa(id))
}

// GetNatureByName: gets nature details by the given name
func (s *Sdk) GetNatureByName(name string) (*Nature, error) {
	// Check to see if the nature is already cached
	cachedNature, found := s.naturesByName[name]
	if found {
		return cachedNature, nil
	}

	return s.getNature(name)
}

// getNature: gets nature details by the given name or ID
func (s *Sdk) getNature(nameOrID string) (*Nature, error) {

	// Fetch the nature from the endpoint
	eNature, err := s.client.getNature(nameOrID)
	if err != nil {
		return nil, err
	}

	// Convert the nature
	nature := convertNature(eNature)

	// Cache the nature
	s.naturesByName[nature.Name] = nature

	// Populate nature decreased stat
	decreasedStat, err := s.GetStatByName(eNature.DecreasedStat.Name)
	if err != nil {
		return nil, err
	}
	nature.DecreasedStat = decreasedStat

	// Populate nature increased stat
	increasedStat, err := s.GetStatByName(eNature.IncreasedStat.Name)
	if err != nil {
		return nil, err
	}
	nature.IncreasedStat = increasedStat

	return nature, nil
}

// GetStatByID: gets stat details by the given ID
func (s *Sdk) GetStatByID(id int) (*Stat, error) {
	return s.getStat(strconv.Itoa(id))
}

// GetStatByName: gets stat details by the given name
func (s *Sdk) GetStatByName(name string) (*Stat, error) {
	// Check to see if the stat is already cached
	cachedStat, found := s.statsByName[name]
	if found {
		return cachedStat, nil
	}

	return s.getStat(name)
}

// getStat: gets stat details by the given name or ID
func (s *Sdk) getStat(nameOrID string) (*Stat, error) {

	// Fetch the stat from the endpoint
	eStat, err := s.client.getStat(nameOrID)
	if err != nil {
		return nil, err
	}

	// Convert the stat
	stat := convertStat(eStat)

	// Cache the stat
	s.statsByName[stat.Name] = stat

	// Populate increase affecting natures
	for i, eNature := range eStat.AffectingNatures.Increase {
		nature, err := s.GetNatureByName(eNature.Name)
		if err != nil {
			return nil, err
		}
		stat.AffectingNatures.Increase[i] = nature
	}

	// Populate decrease affecting natures
	for i, eNature := range eStat.AffectingNatures.Decrease {
		nature, err := s.GetNatureByName(eNature.Name)
		if err != nil {
			return nil, err
		}
		stat.AffectingNatures.Decrease[i] = nature
	}

	return stat, nil
}
