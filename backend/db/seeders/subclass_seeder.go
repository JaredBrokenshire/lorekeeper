package seeders

import (
	"fmt"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
)

func (s *Seeder) SetSubclasses() {
	subclasses := []models.Subclass{
		{
			ID:      1,
			ClassID: 1,
			Name:    "Alchemist",
		},
		{
			ID:      2,
			ClassID: 1,
			Name:    "Armorer",
		},
		{
			ID:      3,
			ClassID: 1,
			Name:    "Artillerist",
		},
		{
			ID:      4,
			ClassID: 1,
			Name:    "Battle Smith",
		},
		{
			ID:      5,
			ClassID: 1,
			Name:    "Archivist",
		},
		{
			ID:      6,
			ClassID: 2,
			Name:    "Path of the Berserker",
		},
		{
			ID:      7,
			ClassID: 2,
			Name:    "Path of the Wild Heart",
		},
		{
			ID:      8,
			ClassID: 2,
			Name:    "Path of the World Tree",
		},
		{
			ID:      9,
			ClassID: 2,
			Name:    "Path of the Zealot",
		},
		{
			ID:      10,
			ClassID: 2,
			Name:    "Path of the Fractured",
		},
		{
			ID:      11,
			ClassID: 2,
			Name:    "Path of the Primal Spirit",
		},
		{
			ID:      12,
			ClassID: 2,
			Name:    "Path of the Wrathful Dead",
		},
		{
			ID:      13,
			ClassID: 3,
			Name:    "College of Dance",
		},
		{
			ID:      14,
			ClassID: 3,
			Name:    "College of Glamour",
		},
		{
			ID:      15,
			ClassID: 3,
			Name:    "College of Lore",
		},
		{
			ID:      16,
			ClassID: 3,
			Name:    "College of the Moon",
		},
		{
			ID:      17,
			ClassID: 3,
			Name:    "College of Valour",
		},
		{
			ID:      18,
			ClassID: 3,
			Name:    "College of Adventures",
		},
		{
			ID:      19,
			ClassID: 3,
			Name:    "College of Fools",
		},
		{
			ID:      20,
			ClassID: 3,
			Name:    "College of Requiems",
		},
		{
			ID:      21,
			ClassID: 4,
			Name:    "Knowledge Domain",
		},
		{
			ID:      22,
			ClassID: 4,
			Name:    "Life Domain",
		},
		{
			ID:      23,
			ClassID: 4,
			Name:    "Light Domain",
		},
		{
			ID:      24,
			ClassID: 4,
			Name:    "Trickery Domain",
		},
		{
			ID:      25,
			ClassID: 4,
			Name:    "War Domain",
		},
		{
			ID:      26,
			ClassID: 4,
			Name:    "Eldritch Domain",
		},
		{
			ID:      27,
			ClassID: 4,
			Name:    "Inquisition Domain",
		},
		{
			ID:      28,
			ClassID: 4,
			Name:    "Purification Domain",
		},
		{
			ID:      29,
			ClassID: 5,
			Name:    "Circle of the Land",
		},
		{
			ID:      30,
			ClassID: 5,
			Name:    "Circle of the Moon",
		},
		{
			ID:      31,
			ClassID: 5,
			Name:    "Circle of the Sea",
		},
		{
			ID:      32,
			ClassID: 5,
			Name:    "Circle of the Stars",
		},
		{
			ID:      33,
			ClassID: 5,
			Name:    "Circle of Blood",
		},
		{
			ID:      34,
			ClassID: 5,
			Name:    "Circle of Entropy",
		},
		{
			ID:      35,
			ClassID: 5,
			Name:    "Circle of Mutation",
		},
		{
			ID:      36,
			ClassID: 6,
			Name:    "Banneret",
		},
		{
			ID:      37,
			ClassID: 6,
			Name:    "Battle Master",
		},
		{
			ID:      38,
			ClassID: 6,
			Name:    "Champion",
		},
		{
			ID:      39,
			ClassID: 6,
			Name:    "Eldritch Knight",
		},
		{
			ID:      40,
			ClassID: 6,
			Name:    "Psi Warrior",
		},
		{
			ID:      41,
			ClassID: 6,
			Name:    "Sorrow Knight",
		},
		{
			ID:      42,
			ClassID: 6,
			Name:    "Bulwark Warrior",
		},
		{
			ID:      43,
			ClassID: 6,
			Name:    "Living Crucible",
		},
		{
			ID:      44,
			ClassID: 6,
			Name:    "Nightwatcher",
		},
		{
			ID:      45,
			ClassID: 7,
			Name:    "Warrior of Mercy",
		},
		{
			ID:      46,
			ClassID: 7,
			Name:    "Warrior of Shadow",
		},
		{
			ID:      47,
			ClassID: 7,
			Name:    "Warrior of the Elements",
		},
		{
			ID:      48,
			ClassID: 7,
			Name:    "Warrior of the Open Hand",
		},
		{
			ID:      49,
			ClassID: 7,
			Name:    "Warrior of Pride",
		},
		{
			ID:      50,
			ClassID: 7,
			Name:    "Warrior of Regret",
		},
		{
			ID:      51,
			ClassID: 7,
			Name:    "Warrior of the Leaden Crown",
		},
		{
			ID:      52,
			ClassID: 8,
			Name:    "Oath of Devotion",
		},
		{
			ID:      53,
			ClassID: 8,
			Name:    "Oath of Glory",
		},
		{
			ID:      54,
			ClassID: 8,
			Name:    "Oath of the Ancients",
		},
		{
			ID:      55,
			ClassID: 8,
			Name:    "Oath of the Noble Genies",
		},
		{
			ID:      56,
			ClassID: 8,
			Name:    "Oath of Vengeance",
		},
		{
			ID:      57,
			ClassID: 8,
			Name:    "Oath of Pestilence",
		},
		{
			ID:      58,
			ClassID: 8,
			Name:    "Oath of Slaughter",
		},
		{
			ID:      59,
			ClassID: 8,
			Name:    "Oath of Zeal",
		},
		{
			ID:      60,
			ClassID: 9,
			Name:    "Beast Master",
		},
		{
			ID:      61,
			ClassID: 9,
			Name:    "Fey Wanderer",
		},
		{
			ID:      62,
			ClassID: 9,
			Name:    "Gloom Stalker",
		},
		{
			ID:      63,
			ClassID: 9,
			Name:    "Hunter",
		},
		{
			ID:      64,
			ClassID: 9,
			Name:    "Winter Walker",
		},
		{
			ID:      65,
			ClassID: 9,
			Name:    "Green Reaper",
		},
		{
			ID:      66,
			ClassID: 9,
			Name:    "Primordial Archer",
		},
		{
			ID:      67,
			ClassID: 9,
			Name:    "Vermin Lord",
		},
		{
			ID:      68,
			ClassID: 10,
			Name:    "Arcane Trickster",
		},
		{
			ID:      69,
			ClassID: 10,
			Name:    "Assassin",
		},
		{
			ID:      70,
			ClassID: 10,
			Name:    "Scion of the Three",
		},
		{
			ID:      71,
			ClassID: 10,
			Name:    "Soulknife",
		},
		{
			ID:      72,
			ClassID: 10,
			Name:    "Thief",
		},
		{
			ID:      73,
			ClassID: 10,
			Name:    "Highway Rider",
		},
		{
			ID:      74,
			ClassID: 10,
			Name:    "Misfortune Bringer",
		},
		{
			ID:      75,
			ClassID: 10,
			Name:    "Sanguine Thief",
		},
		{
			ID:      76,
			ClassID: 11,
			Name:    "Aberrant Sorcery",
		},
		{
			ID:      77,
			ClassID: 11,
			Name:    "Clockwork Sorcery",
		},
		{
			ID:      78,
			ClassID: 11,
			Name:    "Draconic Sorcery",
		},
		{
			ID:      79,
			ClassID: 11,
			Name:    "Spellfire Sorcery",
		},
		{
			ID:      80,
			ClassID: 11,
			Name:    "Wild Magic Sorcery",
		},
		{
			ID:      81,
			ClassID: 11,
			Name:    "Apocalypse Sorcery",
		},
		{
			ID:      82,
			ClassID: 11,
			Name:    "Haunted Sorcery",
		},
		{
			ID:      83,
			ClassID: 12,
			Name:    "Wretched Bloodline",
		},
		{
			ID:      84,
			ClassID: 13,
			Name:    "Archfey Patron",
		},
		{
			ID:      85,
			ClassID: 13,
			Name:    "Celestial Patron",
		},
		{
			ID:      86,
			ClassID: 13,
			Name:    "Fiend Patron",
		},
		{
			ID:      87,
			ClassID: 13,
			Name:    "Great Old One Patron",
		},
		{
			ID:      88,
			ClassID: 13,
			Name:    "Parasite Patron",
		},
		{
			ID:      89,
			ClassID: 13,
			Name:    "The Coven",
		},
		{
			ID:      90,
			ClassID: 13,
			Name:    "The First Vampire Patron",
		},
		{
			ID:      91,
			ClassID: 14,
			Name:    "Abjurer",
		},
		{
			ID:      92,
			ClassID: 14,
			Name:    "Bladesinger",
		},
		{
			ID:      93,
			ClassID: 14,
			Name:    "Diviner",
		},
		{
			ID:      94,
			ClassID: 14,
			Name:    "Evoker",
		},
		{
			ID:      95,
			ClassID: 14,
			Name:    "Illusionist",
		},
		{
			ID:      96,
			ClassID: 14,
			Name:    "Daemonologist",
		},
		{
			ID:      97,
			ClassID: 14,
			Name:    "Plague Doctor",
		},
		{
			ID:      98,
			ClassID: 14,
			Name:    "Sangromancer",
		},
		{
			ID:      99,
			ClassID: 15,
			Name:    "Order of the Mutant",
		},
		{
			ID:      100,
			ClassID: 15,
			Name:    "Order of the Lycan",
		},
		{
			ID:      101,
			ClassID: 15,
			Name:    "Order of the Ghost Slayer",
		},
		{
			ID:      102,
			ClassID: 15,
			Name:    "Order of the Profane Soul",
		},
		{
			ID:      103,
			ClassID: 6,
			Name:    "Echo Knight",
		},
		{
			ID:      104,
			ClassID: 7,
			Name:    "Way of the Drunken Master",
		},
		{
			ID:      105,
			ClassID: 10,
			Name:    "Shadow Magic",
		},
		{
			ID:      106,
			ClassID: 4,
			Name:    "Twilight Domain",
		},
	}

	for _, subclass := range subclasses {
		err := s.DB.Where("id = ?", subclass.ID).FirstOrCreate(&subclass).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating subclass with id `%v` in seeder: %v", subclass.ID, err.Error()))
		}
	}
}
