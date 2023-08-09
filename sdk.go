package pokesdk

import (
	"strconv"
	"time"
)

// Sdk: used to access the PokeSDK
type Sdk struct {
	// client: used to make REST requests
	client *client
	// natures: caches natures returned
	natures map[string]*Nature
	// stats: caches stats returned
	stats map[string]*Stat
}

// NewSdk: creates a new Sdk
func NewSdk() *Sdk {
	return &Sdk{
		client:  newClient(),
		natures: map[string]*Nature{},
		stats:   map[string]*Stat{},
	}
}

// SetTimeout: sets the timeout used for REST requests
func (s *Sdk) SetTimeout(timeout time.Duration) {
	s.client.setTimeout(timeout)
}

// GetPokemonByID: gets Pokemon details by the given ID
func (s *Sdk) GetPokemonByID(id int) (*Pokemon, error) {
	return s.GetPokemonByName(strconv.Itoa(id))
}

// GetPokemonByName: gets Pokemon details by the given name
func (s *Sdk) GetPokemonByName(name string) (*Pokemon, error) {
	// Fetch the pokemon from the endpoint
	ePokemon, err := s.client.getPokemon(name)
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
	return s.GetNatureByName(strconv.Itoa(id))
}

// GetNatureByName: gets nature details by the given name
func (s *Sdk) GetNatureByName(name string) (*Nature, error) {
	// Check to see if the nature is already cached
	cachedNature, found := s.natures[name]
	if found {
		return cachedNature, nil
	}

	// Fetch the nature from the endpoint
	eNature, err := s.client.getNature(name)
	if err != nil {
		return nil, err
	}

	// Convert the nature
	nature := convertNature(eNature)

	// Cache the nature
	s.natures[nature.Name] = nature

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
	return s.GetStatByName(strconv.Itoa(id))
}

// GetStatByName: gets stat details by the given name
func (s *Sdk) GetStatByName(name string) (*Stat, error) {
	// Check to see if the stat is already cached
	cachedStat, found := s.stats[name]
	if found {
		return cachedStat, nil
	}

	// Fetch the stat from the endpoint
	eStat, err := s.client.getStat(name)
	if err != nil {
		return nil, err
	}

	// Convert the stat
	stat := convertStat(eStat)

	// Cache the stat
	s.stats[stat.Name] = stat

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
