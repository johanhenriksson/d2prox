package d2prox

type NPCType struct {
	ID      string
	ClassID int
	Name    string
}

// NPCTypes contains id/name data for all npcs
var NPCTypes = []*NPCType{
	&NPCType{"skeleton1", 0, "Skeleton"},
	&NPCType{"skeleton2", 1, "Returned"},
	&NPCType{"skeleton3", 2, "BoneWarrior"},
	&NPCType{"skeleton4", 3, "BurningDead"},
	&NPCType{"skeleton5", 4, "Horror"},
	&NPCType{"zombie1", 5, "Zombie"},
	&NPCType{"zombie2", 6, "HungryDead"},
	&NPCType{"zombie3", 7, "Ghoul"},
	&NPCType{"zombie4", 8, "DrownedCarcass"},
	&NPCType{"zombie5", 9, "PlagueBearer"},
	&NPCType{"bighead1", 10, "Afflicted"},
	&NPCType{"bighead2", 11, "Tainted"},
	&NPCType{"bighead3", 12, "Misshapen"},
	&NPCType{"bighead4", 13, "Disfigured"},
	&NPCType{"bighead5", 14, "Damned"},
	&NPCType{"foulcrow1", 15, "FoulCrow"},
	&NPCType{"foulcrow2", 16, "BloodHawk"},
	&NPCType{"foulcrow3", 17, "BlackRaptor"},
	&NPCType{"foulcrow4", 18, "CloudStalker"},
	&NPCType{"fallen1", 19, "Fallen"},
	&NPCType{"fallen2", 20, "Carver"},
	&NPCType{"fallen3", 21, "Devilkin"},
	&NPCType{"fallen4", 22, "DarkOne"},
	&NPCType{"fallen5", 23, "WarpedFallen"},
	&NPCType{"brute2", 24, "Brute"},
	&NPCType{"brute3", 25, "Yeti"},
	&NPCType{"brute4", 26, "Crusher"},
	&NPCType{"brute5", 27, "WailingBeast"},
	&NPCType{"brute1", 28, "GargantuanBeast"},
	&NPCType{"sandraider1", 29, "SandRaider"},
	&NPCType{"sandraider2", 30, "Marauder"},
	&NPCType{"sandraider3", 31, "Invader"},
	&NPCType{"sandraider4", 32, "Infidel"},
	&NPCType{"sandraider5", 33, "Assailant"},
	&NPCType{"gorgon1", 34, "unused"},
	&NPCType{"gorgon2", 35, "unused"},
	&NPCType{"gorgon3", 36, "unused"},
	&NPCType{"gorgon4", 37, "unused"},
	&NPCType{"wraith1", 38, "Ghost"},
	&NPCType{"wraith2", 39, "Wraith"},
	&NPCType{"wraith3", 40, "Specter"},
	&NPCType{"wraith4", 41, "Apparition"},
	&NPCType{"wraith5", 42, "DarkShape"},
	&NPCType{"corruptrogue1", 43, "DarkHunter"},
	&NPCType{"corruptrogue2", 44, "VileHunter"},
	&NPCType{"corruptrogue3", 45, "DarkStalker"},
	&NPCType{"corruptrogue4", 46, "BlackRogue"},
	&NPCType{"corruptrogue5", 47, "FleshHunter"},
	&NPCType{"baboon1", 48, "DuneBeast"},
	&NPCType{"baboon2", 49, "RockDweller"},
	&NPCType{"baboon3", 50, "JungleHunter"},
	&NPCType{"baboon4", 51, "DoomApe"},
	&NPCType{"baboon5", 52, "TempleGuard"},
	&NPCType{"goatman1", 53, "MoonClan"},
	&NPCType{"goatman2", 54, "NightClan"},
	&NPCType{"goatman3", 55, "BloodClan"},
	&NPCType{"goatman4", 56, "HellClan"},
	&NPCType{"goatman5", 57, "DeathClan"},
	&NPCType{"fallenshaman1", 58, "FallenShaman"},
	&NPCType{"fallenshaman2", 59, "CarverShaman"},
	&NPCType{"fallenshaman3", 60, "DevilkinShaman"},
	&NPCType{"fallenshaman4", 61, "DarkShaman"},
	&NPCType{"fallenshaman5", 62, "WarpedShaman"},
	&NPCType{"quillrat1", 63, "QuillRat"},
	&NPCType{"quillrat2", 64, "SpikeFiend"},
	&NPCType{"quillrat3", 65, "ThornBeast"},
	&NPCType{"quillrat4", 66, "RazorSpine"},
	&NPCType{"quillrat5", 67, "JungleUrchin"},
	&NPCType{"sandmaggot1", 68, "SandMaggot"},
	&NPCType{"sandmaggot2", 69, "RockWorm"},
	&NPCType{"sandmaggot3", 70, "Devourer"},
	&NPCType{"sandmaggot4", 71, "GiantLamprey"},
	&NPCType{"sandmaggot5", 72, "WorldKiller"},
	&NPCType{"clawviper1", 73, "TombViper"},
	&NPCType{"clawviper2", 74, "ClawViper"},
	&NPCType{"clawviper3", 75, "Salamander"},
	&NPCType{"clawviper4", 76, "PitViper"},
	&NPCType{"clawviper5", 77, "SerpentMagus"},
	&NPCType{"sandleaper1", 78, "SandLeaper"},
	&NPCType{"sandleaper2", 79, "CaveLeaper"},
	&NPCType{"sandleaper3", 80, "TombCreeper"},
	&NPCType{"sandleaper4", 81, "TreeLurker"},
	&NPCType{"sandleaper5", 82, "RazorPitDemon"},
	&NPCType{"pantherwoman1", 83, "Huntress"},
	&NPCType{"pantherwoman2", 84, "SaberCat"},
	&NPCType{"pantherwoman3", 85, "NightTiger"},
	&NPCType{"pantherwoman4", 86, "HellCat"},
	&NPCType{"swarm1", 87, "Itchies"},
	&NPCType{"swarm2", 88, "BlackLocusts"},
	&NPCType{"swarm3", 89, "PlagueBugs"},
	&NPCType{"swarm4", 90, "HellSwarm"},
	&NPCType{"scarab1", 91, "DungSoldier"},
	&NPCType{"scarab2", 92, "SandWarrior"},
	&NPCType{"scarab3", 93, "Scarab"},
	&NPCType{"scarab4", 94, "SteelWeevil"},
	&NPCType{"scarab5", 95, "AlbinoRoach"},
	&NPCType{"mummy1", 96, "DriedCorpse"},
	&NPCType{"mummy2", 97, "Decayed"},
	&NPCType{"mummy3", 98, "Embalmed"},
	&NPCType{"mummy4", 99, "PreservedDead"},
	&NPCType{"mummy5", 100, "Cadaver"},
	&NPCType{"unraveler1", 101, "HollowOne"},
	&NPCType{"unraveler2", 102, "Guardian"},
	&NPCType{"unraveler3", 103, "Unraveler"},
	&NPCType{"unraveler4", 104, "Horadrim Ancient"},
	&NPCType{"unraveler5", 105, "Baal Subject Mummy"},
	&NPCType{"chaoshorde1", 106, "unused"},
	&NPCType{"chaoshorde2", 107, "unused"},
	&NPCType{"chaoshorde3", 108, "unused"},
	&NPCType{"chaoshorde4", 109, "unused"},
	&NPCType{"vulture1", 110, "CarrionBird"},
	&NPCType{"vulture2", 111, "UndeadScavenger"},
	&NPCType{"vulture3", 112, "HellBuzzard"},
	&NPCType{"vulture4", 113, "WingedNightmare"},
	&NPCType{"mosquito1", 114, "Sucker"},
	&NPCType{"mosquito2", 115, "Feeder"},
	&NPCType{"mosquito3", 116, "BloodHook"},
	&NPCType{"mosquito4", 117, "BloodWing"},
	&NPCType{"willowisp1", 118, "Gloam"},
	&NPCType{"willowisp2", 119, "SwampGhost"},
	&NPCType{"willowisp3", 120, "BurningSoul"},
	&NPCType{"willowisp4", 121, "BlackSoul"},
	&NPCType{"arach1", 122, "Arach"},
	&NPCType{"arach2", 123, "SandFisher"},
	&NPCType{"arach3", 124, "PoisonSpinner"},
	&NPCType{"arach4", 125, "FlameSpider"},
	&NPCType{"arach5", 126, "SpiderMagus"},
	&NPCType{"thornhulk1", 127, "ThornedHulk"},
	&NPCType{"thornhulk2", 128, "BrambleHulk"},
	&NPCType{"thornhulk3", 129, "Thrasher"},
	&NPCType{"thornhulk4", 130, "Spikefist"},
	&NPCType{"vampire1", 131, "GhoulLord"},
	&NPCType{"vampire2", 132, "NightLord"},
	&NPCType{"vampire3", 133, "DarkLord"},
	&NPCType{"vampire4", 134, "BloodLord"},
	&NPCType{"vampire5", 135, "Banished"},
	&NPCType{"batdemon1", 136, "DesertWing"},
	&NPCType{"batdemon2", 137, "Fiend"},
	&NPCType{"batdemon3", 138, "Gloombat"},
	&NPCType{"batdemon4", 139, "BloodDiver"},
	&NPCType{"batdemon5", 140, "DarkFamiliar"},
	&NPCType{"fetish1", 141, "RatMan"},
	&NPCType{"fetish2", 142, "Fetish"},
	&NPCType{"fetish3", 143, "Flayer"},
	&NPCType{"fetish4", 144, "SoulKiller"},
	&NPCType{"fetish5", 145, "StygianDoll"},
	&NPCType{"cain1", 146, "DeckardCain"},
	&NPCType{"gheed", 147, "Gheed"},
	&NPCType{"akara", 148, "Akara"},
	&NPCType{"chicken", 149, "dummy"},
	&NPCType{"kashya", 150, "Kashya"},
	&NPCType{"rat", 151, "dummy"},
	&NPCType{"rogue1", 152, "dummy"},
	&NPCType{"hellmeteor", 153, "dummy"},
	&NPCType{"charsi", 154, "Charsi"},
	&NPCType{"warriv1", 155, "Warriv"},
	&NPCType{"andariel", 156, "Andariel"},
	&NPCType{"bird1", 157, "dummy"},
	&NPCType{"bird2", 158, "dummy"},
	&NPCType{"bat", 159, "dummy"},
	&NPCType{"cr_archer1", 160, "DarkRanger"},
	&NPCType{"cr_archer2", 161, "VileArcher"},
	&NPCType{"cr_archer3", 162, "DarkArcher"},
	&NPCType{"cr_archer4", 163, "BlackArcher"},
	&NPCType{"cr_archer5", 164, "FleshArcher"},
	&NPCType{"cr_lancer1", 165, "DarkSpearwoman"},
	&NPCType{"cr_lancer2", 166, "VileLancer"},
	&NPCType{"cr_lancer3", 167, "DarkLancer"},
	&NPCType{"cr_lancer4", 168, "BlackLancer"},
	&NPCType{"cr_lancer5", 169, "FleshLancer"},
	&NPCType{"sk_archer1", 170, "SkeletonArcher"},
	&NPCType{"sk_archer2", 171, "ReturnedArcher"},
	&NPCType{"sk_archer3", 172, "BoneArcher"},
	&NPCType{"sk_archer4", 173, "BurningDeadArcher"},
	&NPCType{"sk_archer5", 174, "HorrorArcher"},
	&NPCType{"warriv2", 175, "Warriv"},
	&NPCType{"atma", 176, "Atma"},
	&NPCType{"drognan", 177, "Drognan"},
	&NPCType{"fara", 178, "Fara"},
	&NPCType{"cow", 179, "dummy"},
	&NPCType{"maggotbaby1", 180, "SandMaggotYoung"},
	&NPCType{"maggotbaby2", 181, "RockWormYoung"},
	&NPCType{"maggotbaby3", 182, "DevourerYoung"},
	&NPCType{"maggotbaby4", 183, "GiantLampreyYoung"},
	&NPCType{"maggotbaby5", 184, "WorldKillerYoung"},
	&NPCType{"camel", 185, "dummy"},
	&NPCType{"blunderbore1", 186, "Blunderbore"},
	&NPCType{"blunderbore2", 187, "Gorbelly"},
	&NPCType{"blunderbore3", 188, "Mauler"},
	&NPCType{"blunderbore4", 189, "Urdar"},
	&NPCType{"maggotegg1", 190, "SandMaggotEgg"},
	&NPCType{"maggotegg2", 191, "RockWormEgg"},
	&NPCType{"maggotegg3", 192, "DevourerEgg"},
	&NPCType{"maggotegg4", 193, "GiantLampreyEgg"},
	&NPCType{"maggotegg5", 194, "WorldKillerEgg"},
	&NPCType{"act2male", 195, "dummy"},
	&NPCType{"act2female", 196, "dummy"},
	&NPCType{"act2child", 197, "dummy"},
	&NPCType{"greiz", 198, "Greiz"},
	&NPCType{"elzix", 199, "Elzix"},
	&NPCType{"geglash", 200, "Geglash"},
	&NPCType{"jerhyn", 201, "Jerhyn"},
	&NPCType{"lysander", 202, "Lysander"},
	&NPCType{"act2guard1", 203, "dummy"},
	&NPCType{"act2vendor1", 204, "dummy"},
	&NPCType{"act2vendor2", 205, "dummy"},
	&NPCType{"crownest1", 206, "FoulCrowNest"},
	&NPCType{"crownest2", 207, "BloodHawkNest"},
	&NPCType{"crownest3", 208, "BlackVultureNest"},
	&NPCType{"crownest4", 209, "CloudStalkerNest"},
	&NPCType{"meshif1", 210, "Meshif"},
	&NPCType{"duriel", 211, "Duriel"},
	&NPCType{"bonefetish1", 212, "Undead RatMan"},
	&NPCType{"bonefetish2", 213, "Undead Fetish"},
	&NPCType{"bonefetish3", 214, "Undead Flayer"},
	&NPCType{"bonefetish4", 215, "Undead SoulKiller"},
	&NPCType{"bonefetish5", 216, "Undead StygianDoll"},
	&NPCType{"darkguard1", 217, "unused"},
	&NPCType{"darkguard2", 218, "unused"},
	&NPCType{"darkguard3", 219, "unused"},
	&NPCType{"darkguard4", 220, "unused"},
	&NPCType{"darkguard5", 221, "unused"},
	&NPCType{"bloodmage1", 222, "unused"},
	&NPCType{"bloodmage2", 223, "unused"},
	&NPCType{"bloodmage3", 224, "unused"},
	&NPCType{"bloodmage4", 225, "unused"},
	&NPCType{"bloodmage5", 226, "unused"},
	&NPCType{"maggot", 227, "Maggot"},
	&NPCType{"sarcophagus", 228, "MummyGenerator"},
	&NPCType{"radament", 229, "Radament"},
	&NPCType{"firebeast", 230, "unused"},
	&NPCType{"iceglobe", 231, "unused"},
	&NPCType{"lightningbeast", 232, "unused"},
	&NPCType{"poisonorb", 233, "unused"},
	&NPCType{"flyingscimitar", 234, "FlyingScimitar"},
	&NPCType{"zealot1", 235, "Zakarumite"},
	&NPCType{"zealot2", 236, "Faithful"},
	&NPCType{"zealot3", 237, "Zealot"},
	&NPCType{"cantor1", 238, "Sexton"},
	&NPCType{"cantor2", 239, "Cantor"},
	&NPCType{"cantor3", 240, "Heirophant"},
	&NPCType{"cantor4", 241, "Heirophant"},
	&NPCType{"mephisto", 242, "Mephisto"},
	&NPCType{"diablo", 243, "Diablo"},
	&NPCType{"cain2", 244, "DeckardCain"},
	&NPCType{"cain3", 245, "DeckardCain"},
	&NPCType{"cain4", 246, "DeckardCain"},
	&NPCType{"frogdemon1", 247, "Swamp Dweller"},
	&NPCType{"frogdemon2", 248, "Bog Creature"},
	&NPCType{"frogdemon3", 249, "Slime Prince"},
	&NPCType{"summoner", 250, "Summoner"},
	&NPCType{"tyrael1", 251, "tyrael"},
	&NPCType{"asheara", 252, "asheara"},
	&NPCType{"hratli", 253, "hratli"},
	&NPCType{"alkor", 254, "alkor"},
	&NPCType{"ormus", 255, "ormus"},
	&NPCType{"izual", 256, "izual"},
	&NPCType{"halbu", 257, "halbu"},
	&NPCType{"tentacle1", 258, "WaterWatcherLimb"},
	&NPCType{"tentacle2", 259, "RiverStalkerLimb"},
	&NPCType{"tentacle3", 260, "StygianWatcherLimb"},
	&NPCType{"tentaclehead1", 261, "WaterWatcherHead"},
	&NPCType{"tentaclehead2", 262, "RiverStalkerHead"},
	&NPCType{"tentaclehead3", 263, "StygianWatcherHead"},
	&NPCType{"meshif2", 264, "meshif"},
	&NPCType{"cain5", 265, "DeckardCain"},
	&NPCType{"navi", 266, "navi"},
	&NPCType{"bloodraven", 267, "Bloodraven"},
	&NPCType{"bug", 268, "dummy"},
	&NPCType{"scorpion", 269, "dummy"},
	&NPCType{"rogue2", 270, "RogueScout"},
	&NPCType{"roguehire", 271, "dummy"},
	&NPCType{"rogue3", 272, "dummy"},
	&NPCType{"gargoyletrap", 273, "GargoyleTrap"},
	&NPCType{"skmage_pois1", 274, "ReturnedMage"},
	&NPCType{"skmage_pois2", 275, "BoneMage"},
	&NPCType{"skmage_pois3", 276, "BurningDeadMage"},
	&NPCType{"skmage_pois4", 277, "HorrorMage"},
	&NPCType{"fetishshaman1", 278, "RatManShaman"},
	&NPCType{"fetishshaman2", 279, "FetishShaman"},
	&NPCType{"fetishshaman3", 280, "FlayerShaman"},
	&NPCType{"fetishshaman4", 281, "SoulKillerShaman"},
	&NPCType{"fetishshaman5", 282, "StygianDollShaman"},
	&NPCType{"larva", 283, "larva"},
	&NPCType{"maggotqueen1", 284, "SandMaggotQueen"},
	&NPCType{"maggotqueen2", 285, "RockWormQueen"},
	&NPCType{"maggotqueen3", 286, "DevourerQueen"},
	&NPCType{"maggotqueen4", 287, "GiantLampreyQueen"},
	&NPCType{"maggotqueen5", 288, "WorldKillerQueen"},
	&NPCType{"claygolem", 289, "ClayGolem"},
	&NPCType{"bloodgolem", 290, "BloodGolem"},
	&NPCType{"irongolem", 291, "IronGolem"},
	&NPCType{"firegolem", 292, "FireGolem"},
	&NPCType{"familiar", 293, "dummy"},
	&NPCType{"act3male", 294, "dummy"},
	&NPCType{"baboon6", 295, "NightMarauder"},
	&NPCType{"act3female", 296, "dummy"},
	&NPCType{"natalya", 297, "Natalya"},
	&NPCType{"vilemother1", 298, "FleshSpawner"},
	&NPCType{"vilemother2", 299, "StygianHag"},
	&NPCType{"vilemother3", 300, "Grotesque"},
	&NPCType{"vilechild1", 301, "FleshBeast"},
	&NPCType{"vilechild2", 302, "StygianDog"},
	&NPCType{"vilechild3", 303, "GrotesqueWyrm"},
	&NPCType{"fingermage1", 304, "Groper"},
	&NPCType{"fingermage2", 305, "Strangler"},
	&NPCType{"fingermage3", 306, "StormCaster"},
	&NPCType{"regurgitator1", 307, "Corpulent"},
	&NPCType{"regurgitator2", 308, "CorpseSpitter"},
	&NPCType{"regurgitator3", 309, "MawFiend"},
	&NPCType{"doomknight1", 310, "DoomKnight"},
	&NPCType{"doomknight2", 311, "AbyssKnight"},
	&NPCType{"doomknight3", 312, "OblivionKnight"},
	&NPCType{"quillbear1", 313, "QuillBear"},
	&NPCType{"quillbear2", 314, "SpikeGiant"},
	&NPCType{"quillbear3", 315, "ThornBrute"},
	&NPCType{"quillbear4", 316, "RazorBeast"},
	&NPCType{"quillbear5", 317, "GiantUrchin"},
	&NPCType{"snake", 318, "dummy"},
	&NPCType{"parrot", 319, "dummy"},
	&NPCType{"fish", 320, "dummy"},
	&NPCType{"evilhole1", 321, "dummy"},
	&NPCType{"evilhole2", 322, "dummy"},
	&NPCType{"evilhole3", 323, "dummy"},
	&NPCType{"evilhole4", 324, "dummy"},
	&NPCType{"evilhole5", 325, "dummy"},
	&NPCType{"trap-firebolt", 326, "FireTrap"},
	&NPCType{"trap-horzmissile", 327, "MissleTrap"},
	&NPCType{"trap-vertmissile", 328, "MissileTrap"},
	&NPCType{"trap-poisoncloud", 329, "PoisonCloudTrap"},
	&NPCType{"trap-lightning", 330, "LightningTrap"},
	&NPCType{"act2guard2", 331, "Kaelan"},
	&NPCType{"invisospawner", 332, "dummy"},
	&NPCType{"diabloclone", 333, "Diablo"},
	&NPCType{"suckernest1", 334, "SuckerNest"},
	&NPCType{"suckernest2", 335, "FeederNest"},
	&NPCType{"suckernest3", 336, "BloodHookNest"},
	&NPCType{"suckernest4", 337, "BloodWingNest"},
	&NPCType{"act2hire", 338, "Guard"},
	&NPCType{"minispider", 339, "dummy"},
	&NPCType{"boneprison1", 340, "BonePrison"},
	&NPCType{"boneprison2", 341, "BonePrison"},
	&NPCType{"boneprison3", 342, "BonePrison"},
	&NPCType{"boneprison4", 343, "BonePrison"},
	&NPCType{"bonewall", 344, "Bonewall"},
	&NPCType{"councilmember1", 345, "Council Member"},
	&NPCType{"councilmember2", 346, "Council Member"},
	&NPCType{"councilmember3", 347, "Council Member"},
	&NPCType{"turret1", 348, "Turret"},
	&NPCType{"turret2", 349, "Turret"},
	&NPCType{"turret3", 350, "Turret"},
	&NPCType{"hydra1", 351, "Hydra"},
	&NPCType{"hydra2", 352, "Hydra"},
	&NPCType{"hydra3", 353, "Hydra"},
	&NPCType{"trap-melee", 354, "MeleeTrap"},
	&NPCType{"seventombs", 355, "dummy"},
	&NPCType{"dopplezon", 356, "Dopplezon"},
	&NPCType{"valkyrie", 357, "Valkyrie"},
	&NPCType{"act2guard3", 358, "dummy"},
	&NPCType{"act3hire", 359, "Iron Wolf"},
	&NPCType{"megademon1", 360, "Balrog"},
	&NPCType{"megademon2", 361, "PitLord"},
	&NPCType{"megademon3", 362, "VenomLord"},
	&NPCType{"necroskeleton", 363, "NecroSkeleton"},
	&NPCType{"necromage", 364, "NecroMage"},
	&NPCType{"griswold", 365, "Griswold"},
	&NPCType{"compellingorb", 366, "compellingorb"},
	&NPCType{"tyrael2", 367, "tyrael"},
	&NPCType{"darkwanderer", 368, "youngdiablo"},
	&NPCType{"trap-nova", 369, "NovaTrap"},
	&NPCType{"spiritmummy", 370, "dummy"},
	&NPCType{"lightningspire", 371, "LightningSpire"},
	&NPCType{"firetower", 372, "FireTower"},
	&NPCType{"slinger1", 373, "Slinger"},
	&NPCType{"slinger2", 374, "SpearCat"},
	&NPCType{"slinger3", 375, "NightSlinger"},
	&NPCType{"slinger4", 376, "HellSlinger"},
	&NPCType{"act2guard4", 377, "dummy"},
	&NPCType{"act2guard5", 378, "dummy"},
	&NPCType{"skmage_cold1", 379, "ReturnedMage"},
	&NPCType{"skmage_cold2", 380, "BoneMage"},
	&NPCType{"skmage_cold3", 381, "BaalColdMage"},
	&NPCType{"skmage_cold4", 382, "HorrorMage"},
	&NPCType{"skmage_fire1", 383, "ReturnedMage"},
	&NPCType{"skmage_fire2", 384, "BoneMage"},
	&NPCType{"skmage_fire3", 385, "BurningDeadMage"},
	&NPCType{"skmage_fire4", 386, "HorrorMage"},
	&NPCType{"skmage_ltng1", 387, "ReturnedMage"},
	&NPCType{"skmage_ltng2", 388, "BoneMage"},
	&NPCType{"skmage_ltng3", 389, "BurningDeadMage"},
	&NPCType{"skmage_ltng4", 390, "HorrorMage"},
	&NPCType{"hellbovine", 391, "Hell Bovine"},
	&NPCType{"window1", 392, "Window"},
	&NPCType{"window2", 393, "Window"},
	&NPCType{"slinger5", 394, "SpearCat"},
	&NPCType{"slinger6", 395, "NightSlinger"},
	&NPCType{"fetishblow1", 396, "RatMan"},
	&NPCType{"fetishblow2", 397, "Fetish"},
	&NPCType{"fetishblow3", 398, "Flayer"},
	&NPCType{"fetishblow4", 399, "SoulKiller"},
	&NPCType{"fetishblow5", 400, "StygianDoll"},
	&NPCType{"mephistospirit", 401, "dummy"},
	&NPCType{"smith", 402, "The Smith"},
	&NPCType{"trappedsoul1", 403, "TrappedSoul"},
	&NPCType{"trappedsoul2", 404, "TrappedSoul"},
	&NPCType{"jamella", 405, "Jamella"},
	&NPCType{"izualghost", 406, "Izual"},
	&NPCType{"fetish11", 407, "RatMan"},
	&NPCType{"malachai", 408, "Malachai"},
	&NPCType{"hephasto", 409, "The Feature Creep"},
	&NPCType{"wakeofdestruction", 410, "Wake of Destruction"},
	&NPCType{"chargeboltsentry", 411, "Charged Bolt Sentry"},
	&NPCType{"lightningsentry", 412, "Lightning Sentry"},
	&NPCType{"bladecreeper", 413, "Blade Creeper"},
	&NPCType{"invisopet", 414, "Invis Pet"},
	&NPCType{"infernosentry", 415, "Inferno Sentry"},
	&NPCType{"deathsentry", 416, "Death Sentry"},
	&NPCType{"shadowwarrior", 417, "Shadow Warrior"},
	&NPCType{"shadowmaster", 418, "Shadow Master"},
	&NPCType{"druidhawk", 419, "Druid Hawk"},
	&NPCType{"spiritwolf", 420, "Druid Spirit Wolf"},
	&NPCType{"fenris", 421, "Druid Fenris"},
	&NPCType{"spiritofbarbs", 422, "Spirit of Barbs"},
	&NPCType{"heartofwolverine", 423, "Heart of Wolverine"},
	&NPCType{"oaksage", 424, "Oak Sage"},
	&NPCType{"plaguepoppy", 425, "Druid Plague Poppy"},
	&NPCType{"cycleoflife", 426, "Druid Cycle of Life"},
	&NPCType{"vinecreature", 427, "Vine Creature"},
	&NPCType{"druidbear", 428, "Druid Bear"},
	&NPCType{"eagle", 429, "Eagle"},
	&NPCType{"wolf", 430, "Wolf"},
	&NPCType{"bear", 431, "Bear"},
	&NPCType{"barricadedoor1", 432, "Barricade Door"},
	&NPCType{"barricadedoor2", 433, "Barricade Door"},
	&NPCType{"prisondoor", 434, "Prison Door"},
	&NPCType{"barricadetower", 435, "Barricade Tower"},
	&NPCType{"reanimatedhorde1", 436, "RotWalker"},
	&NPCType{"reanimatedhorde2", 437, "ReanimatedHorde"},
	&NPCType{"reanimatedhorde3", 438, "ProwlingDead"},
	&NPCType{"reanimatedhorde4", 439, "UnholyCorpse"},
	&NPCType{"reanimatedhorde5", 440, "DefiledWarrior"},
	&NPCType{"siegebeast1", 441, "Siege Beast"},
	&NPCType{"siegebeast2", 442, "CrushBiest"},
	&NPCType{"siegebeast3", 443, "BloodBringer"},
	&NPCType{"siegebeast4", 444, "GoreBearer"},
	&NPCType{"siegebeast5", 445, "DeamonSteed"},
	&NPCType{"snowyeti1", 446, "SnowYeti1"},
	&NPCType{"snowyeti2", 447, "SnowYeti2"},
	&NPCType{"snowyeti3", 448, "SnowYeti3"},
	&NPCType{"snowyeti4", 449, "SnowYeti4"},
	&NPCType{"wolfrider1", 450, "WolfRider1"},
	&NPCType{"wolfrider2", 451, "WolfRider2"},
	&NPCType{"wolfrider3", 452, "WolfRider3"},
	&NPCType{"minion1", 453, "Minionexp"},
	&NPCType{"minion2", 454, "Slayerexp"},
	&NPCType{"minion3", 455, "IceBoar"},
	&NPCType{"minion4", 456, "FireBoar"},
	&NPCType{"minion5", 457, "HellSpawn"},
	&NPCType{"minion6", 458, "IceSpawn"},
	&NPCType{"minion7", 459, "GreaterHellSpawn"},
	&NPCType{"minion8", 460, "GreaterIceSpawn"},
	&NPCType{"suicideminion1", 461, "FanaticMinion"},
	&NPCType{"suicideminion2", 462, "BerserkSlayer"},
	&NPCType{"suicideminion3", 463, "ConsumedIceBoar"},
	&NPCType{"suicideminion4", 464, "ConsumedFireBoar"},
	&NPCType{"suicideminion5", 465, "FrenziedHellSpawn"},
	&NPCType{"suicideminion6", 466, "FrenziedIceSpawn"},
	&NPCType{"suicideminion7", 467, "InsaneHellSpawn"},
	&NPCType{"suicideminion8", 468, "InsaneIceSpawn"},
	&NPCType{"succubus1", 469, "Succubusexp"},
	&NPCType{"succubus2", 470, "VileTemptress"},
	&NPCType{"succubus3", 471, "StygianHarlot"},
	&NPCType{"succubus4", 472, "Hell Temptress"},
	&NPCType{"succubus5", 473, "Blood Temptress"},
	&NPCType{"succubuswitch1", 474, "Dominus"},
	&NPCType{"succubuswitch2", 475, "VileWitch"},
	&NPCType{"succubuswitch3", 476, "StygianFury"},
	&NPCType{"succubuswitch4", 477, "Blood Witch"},
	&NPCType{"succubuswitch5", 478, "Hell Witch"},
	&NPCType{"overseer1", 479, "OverSeer"},
	&NPCType{"overseer2", 480, "Lasher"},
	&NPCType{"overseer3", 481, "OverLord"},
	&NPCType{"overseer4", 482, "BloodBoss"},
	&NPCType{"overseer5", 483, "HellWhip"},
	&NPCType{"minionspawner1", 484, "MinionSpawner"},
	&NPCType{"minionspawner2", 485, "MinionSlayerSpawner"},
	&NPCType{"minionspawner3", 486, "MinionIce/fireBoarSpawner"},
	&NPCType{"minionspawner4", 487, "MinionIce/fireBoarSpawner"},
	&NPCType{"minionspawner5", 488, "Minionice/hellSpawnSpawner"},
	&NPCType{"minionspawner6", 489, "MinionIce/fireBoarSpawner"},
	&NPCType{"minionspawner7", 490, "MinionIce/fireBoarSpawner"},
	&NPCType{"minionspawner8", 491, "Minionice/hellSpawnSpawner"},
	&NPCType{"imp1", 492, "Imp1"},
	&NPCType{"imp2", 493, "Imp2"},
	&NPCType{"imp3", 494, "Imp3"},
	&NPCType{"imp4", 495, "Imp4"},
	&NPCType{"imp5", 496, "Imp5"},
	&NPCType{"catapult1", 497, "CatapultS"},
	&NPCType{"catapult2", 498, "CatapultE"},
	&NPCType{"catapult3", 499, "CatapultSiege"},
	&NPCType{"catapult4", 500, "CatapultW"},
	&NPCType{"frozenhorror1", 501, "Frozen Horror1"},
	&NPCType{"frozenhorror2", 502, "Frozen Horror2"},
	&NPCType{"frozenhorror3", 503, "Frozen Horror3"},
	&NPCType{"frozenhorror4", 504, "Frozen Horror4"},
	&NPCType{"frozenhorror5", 505, "Frozen Horror5"},
	&NPCType{"bloodlord1", 506, "Blood Lord1"},
	&NPCType{"bloodlord2", 507, "Blood Lord2"},
	&NPCType{"bloodlord3", 508, "Blood Lord3"},
	&NPCType{"bloodlord4", 509, "Blood Lord4"},
	&NPCType{"bloodlord5", 510, "Blood Lord5"},
	&NPCType{"larzuk", 511, "Larzuk"},
	&NPCType{"drehya", 512, "Drehya"},
	&NPCType{"malah", 513, "Malah"},
	&NPCType{"nihlathak", 514, "Nihlathak Town"},
	&NPCType{"qual-kehk", 515, "Qual-Kehk"},
	&NPCType{"catapultspotter1", 516, "Catapult Spotter S"},
	&NPCType{"catapultspotter2", 517, "Catapult Spotter E"},
	&NPCType{"catapultspotter3", 518, "Catapult Spotter Siege"},
	&NPCType{"catapultspotter4", 519, "Catapult Spotter W"},
	&NPCType{"cain6", 520, "DeckardCain"},
	&NPCType{"tyrael3", 521, "tyrael"},
	&NPCType{"act5barb1", 522, "Act 5 Combatant"},
	&NPCType{"act5barb2", 523, "Act 5 Combatant"},
	&NPCType{"barricadewall1", 524, "Barricade Wall Right"},
	&NPCType{"barricadewall2", 525, "Barricade Wall Left"},
	&NPCType{"nihlathakboss", 526, "Nihlathak"},
	&NPCType{"drehyaiced", 527, "Drehya"},
	&NPCType{"evilhut", 528, "Evil hut"},
	&NPCType{"deathmauler1", 529, "Death Mauler1"},
	&NPCType{"deathmauler2", 530, "Death Mauler2"},
	&NPCType{"deathmauler3", 531, "Death Mauler3"},
	&NPCType{"deathmauler4", 532, "Death Mauler4"},
	&NPCType{"deathmauler5", 533, "Death Mauler5"},
	&NPCType{"act5pow", 534, "POW"},
	&NPCType{"act5barb3", 535, "Act 5 Townguard"},
	&NPCType{"act5barb4", 536, "Act 5 Townguard"},
	&NPCType{"ancientstatue1", 537, "Ancient Statue 1"},
	&NPCType{"ancientstatue2", 538, "Ancient Statue 2"},
	&NPCType{"ancientstatue3", 539, "Ancient Statue 3"},
	&NPCType{"ancientbarb1", 540, "Ancient Barbarian 1"},
	&NPCType{"ancientbarb2", 541, "Ancient Barbarian 2"},
	&NPCType{"ancientbarb3", 542, "Ancient Barbarian 3"},
	&NPCType{"baalthrone", 543, "Baal Throne"},
	&NPCType{"baalcrab", 544, "Baal Crab"},
	&NPCType{"baaltaunt", 545, "Baal Taunt"},
	&NPCType{"putriddefiler1", 546, "Putrid Defiler1"},
	&NPCType{"putriddefiler2", 547, "Putrid Defiler2"},
	&NPCType{"putriddefiler3", 548, "Putrid Defiler3"},
	&NPCType{"putriddefiler4", 549, "Putrid Defiler4"},
	&NPCType{"putriddefiler5", 550, "Putrid Defiler5"},
	&NPCType{"painworm1", 551, "Pain Worm1"},
	&NPCType{"painworm2", 552, "Pain Worm2"},
	&NPCType{"painworm3", 553, "Pain Worm3"},
	&NPCType{"painworm4", 554, "Pain Worm4"},
	&NPCType{"painworm5", 555, "Pain Worm5"},
	&NPCType{"bunny", 556, "Bunny"},
	&NPCType{"baalhighpriest", 557, "Council Member"},
	&NPCType{"venomlord", 558, "VenomLord"},
	&NPCType{"baalcrabstairs", 559, "Baal Crab to Stairs"},
	&NPCType{"act5hire1", 560, "Act 5 Hireling 1hs"},
	&NPCType{"act5hire2", 561, "Act 5 Hireling 2hs"},
	&NPCType{"baaltentacle1", 562, "Baal Tentacle"},
	&NPCType{"baaltentacle2", 563, "Baal Tentacle"},
	&NPCType{"baaltentacle3", 564, "Baal Tentacle"},
	&NPCType{"baaltentacle4", 565, "Baal Tentacle"},
	&NPCType{"baaltentacle5", 566, "Baal Tentacle"},
	&NPCType{"injuredbarb1", 567, "Injured Barbarian 1"},
	&NPCType{"injuredbarb2", 568, "Injured Barbarian 2"},
	&NPCType{"injuredbarb3", 569, "Injured Barbarian 3"},
	&NPCType{"baalclone", 570, "Baal Crab Clone"},
	&NPCType{"baalminion1", 571, "Baals Minion"},
	&NPCType{"baalminion2", 572, "Baals Minion"},
	&NPCType{"baalminion3", 573, "Baals Minion"},
	&NPCType{"worldstoneeffect", 574, "Worldstone Effect"},
	&NPCType{"sk_archer6", 575, "BurningDeadArcher"},
	&NPCType{"sk_archer7", 576, "BoneArcher"},
	&NPCType{"sk_archer8", 577, "BurningDeadArcher"},
	&NPCType{"sk_archer9", 578, "ReturnedArcher"},
	&NPCType{"sk_archer10", 579, "HorrorArcher"},
	&NPCType{"bighead6", 580, "Afflicted"},
	&NPCType{"bighead7", 581, "Tainted"},
	&NPCType{"bighead8", 582, "Misshapen"},
	&NPCType{"bighead9", 583, "Disfigured"},
	&NPCType{"bighead10", 584, "Damned"},
	&NPCType{"goatman6", 585, "MoonClan"},
	&NPCType{"goatman7", 586, "NightClan"},
	&NPCType{"goatman8", 587, "HellClan"},
	&NPCType{"goatman9", 588, "BloodClan"},
	&NPCType{"goatman10", 589, "DeathClan"},
	&NPCType{"foulcrow5", 590, "FoulCrow"},
	&NPCType{"foulcrow6", 591, "BloodHawk"},
	&NPCType{"foulcrow7", 592, "BlackRaptor"},
	&NPCType{"foulcrow8", 593, "CloudStalker"},
	&NPCType{"clawviper6", 594, "ClawViper"},
	&NPCType{"clawviper7", 595, "PitViper"},
	&NPCType{"clawviper8", 596, "Salamander"},
	&NPCType{"clawviper9", 597, "TombViper"},
	&NPCType{"clawviper10", 598, "SerpentMagus"},
	&NPCType{"sandraider6", 599, "Marauder"},
	&NPCType{"sandraider7", 600, "Infidel"},
	&NPCType{"sandraider8", 601, "SandRaider"},
	&NPCType{"sandraider9", 602, "Invader"},
	&NPCType{"sandraider10", 603, "Assailant"},
	&NPCType{"deathmauler6", 604, "Death Mauler1"},
	&NPCType{"quillrat6", 605, "QuillRat"},
	&NPCType{"quillrat7", 606, "SpikeFiend"},
	&NPCType{"quillrat8", 607, "RazorSpine"},
	&NPCType{"vulture5", 608, "CarrionBird"},
	&NPCType{"thornhulk5", 609, "ThornedHulk"},
	&NPCType{"slinger7", 610, "Slinger"},
	&NPCType{"slinger8", 611, "Slinger"},
	&NPCType{"slinger9", 612, "Slinger"},
	&NPCType{"cr_archer6", 613, "VileArcher"},
	&NPCType{"cr_archer7", 614, "DarkArcher"},
	&NPCType{"cr_lancer6", 615, "VileLancer"},
	&NPCType{"cr_lancer7", 616, "DarkLancer"},
	&NPCType{"cr_lancer8", 617, "BlackLancer"},
	&NPCType{"blunderbore5", 618, "Blunderbore"},
	&NPCType{"blunderbore6", 619, "Mauler"},
	&NPCType{"skmage_fire5", 620, "ReturnedMage"},
	&NPCType{"skmage_fire6", 621, "BurningDeadMage"},
	&NPCType{"skmage_ltng5", 622, "ReturnedMage"},
	&NPCType{"skmage_ltng6", 623, "HorrorMage"},
	&NPCType{"skmage_cold5", 624, "BoneMage"},
	&NPCType{"skmage_pois5", 625, "HorrorMage"},
	&NPCType{"skmage_pois6", 626, "HorrorMage"},
	&NPCType{"pantherwoman5", 627, "Huntress"},
	&NPCType{"pantherwoman6", 628, "SaberCat"},
	&NPCType{"sandleaper6", 629, "CaveLeaper"},
	&NPCType{"sandleaper7", 630, "TombCreeper"},
	&NPCType{"wraith6", 631, "Ghost"},
	&NPCType{"wraith7", 632, "Wraith"},
	&NPCType{"wraith8", 633, "Specter"},
	&NPCType{"succubus6", 634, "Succubusexp"},
	&NPCType{"succubus7", 635, "Hell Temptress"},
	&NPCType{"succubuswitch6", 636, "Dominus"},
	&NPCType{"succubuswitch7", 637, "Hell Witch"},
	&NPCType{"succubuswitch8", 638, "VileWitch"},
	&NPCType{"willowisp5", 639, "Gloam"},
	&NPCType{"willowisp6", 640, "BlackSoul"},
	&NPCType{"willowisp7", 641, "BurningSoul"},
	&NPCType{"fallen6", 642, "Carver"},
	&NPCType{"fallen7", 643, "Devilkin"},
	&NPCType{"fallen8", 644, "DarkOne"},
	&NPCType{"fallenshaman6", 645, "CarverShaman"},
	&NPCType{"fallenshaman7", 646, "DevilkinShaman"},
	&NPCType{"fallenshaman8", 647, "DarkShaman"},
	&NPCType{"skeleton6", 648, "BoneWarrior"},
	&NPCType{"skeleton7", 649, "Returned"},
	&NPCType{"batdemon6", 650, "Gloombat"},
	&NPCType{"batdemon7", 651, "Fiend"},
	&NPCType{"bloodlord6", 652, "Blood Lord1"},
	&NPCType{"bloodlord7", 653, "Blood Lord4"},
	&NPCType{"scarab6", 654, "Scarab"},
	&NPCType{"scarab7", 655, "SteelWeevil"},
	&NPCType{"fetish6", 656, "Flayer"},
	&NPCType{"fetish7", 657, "StygianDoll"},
	&NPCType{"fetish8", 658, "SoulKiller"},
	&NPCType{"fetishblow6", 659, "Flayer"},
	&NPCType{"fetishblow7", 660, "StygianDoll"},
	&NPCType{"fetishblow8", 661, "SoulKiller"},
	&NPCType{"fetishshaman6", 662, "FlayerShaman"},
	&NPCType{"fetishshaman7", 663, "StygianDollShaman"},
	&NPCType{"fetishshaman8", 664, "SoulKillerShaman"},
	&NPCType{"baboon7", 665, "TempleGuard"},
	&NPCType{"baboon8", 666, "TempleGuard"},
	&NPCType{"unraveler6", 667, "Guardian"},
	&NPCType{"unraveler7", 668, "Unraveler"},
	&NPCType{"unraveler8", 669, "Horadrim Ancient"},
	&NPCType{"unraveler9", 670, "Horadrim Ancient"},
	&NPCType{"zealot4", 671, "Zealot"},
	&NPCType{"zealot5", 672, "Zealot"},
	&NPCType{"cantor5", 673, "Heirophant"},
	&NPCType{"cantor6", 674, "Heirophant"},
	&NPCType{"vilemother4", 675, "Grotesque"},
	&NPCType{"vilemother5", 676, "FleshSpawner"},
	&NPCType{"vilechild4", 677, "GrotesqueWyrm"},
	&NPCType{"vilechild5", 678, "FleshBeast"},
	&NPCType{"sandmaggot6", 679, "WorldKiller"},
	&NPCType{"maggotbaby6", 680, "WorldKillerYoung"},
	&NPCType{"maggotegg6", 681, "WorldKillerEgg"},
	&NPCType{"minion9", 682, "Slayerexp"},
	&NPCType{"minion10", 683, "HellSpawn"},
	&NPCType{"minion11", 684, "GreaterHellSpawn"},
	&NPCType{"arach6", 685, "Arach"},
	&NPCType{"megademon4", 686, "Balrog"},
	&NPCType{"megademon5", 687, "PitLord"},
	&NPCType{"imp6", 688, "Imp1"},
	&NPCType{"imp7", 689, "Imp4"},
	&NPCType{"bonefetish6", 690, "Undead StygianDoll"},
	&NPCType{"bonefetish7", 691, "Undead SoulKiller"},
	&NPCType{"fingermage4", 692, "Strangler"},
	&NPCType{"fingermage5", 693, "StormCaster"},
	&NPCType{"regurgitator4", 694, "MawFiend"},
	&NPCType{"vampire6", 695, "BloodLord"},
	&NPCType{"vampire7", 696, "GhoulLord"},
	&NPCType{"vampire8", 697, "DarkLord"},
	&NPCType{"reanimatedhorde6", 698, "UnholyCorpse"},
	&NPCType{"dkfig1", 699, "DoomKnight"},
	&NPCType{"dkfig2", 700, "DoomKnight"},
	&NPCType{"dkmag1", 701, "OblivionKnight"},
	&NPCType{"dkmag2", 702, "OblivionKnight"},
	&NPCType{"mummy6", 703, "Cadaver"},
}