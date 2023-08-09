package pokesdk

// convertNature: converts an endpoint nature to a nature
// Leaves stat structs empty as they can be populated with further API calls
func convertNature(eNature *apiNature) *Nature {
	nature := &Nature{
		ID:          eNature.ID,
		Name:        eNature.Name,
		LikesFlavor: eNature.LikesFlavor,
		HatesFlavor: eNature.HatesFlavor,
	}

	// Create empty stats for now with just the Name set
	// The full details may be fetched later
	nature.DecreasedStat = &Stat{
		Name: eNature.DecreasedStat.Name,
	}
	nature.IncreasedStat = &Stat{
		Name: eNature.IncreasedStat.Name,
	}

	if eNature.PokeathlonStatChanges != nil {
		nature.PokeathlonStatChanges = make([]struct {
			MaxChange      int "json:\"max_change\""
			PokeathlonStat struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"pokeathlon_stat\""
		}, len(eNature.PokeathlonStatChanges))
		copy(nature.PokeathlonStatChanges, eNature.PokeathlonStatChanges)
	}
	if eNature.MoveBattleStylePreferences != nil {
		nature.MoveBattleStylePreferences = make([]struct {
			LowHpPreference  int "json:\"low_hp_preference\""
			HighHpPreference int "json:\"high_hp_preference\""
			MoveBattleStyle  struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"move_battle_style\""
		}, len(eNature.MoveBattleStylePreferences))
		copy(nature.MoveBattleStylePreferences, eNature.MoveBattleStylePreferences)
	}
	if eNature.Names != nil {
		nature.Names = make([]struct {
			Name     string "json:\"name\""
			Language struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"language\""
		}, len(eNature.Names))
		copy(nature.Names, eNature.Names)
	}
	return nature
}

// convertPokemon: converts an endpoint pokemon to a Pokemon
// Leaves stat structs empty as they can be populated with further API calls
func convertPokemon(ePokemon *apiPokemon) *Pokemon {
	pokemon := &Pokemon{
		ID:                     ePokemon.ID,
		Name:                   ePokemon.Name,
		BaseExperience:         ePokemon.BaseExperience,
		Height:                 ePokemon.Height,
		IsDefault:              ePokemon.IsDefault,
		Order:                  ePokemon.Order,
		Weight:                 ePokemon.Weight,
		LocationAreaEncounters: ePokemon.LocationAreaEncounters,
		Species:                ePokemon.Species,
		Sprites:                ePokemon.Sprites,
	}

	if ePokemon.Abilities != nil {
		pokemon.Abilities = make([]struct {
			IsHidden bool "json:\"is_hidden\""
			Slot     int  "json:\"slot\""
			Ability  struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"ability\""
		}, len(ePokemon.Abilities))
		copy(pokemon.Abilities, ePokemon.Abilities)
	}
	if ePokemon.Forms != nil {
		pokemon.Forms = make([]struct {
			Name string "json:\"name\""
			URL  string "json:\"url\""
		}, len(ePokemon.Forms))
		copy(pokemon.Forms, ePokemon.Forms)
	}
	if ePokemon.GameIndices != nil {
		pokemon.GameIndices = make([]struct {
			GameIndex int "json:\"game_index\""
			Version   struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"version\""
		}, len(ePokemon.GameIndices))
		copy(pokemon.GameIndices, ePokemon.GameIndices)
	}
	if ePokemon.HeldItems != nil {
		pokemon.HeldItems = make([]struct {
			Item struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"item\""
			VersionDetails []struct {
				Rarity  int "json:\"rarity\""
				Version struct {
					Name string "json:\"name\""
					URL  string "json:\"url\""
				} "json:\"version\""
			} "json:\"version_details\""
		}, len(ePokemon.HeldItems))
		copy(pokemon.HeldItems, ePokemon.HeldItems)
		for i2 := range ePokemon.HeldItems {
			if ePokemon.HeldItems[i2].VersionDetails != nil {
				pokemon.HeldItems[i2].VersionDetails = make([]struct {
					Rarity  int "json:\"rarity\""
					Version struct {
						Name string "json:\"name\""
						URL  string "json:\"url\""
					} "json:\"version\""
				}, len(ePokemon.HeldItems[i2].VersionDetails))
				copy(pokemon.HeldItems[i2].VersionDetails, ePokemon.HeldItems[i2].VersionDetails)
			}
		}
	}
	if ePokemon.Moves != nil {
		pokemon.Moves = make([]struct {
			Move struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"move\""
			VersionGroupDetails []struct {
				LevelLearnedAt int "json:\"level_learned_at\""
				VersionGroup   struct {
					Name string "json:\"name\""
					URL  string "json:\"url\""
				} "json:\"version_group\""
				MoveLearnMethod struct {
					Name string "json:\"name\""
					URL  string "json:\"url\""
				} "json:\"move_learn_method\""
			} "json:\"version_group_details\""
		}, len(ePokemon.Moves))
		copy(pokemon.Moves, ePokemon.Moves)
		for i2 := range ePokemon.Moves {
			if ePokemon.Moves[i2].VersionGroupDetails != nil {
				pokemon.Moves[i2].VersionGroupDetails = make([]struct {
					LevelLearnedAt int "json:\"level_learned_at\""
					VersionGroup   struct {
						Name string "json:\"name\""
						URL  string "json:\"url\""
					} "json:\"version_group\""
					MoveLearnMethod struct {
						Name string "json:\"name\""
						URL  string "json:\"url\""
					} "json:\"move_learn_method\""
				}, len(ePokemon.Moves[i2].VersionGroupDetails))
				copy(pokemon.Moves[i2].VersionGroupDetails, ePokemon.Moves[i2].VersionGroupDetails)
			}
		}
	}
	if ePokemon.Stats != nil {
		pokemon.Stats = make([]struct {
			BaseStat int   "json:\"base_stat\""
			Effort   int   "json:\"effort\""
			Stat     *Stat "json:\"stat\""
		}, len(ePokemon.Stats))

		for i, stat := range ePokemon.Stats {
			pokemon.Stats[i] = struct {
				BaseStat int   "json:\"base_stat\""
				Effort   int   "json:\"effort\""
				Stat     *Stat "json:\"stat\""
			}{
				BaseStat: stat.BaseStat,
				Effort:   stat.Effort,
				// Stat struct may be populated later
				Stat: &Stat{
					Name: stat.Stat.Name,
				},
			}
		}
	}
	if ePokemon.Types != nil {
		pokemon.Types = make([]struct {
			Slot int "json:\"slot\""
			Type struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"type\""
		}, len(ePokemon.Types))
		copy(pokemon.Types, ePokemon.Types)
	}
	if ePokemon.PastTypes != nil {
		pokemon.PastTypes = make([]struct {
			Generation struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"generation\""
			Types []struct {
				Slot int "json:\"slot\""
				Type struct {
					Name string "json:\"name\""
					URL  string "json:\"url\""
				} "json:\"type\""
			} "json:\"types\""
		}, len(ePokemon.PastTypes))
		copy(pokemon.PastTypes, ePokemon.PastTypes)
		for i2 := range ePokemon.PastTypes {
			if ePokemon.PastTypes[i2].Types != nil {
				pokemon.PastTypes[i2].Types = make([]struct {
					Slot int "json:\"slot\""
					Type struct {
						Name string "json:\"name\""
						URL  string "json:\"url\""
					} "json:\"type\""
				}, len(ePokemon.PastTypes[i2].Types))
				copy(pokemon.PastTypes[i2].Types, ePokemon.PastTypes[i2].Types)
			}
		}
	}
	return pokemon
}

// convertStat: converts an endpoint stat to a stat
// Leaves nature structs empty as they can be populated with further API calls
func convertStat(eStat *apiStat) *Stat {
	stat := &Stat{
		ID:             eStat.ID,
		Name:           eStat.Name,
		GameIndex:      eStat.GameIndex,
		IsBattleOnly:   eStat.IsBattleOnly,
		AffectingMoves: eStat.AffectingMoves,
		AffectingNatures: struct {
			Increase []*Nature "json:\"increase\""
			Decrease []*Nature "json:\"decrease\""
		}{},
	}
	if eStat.AffectingMoves.Increase != nil {
		stat.AffectingMoves.Increase = make([]struct {
			Change int "json:\"change\""
			Move   struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"move\""
		}, len(eStat.AffectingMoves.Increase))
		copy(stat.AffectingMoves.Increase, eStat.AffectingMoves.Increase)
	}
	if eStat.AffectingMoves.Decrease != nil {
		stat.AffectingMoves.Decrease = make([]struct {
			Change int "json:\"change\""
			Move   struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"move\""
		}, len(eStat.AffectingMoves.Decrease))
		copy(stat.AffectingMoves.Decrease, eStat.AffectingMoves.Decrease)
	}
	if eStat.AffectingNatures.Increase != nil {
		stat.AffectingNatures.Increase = make([]*Nature, len(eStat.AffectingNatures.Increase))

		for i, nature := range eStat.AffectingNatures.Increase {
			// Nature struct may be populated later
			stat.AffectingNatures.Increase[i] = &Nature{
				Name: nature.Name,
			}
		}
	}
	if eStat.AffectingNatures.Decrease != nil {
		stat.AffectingNatures.Decrease = make([]*Nature, len(eStat.AffectingNatures.Decrease))

		for i, nature := range eStat.AffectingNatures.Decrease {
			// Nature struct may be populated later
			stat.AffectingNatures.Decrease[i] = &Nature{
				Name: nature.Name,
			}
		}
	}
	if eStat.Characteristics != nil {
		stat.Characteristics = make([]struct {
			URL string "json:\"url\""
		}, len(eStat.Characteristics))
		copy(stat.Characteristics, eStat.Characteristics)
	}
	if eStat.Names != nil {
		stat.Names = make([]struct {
			Name     string "json:\"name\""
			Language struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"language\""
		}, len(eStat.Names))
		copy(stat.Names, eStat.Names)
	}
	return stat
}
