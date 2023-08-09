package pokesdk

// convertNature: converts an API nature to a nature
// Leaves stat structs empty as they can be populated with further API calls
func convertNature(aNature *apiNature) *Nature {
	nature := &Nature{
		ID:          aNature.ID,
		Name:        aNature.Name,
		LikesFlavor: aNature.LikesFlavor,
		HatesFlavor: aNature.HatesFlavor,
	}

	// Create empty stats for now with just the Name set
	// The full details may be fetched later
	nature.DecreasedStat = &Stat{
		Name: aNature.DecreasedStat.Name,
	}
	nature.IncreasedStat = &Stat{
		Name: aNature.IncreasedStat.Name,
	}

	if aNature.PokeathlonStatChanges != nil {
		nature.PokeathlonStatChanges = make([]struct {
			MaxChange      int "json:\"max_change\""
			PokeathlonStat struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"pokeathlon_stat\""
		}, len(aNature.PokeathlonStatChanges))
		copy(nature.PokeathlonStatChanges, aNature.PokeathlonStatChanges)
	}
	if aNature.MoveBattleStylePreferences != nil {
		nature.MoveBattleStylePreferences = make([]struct {
			LowHpPreference  int "json:\"low_hp_preference\""
			HighHpPreference int "json:\"high_hp_preference\""
			MoveBattleStyle  struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"move_battle_style\""
		}, len(aNature.MoveBattleStylePreferences))
		copy(nature.MoveBattleStylePreferences, aNature.MoveBattleStylePreferences)
	}
	if aNature.Names != nil {
		nature.Names = make([]struct {
			Name     string "json:\"name\""
			Language struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"language\""
		}, len(aNature.Names))
		copy(nature.Names, aNature.Names)
	}
	return nature
}

// convertPokemon: converts an API pokemon to a pokemon
// Leaves stat structs empty as they can be populated with further API calls
func convertPokemon(aPokemon *apiPokemon) *Pokemon {
	pokemon := &Pokemon{
		ID:                     aPokemon.ID,
		Name:                   aPokemon.Name,
		BaseExperience:         aPokemon.BaseExperience,
		Height:                 aPokemon.Height,
		IsDefault:              aPokemon.IsDefault,
		Order:                  aPokemon.Order,
		Weight:                 aPokemon.Weight,
		LocationAreaEncounters: aPokemon.LocationAreaEncounters,
		Species:                aPokemon.Species,
		Sprites:                aPokemon.Sprites,
	}

	if aPokemon.Abilities != nil {
		pokemon.Abilities = make([]struct {
			IsHidden bool "json:\"is_hidden\""
			Slot     int  "json:\"slot\""
			Ability  struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"ability\""
		}, len(aPokemon.Abilities))
		copy(pokemon.Abilities, aPokemon.Abilities)
	}
	if aPokemon.Forms != nil {
		pokemon.Forms = make([]struct {
			Name string "json:\"name\""
			URL  string "json:\"url\""
		}, len(aPokemon.Forms))
		copy(pokemon.Forms, aPokemon.Forms)
	}
	if aPokemon.GameIndices != nil {
		pokemon.GameIndices = make([]struct {
			GameIndex int "json:\"game_index\""
			Version   struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"version\""
		}, len(aPokemon.GameIndices))
		copy(pokemon.GameIndices, aPokemon.GameIndices)
	}
	if aPokemon.HeldItems != nil {
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
		}, len(aPokemon.HeldItems))
		copy(pokemon.HeldItems, aPokemon.HeldItems)
		for i2 := range aPokemon.HeldItems {
			if aPokemon.HeldItems[i2].VersionDetails != nil {
				pokemon.HeldItems[i2].VersionDetails = make([]struct {
					Rarity  int "json:\"rarity\""
					Version struct {
						Name string "json:\"name\""
						URL  string "json:\"url\""
					} "json:\"version\""
				}, len(aPokemon.HeldItems[i2].VersionDetails))
				copy(pokemon.HeldItems[i2].VersionDetails, aPokemon.HeldItems[i2].VersionDetails)
			}
		}
	}
	if aPokemon.Moves != nil {
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
		}, len(aPokemon.Moves))
		copy(pokemon.Moves, aPokemon.Moves)
		for i2 := range aPokemon.Moves {
			if aPokemon.Moves[i2].VersionGroupDetails != nil {
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
				}, len(aPokemon.Moves[i2].VersionGroupDetails))
				copy(pokemon.Moves[i2].VersionGroupDetails, aPokemon.Moves[i2].VersionGroupDetails)
			}
		}
	}
	if aPokemon.Stats != nil {
		pokemon.Stats = make([]struct {
			BaseStat int   "json:\"base_stat\""
			Effort   int   "json:\"effort\""
			Stat     *Stat "json:\"stat\""
		}, len(aPokemon.Stats))

		for i, stat := range aPokemon.Stats {
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
	if aPokemon.Types != nil {
		pokemon.Types = make([]struct {
			Slot int "json:\"slot\""
			Type struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"type\""
		}, len(aPokemon.Types))
		copy(pokemon.Types, aPokemon.Types)
	}
	if aPokemon.PastTypes != nil {
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
		}, len(aPokemon.PastTypes))
		copy(pokemon.PastTypes, aPokemon.PastTypes)
		for i2 := range aPokemon.PastTypes {
			if aPokemon.PastTypes[i2].Types != nil {
				pokemon.PastTypes[i2].Types = make([]struct {
					Slot int "json:\"slot\""
					Type struct {
						Name string "json:\"name\""
						URL  string "json:\"url\""
					} "json:\"type\""
				}, len(aPokemon.PastTypes[i2].Types))
				copy(pokemon.PastTypes[i2].Types, aPokemon.PastTypes[i2].Types)
			}
		}
	}
	return pokemon
}

// convertStat: converts an API stat to a stat
// Leaves nature structs empty as they can be populated with further API calls
func convertStat(aStat *apiStat) *Stat {
	stat := &Stat{
		ID:             aStat.ID,
		Name:           aStat.Name,
		GameIndex:      aStat.GameIndex,
		IsBattleOnly:   aStat.IsBattleOnly,
		AffectingMoves: aStat.AffectingMoves,
		AffectingNatures: struct {
			Increase []*Nature "json:\"increase\""
			Decrease []*Nature "json:\"decrease\""
		}{},
	}
	if aStat.AffectingMoves.Increase != nil {
		stat.AffectingMoves.Increase = make([]struct {
			Change int "json:\"change\""
			Move   struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"move\""
		}, len(aStat.AffectingMoves.Increase))
		copy(stat.AffectingMoves.Increase, aStat.AffectingMoves.Increase)
	}
	if aStat.AffectingMoves.Decrease != nil {
		stat.AffectingMoves.Decrease = make([]struct {
			Change int "json:\"change\""
			Move   struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"move\""
		}, len(aStat.AffectingMoves.Decrease))
		copy(stat.AffectingMoves.Decrease, aStat.AffectingMoves.Decrease)
	}
	if aStat.AffectingNatures.Increase != nil {
		stat.AffectingNatures.Increase = make([]*Nature, len(aStat.AffectingNatures.Increase))

		for i, nature := range aStat.AffectingNatures.Increase {
			// Nature struct may be populated later
			stat.AffectingNatures.Increase[i] = &Nature{
				Name: nature.Name,
			}
		}
	}
	if aStat.AffectingNatures.Decrease != nil {
		stat.AffectingNatures.Decrease = make([]*Nature, len(aStat.AffectingNatures.Decrease))

		for i, nature := range aStat.AffectingNatures.Decrease {
			// Nature struct may be populated later
			stat.AffectingNatures.Decrease[i] = &Nature{
				Name: nature.Name,
			}
		}
	}
	if aStat.Characteristics != nil {
		stat.Characteristics = make([]struct {
			URL string "json:\"url\""
		}, len(aStat.Characteristics))
		copy(stat.Characteristics, aStat.Characteristics)
	}
	if aStat.Names != nil {
		stat.Names = make([]struct {
			Name     string "json:\"name\""
			Language struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"language\""
		}, len(aStat.Names))
		copy(stat.Names, aStat.Names)
	}
	return stat
}
