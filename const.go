package d2prox

import "fmt"

//
// enums & constants
//

const JoinGameOk = 0x00

//
// battle.net server packets
//

const SidAuthInfo = 0x50
const SidAuthCheck = 0x51
const SidLogonResponse2 = 0x3A
const SidLogonRealmEx = 0x3E

//
// realm server packets
//

const McpStartup = 0x01
const McpCreateGame = 0x03
const McpJoinGame = 0x04

//
// unit type
//

type UnitType int

const UnitTypePlayer = UnitType(0x00)
const UnitTypeNPC = UnitType(0x01)
const UnitTypeObject = UnitType(0x02)
const UnitTypeMissile = UnitType(0x03)
const UnitTypeItem = UnitType(0x04)
const UnitTypeWarp = UnitType(0x05)

//
// player class
//

type PlayerClass int

const ClassAmazon = PlayerClass(0x00)
const ClassSorceress = PlayerClass(0x01)
const ClassNecromancer = PlayerClass(0x2)
const ClassPaladin = PlayerClass(0x03)
const ClassBarbarian = PlayerClass(0x4)
const ClassDruid = PlayerClass(0x5)
const ClassAssassin = PlayerClass(0x6)

func (c PlayerClass) String() string {
	switch c {
	case ClassAmazon:
		return "Amazon"
	case ClassSorceress:
		return "Sorceress"
	case ClassNecromancer:
		return "Necromancer"
	case ClassPaladin:
		return "Paladin"
	case ClassBarbarian:
		return "Barbarian"
	case ClassDruid:
		return "Druid"
	case ClassAssassin:
		return "Assassin"
	}
	return fmt.Sprintf("Unknown (%d)", c)
}

//
// item quality
//

type Quality int

const QualityNone = Quality(0x00)
const QualityInferior = Quality(0x01)
const QualityNormal = Quality(0x02)
const QualitySuperior = Quality(0x03)
const QualityMagic = Quality(0x04)
const QualitySet = Quality(0x05)
const QualityRare = Quality(0x06)
const QualityUnique = Quality(0x07)
const QualityCraft = Quality(0x08)

func (q Quality) String() string {
	switch q {
	case QualityInferior:
		return "Inferior"
	case QualityNormal:
		return ""
	case QualitySuperior:
		return "Superior"
	case QualityMagic:
		return "Magic"
	case QualitySet:
		return "Set"
	case QualityRare:
		return "Rare"
	case QualityUnique:
		return "Unique"
	case QualityCraft:
		return "Crafted"
	}
	return fmt.Sprintf("None [%d]", int(q))
}

//
// item actions
//

type ItemAction int

const ItemActionNewGround = ItemAction(0x00)
const ItemActionPickup = ItemAction(0x01)
const ItemActionDrop = ItemAction(0x02)
const ItemActionOldGround = ItemAction(0x03)
const ItemActionToStorage = ItemAction(0x04)
const ItemActionFromStorage = ItemAction(0x05)
const ItemActionEquip = ItemAction(0x06)
const ItemActionIndirectSwapBody = ItemAction(0x07)
const ItemActionUnequip = ItemAction(0x08)
const ItemActionSwapBody = ItemAction(0x09)
const ItemActionAddQuantity = ItemAction(0x0a)
const ItemActionToStore = ItemAction(0x0b)
const ItemActionFromStore = ItemAction(0x0c)
const ItemActionSwapInContainer = ItemAction(0x0d)
const ItemActionPlaceBelt = ItemAction(0x0e)
const ItemActionRemoveBelt = ItemAction(0x0f)
const ItemActionSwapBelt = ItemAction(0x10)
const ItemActionAutoUnequip = ItemAction(0x11)
const ItemActionToCursor = ItemAction(0x12)
const ItemActionItemInSocket = ItemAction(0x13)
const ItemActionUpdateStats = ItemAction(0x15)
const ItemActionWeaponSwitch = ItemAction(0x17)

func (a ItemAction) String() string {
	switch a {
	case ItemActionDrop:
		return "Drop"
	case ItemActionEquip:
		return "Equip"
	case ItemActionPickup:
		return "Pickup"
	case ItemActionToCursor:
		return "ToCursor"
	}
	return fmt.Sprintf("ItemAction(%X)", int(a))
}

//
// item group
//

const ItemGroupAllArmor = 0x20000000
const ItemGroupAllWeapons = 0x40000000

//
// stats
//

const STAT_STRENGTH = 0  // str
const STAT_ENERGY = 1    // energy
const STAT_DEXTERITY = 2 // dexterity
const STAT_VITALITY = 3  // vitality
const STAT_STATPOINTSLEFT = 4
const STAT_NEWSKILLS = 5
const STAT_HP = 6          // life
const STAT_MAXHP = 7       // max life
const STAT_MANA = 8        // mana
const STAT_MAXMANA = 9     // max mana
const STAT_STAMINA = 10    // stamina
const STAT_MAXSTAMINA = 11 // max stamina
const STAT_LEVEL = 12      // level
const STAT_EXP = 13        // experience
const STAT_GOLD = 14       // gold
const STAT_GOLDBANK = 15   // stash gold
const STAT_ENHANCEDDEFENSE = 16
const STAT_ENHANCEDMAXIMUMDAMAGE = 17
const STAT_ENHANCEDMINIMUMDAMAGE = 18
const STAT_ATTACKRATING = 19
const STAT_TOBLOCK = 20 // to block
const STAT_MINIMUMDAMAGE = 21
const STAT_MAXIMUMDAMAGE = 22
const STAT_SECONDARYMINIMUMDAMAGE = 23
const STAT_SECONDARYMAXIMUMDAMAGE = 24
const STAT_ENHANCEDDAMAGE = 25
const STAT_MANARECOVERY = 26
const STAT_MANARECOVERYBONUS = 27
const STAT_STAMINARECOVERYBONUS = 28
const STAT_LASTEXPERIENCE = 29
const STAT_NEXTEXPERIENCE = 30
const STAT_DEFENSE = 31
const STAT_DEFENSEVSMISSILES = 32
const STAT_DEFENSEVSMELEE = 33
const STAT_DMGREDUCTION = 34
const STAT_MAGICDMGREDUCTION = 35    // magic damage reduction
const STAT_DMGREDUCTIONPCT = 36      // damage reduction
const STAT_MAGICDMGREDUCTIONPCT = 37 // magic damage reduction percentage
const STAT_MAXMAGICDMGREDUCTPCT = 38 // max magic damage reduction percentage
const STAT_FIRERESIST = 39           // fire resist
const STAT_MAXFIRERESIST = 40        // max fire resist
const STAT_LIGHTNINGRESIST = 41      // lightning resist
const STAT_MAXLIGHTNINGRESIST = 42   // max lightning resist
const STAT_COLDRESIST = 43           // cold resist
const STAT_MAXCOLDRESIST = 44        // max cold resist
const STAT_POISONRESIST = 45         // poison resist
const STAT_MAXPOISONRESIST = 46      // max poison resist
const STAT_DAMAGEAURA = 47
const STAT_MINIMUMFIREDAMAGE = 48
const STAT_MAXIMUMFIREDAMAGE = 49
const STAT_MINIMUMLIGHTNINGDAMAGE = 50
const STAT_MAXIMUMLIGHTNINGDAMAGE = 51
const STAT_MINIMUMMAGICALDAMAGE = 52
const STAT_MAXIMUMMAGICALDAMAGE = 53
const STAT_MINIMUMCOLDDAMAGE = 54
const STAT_MAXIMUMCOLDDAMAGE = 55
const STAT_COLDDAMAGELENGTH = 56
const STAT_MINIMUMPOISONDAMAGE = 57
const STAT_MAXIMUMPOISONDAMAGE = 58
const STAT_POISONDAMAGELENGTH = 59
const STAT_LIFELEECH = 60 // Life Leech (min life stolen per hit)
const STAT_MAXLIFESTOLENPERHIT = 61
const STAT_MANALEECH = 62 // Mana Leech (min mana stolen per hit)
const STAT_MAXMANASTOLENPERHIT = 63
const STAT_MINIMUMSTAMINADRAIN = 64
const STAT_MAXIMUMSTAMINADRAIN = 65
const STAT_STUNLENGTH = 66
const STAT_VELOCITYPERCENT = 67 // effective run/walk
const STAT_ATTACKRATE = 68
const STAT_OTHERANIMATIONRATE = 69
const STAT_AMMOQUANTITY = 70 // ammo quantity(arrow/bolt/throwing)
const STAT_VALUE = 71
const STAT_DURABILITY = 72    // item durability
const STAT_MAXDURABILITY = 73 // max item durability
const STAT_REPLENISHLIFE = 74
const STAT_ENHANCEDMAXDURABILITY = 75
const STAT_ENHANCEDLIFE = 76
const STAT_ENHANCEDMANA = 77
const STAT_ATTACKERTAKESDAMAGE = 78
const STAT_GOLDFIND = 79  // Gold find (GF)
const STAT_MAGICFIND = 80 // magic find (MF)
const STAT_KNOCKBACK = 81
const STAT_TIMEDURATION = 82
const STAT_CLASSSKILLS = 83
const STAT_UNSENTPARAMETER = 84
const STAT_ADDEXPERIENCE = 85
const STAT_LIFEAFTEREACHKILL = 86
const STAT_REDUCEVENDORPRICES = 87
const STAT_DOUBLEHERBDURATION = 88
const STAT_LIGHTRADIUS = 89
const STAT_LIGHTCOLOUR = 90
const STAT_REDUCEDREQUIREMENTS = 91
const STAT_REDUCEDLEVELREQ = 92
const STAT_IAS = 93 // IAS
const STAT_REDUCEDLEVELREQPCT = 94
const STAT_LASTBLOCKFRAME = 95
const STAT_FASTERRUNWALK = 96 // faster run/walk
const STAT_NONCLASSSKILL = 97
const STAT_STATE = 98
const STAT_FASTERHITRECOVERY = 99 // faster hit recovery
const STAT_MONSTERPLAYERCOUNT = 100
const STAT_SKILLPOISONOVERRIDELEN = 101
const STAT_FASTERBLOCK = 102 // faster block rate
const STAT_SKILLBYPASSUNDEAD = 103
const STAT_SKILLBYPASSDEMONS = 104
const STAT_FASTERCAST = 105 // faster cast rate
const STAT_SKILLBYPASSBEASTS = 106
const STAT_SINGLESKILL = 107
const STAT_SLAINMONSTERSRIP = 108
const STAT_CURSERESISTANCE = 109
const STAT_POISONLENGTHREDUCTION = 110 // Poison length reduction
const STAT_ADDSDAMAGE = 111
const STAT_HITCAUSESMONSTERTOFLEE = 112
const STAT_HITBLINDSTARGET = 113
const STAT_DAMAGETOMANA = 114
const STAT_IGNORETARGETSDEFENSE = 115
const STAT_REDUCETARGETSDEFENSE = 116
const STAT_PREVENTMONSTERHEAL = 117
const STAT_HALFFREEZEDURATION = 118
const STAT_TOHITPERCENT = 119
const STAT_MONSTERDEFDUCTPERHIT = 120
const STAT_DAMAGETODEMONS = 121
const STAT_DAMAGETOUNDEAD = 122
const STAT_ATTACKRATINGVSDEMONS = 123
const STAT_ATTACKRATINGVSUNDEAD = 124
const STAT_THROWABLE = 125
const STAT_ELEMENTALSKILLS = 126
const STAT_ALLSKILLS = 127
const STAT_ATTACKERTAKESLTNGDMG = 128
const STAT_IRONMAIDENLEVEL = 129
const STAT_LIFETAPLEVEL = 130
const STAT_THORNSPERCENT = 131
const STAT_BONEARMOR = 132
const STAT_MAXIMUMBONEARMOR = 133
const STAT_FREEZESTARGET = 134
const STAT_OPENWOUNDS = 135   // Open Wounds
const STAT_CRUSHINGBLOW = 136 // crushing blow
const STAT_KICKDAMAGE = 137
const STAT_MANAAFTEREACHKILL = 138
const STAT_LIFEAFTEREACHDEMONKILL = 139
const STAT_EXTRABLOOD = 140
const STAT_DEADLYSTRIKE = 141           // deadly strike
const STAT_FIREABSORBPERCENT = 142      // fire absorb %
const STAT_FIREABSORB = 143             // fire absorb
const STAT_LIGHTNINGABSORBPERCENT = 144 // lightning absorb %
const STAT_LIGHTNINGABSORB = 145        // lightning absorb
const STAT_MAGICABSORBPERCENT = 146     // magic absorb %
const STAT_MAGICABSORB = 147
const STAT_COLDABSORBPERCENT = 148 // cold absorb %
const STAT_COLDABSORB = 149        // cold absorb
const STAT_SLOW = 150              // slow %
const STAT_AURA = 151
const STAT_INDESTRUCTIBLE = 152
const STAT_CANNOTBEFROZEN = 153
const STAT_STAMINADRAINPERCENT = 154
const STAT_REANIMATE = 155
const STAT_PIERCINGATTACK = 156
const STAT_FIRESMAGICARROWS = 157
const STAT_FIREEXPLOSIVEARROWS = 158
const STAT_MINIMUMTHROWINGDAMAGE = 159
const STAT_MAXIMUMTHROWINGDAMAGE = 160
const STAT_SKILLHANDOFATHENA = 161
const STAT_SKILLSTAMINAPERCENT = 162
const STAT_SKILLPASSIVESTAMINAPCT = 163
const STAT_CONCENTRATION = 164
const STAT_ENCHANT = 165
const STAT_PIERCE = 166
const STAT_CONVICTION = 167
const STAT_CHILLINGARMOR = 168
const STAT_FRENZY = 169
const STAT_DECREPIFY = 170
const STAT_SKILLARMORPERCENT = 171
const STAT_ALIGNMENT = 172
const STAT_TARGET0 = 173
const STAT_TARGET1 = 174
const STAT_GOLDLOST = 175
const STAT_CONVERSIONLEVEL = 176
const STAT_CONVERSIONMAXIMUMLIFE = 177
const STAT_UNITDOOVERLAY = 178
const STAT_ATTCKRTNGVSMONSTERTYPE = 179
const STAT_DAMAGETOMONSTERTYPE = 180
const STAT_FADE = 181
const STAT_ARMOROVERRIDEPERCENT = 182
const STAT_UNUSED183 = 183
const STAT_UNUSED184 = 184
const STAT_UNUSED185 = 185
const STAT_UNUSED186 = 186
const STAT_UNUSED187 = 187
const STAT_SKILLTAB = 188
const STAT_UNUSED189 = 189
const STAT_UNUSED190 = 190
const STAT_UNUSED191 = 191
const STAT_UNUSED192 = 192
const STAT_UNUSED193 = 193
const STAT_SOCKETS = 194
const STAT_SKILLONSTRIKING = 195
const STAT_SKILLONKILL = 196
const STAT_SKILLONDEATH = 197
const STAT_SKILLONHIT = 198
const STAT_SKILLONLEVELUP = 199
const STAT_UNUSED200 = 200
const STAT_SKILLWHENSTRUCK = 201
const STAT_UNUSED202 = 202
const STAT_UNUSED203 = 203
const STAT_CHARGED = 204
const STAT_UNUSED204 = 205
const STAT_UNUSED205 = 206
const STAT_UNUSED206 = 207
const STAT_UNUSED207 = 208
const STAT_UNUSED208 = 209
const STAT_UNUSED209 = 210
const STAT_UNUSED210 = 211
const STAT_UNUSED211 = 212
const STAT_UNUSED212 = 213
const STAT_DEFENSEPERLEVEL = 214
const STAT_ENHANCEDDEFENSEPERLVL = 215
const STAT_LIFEPERLEVEL = 216
const STAT_MANAPERLEVEL = 217
const STAT_MAXDAMAGEPERLEVEL = 218
const STAT_MAXENHANCEDDMGPERLEVEL = 219
const STAT_STRENGTHPERLEVEL = 220
const STAT_DEXTERITYPERLEVEL = 221
const STAT_ENERGYPERLEVEL = 222
const STAT_VITALITYPERLEVEL = 223
const STAT_ATTACKRATINGPERLEVEL = 224
const STAT_BONUSATTCKRTNGPERLEVEL = 225
const STAT_MAXCOLDDMGPERLVL = 226
const STAT_MAXFIREDMGPERLVL = 227
const STAT_MAXLIGHTNINGDMGPERLVL = 228
const STAT_MAXPOISONDMGPERLVL = 229
const STAT_COLDRESPERLEVEL = 230
const STAT_FIRERESPERLEVEL = 231
const STAT_LIGHTNINGRESPERLEVEL = 232
const STAT_POISONRESPERLEVEL = 233
const STAT_COLDABSORBPERLVL = 234
const STAT_FIREABSORBPERLVL = 235
const STAT_LIGHTNINGABSORBPERLVL = 236
const STAT_POISONABSORBPERLVL = 237
const STAT_THORNSPERLEVEL = 238
const STAT_EXTRAGOLDPERLEVEL = 239
const STAT_MAGICFINDPERLEVEL = 240
const STAT_STAMINAREGENPERLEVEL = 241
const STAT_STAMINAPERLEVEL = 242
const STAT_DAMAGETODEMONSPERLEVEL = 243
const STAT_DAMAGETOUNDEADPERLEVEL = 244
const STAT_ATTKRTNGVSDEMONSPERLVL = 245
const STAT_ATTKRTNGVSUNDEADPERLVL = 246
const STAT_CRUSHINGBLOWPERLEVEL = 247
const STAT_OPENWOUNDSPERLEVEL = 248
const STAT_KICKDAMAGEPERLEVEL = 249
const STAT_DEADLYSTRIKEPERLEVEL = 250
const STAT_FINDGEMSPERLEVEL = 251
const STAT_REPAIRSDURABILITY = 252
const STAT_REPLENISHESQUANTITY = 253
const STAT_INCREASEDSTACKSIZE = 254
const STAT_FINDITEM = 255
const STAT_SLASHDAMAGE = 256
const STAT_SLASHDAMAGEPERCENT = 257
const STAT_CRUSHDAMAGE = 258
const STAT_CRUSHDAMAGEPERCENT = 259
const STAT_THRUSTDAMAGE = 260
const STAT_THRUSTDAMAGEPERCENT = 261
const STAT_SLASHDAMAGEABSORPTION = 262
const STAT_CRUSHDAMAGEABSORPTION = 263
const STAT_THRUSTDAMAGEABSORPTION = 264
const STAT_SLASHDAMAGEABSORBPCT = 265
const STAT_CRUSHDAMAGEABSORBPCT = 266
const STAT_THRUSTDAMAGEABSORBPCT = 267
const STAT_DEFENSEPERTIME = 268
const STAT_ENHANCEDDEFENSEPERTIME = 269
const STAT_LIFEPERTIME = 270
const STAT_MANAPERTIME = 271
const STAT_MAXDAMAGEPERTIME = 272
const STAT_MAXENHANCEDDMGPERTIME = 273
const STAT_STRENGTHPERTIME = 274
const STAT_DEXTERITYPERTIME = 275
const STAT_ENERGYPERTIME = 276
const STAT_VITALITYPERTIME = 277
const STAT_ATTACKRATINGPERTIME = 278
const STAT_CHANCETOHITPERTIME = 279
const STAT_MAXCOLDDAMAGEPERTIME = 280
const STAT_MAXFIREDAMAGEPERTIME = 281
const STAT_MAXLIGHTNINGDMGPERTIME = 282
const STAT_MAXDAMAGEPERPOISON = 283
const STAT_COLDRESPERTIME = 284
const STAT_FIRERESPERTIME = 285
const STAT_LIGHTNINGRESPERTIME = 286
const STAT_POISONRESPERTIME = 287
const STAT_COLDABSORPTIONPERTIME = 288
const STAT_FIREABSORPTIONPERTIME = 289
const STAT_LIGHTNINGABSORBPERTIME = 290
const STAT_POISONABSORBPERTIME = 291
const STAT_EXTRAGOLDPERTIME = 292
const STAT_MAGICFINDPERTIME = 293
const STAT_REGENSTAMINAPERTIME = 294
const STAT_STAMINAPERTIME = 295
const STAT_DAMAGETODEMONSPERTIME = 296
const STAT_DAMAGETOUNDEADPERTIME = 297
const STAT_ATTRTNGVSDEMONSPERTIME = 298
const STAT_ATTRTNGVSUNDEADPERTIME = 299
const STAT_CRUSHINGBLOWPERTIME = 300
const STAT_OPENWOUNDSPERTIME = 301
const STAT_KICKDAMAGEPERTIME = 302
const STAT_DEADLYSTRIKEPERTIME = 303
const STAT_FINDGEMSPERTIME = 304
const STAT_ENEMYCOLDRESREDUCTION = 305
const STAT_ENEMYFIRERESREDUCTION = 306
const STAT_ENEMYLIGHTRESREDUCTION = 307
const STAT_ENEMYPSNRESREDUCTION = 308
const STAT_DAMAGEVSMONSTERS = 309
const STAT_ENHANCEDDMGVSMONSTERS = 310
const STAT_ATTACKRATINGVSMONSTERS = 311
const STAT_BONUSATTRTNGVSMONSTERS = 312
const STAT_DEFENSEVSMONSTERS = 313
const STAT_ENHANCEDDEFVSMONSTERS = 314
const STAT_FIREDAMAGELENGTH = 315
const STAT_MINFIREDAMAGELENGTH = 316
const STAT_MAXFIREDAMAGELENGTH = 317
const STAT_PROGRESSIVEDAMAGE = 318
const STAT_PROGRESSIVESTEAL = 319
const STAT_PROGRESSIVEOTHER = 320
const STAT_PROGRESSIVEFIRE = 321
const STAT_PROGRESSIVECOLD = 322
const STAT_PROGRESSIVELIGHTNING = 323
const STAT_EXTRACHARGES = 324
const STAT_PROGRESSIVEATTACKRTNG = 325
const STAT_POISONCOUNT = 326
const STAT_DAMAGEFRAMERATE = 327
const STAT_PIERCEIDX = 328
const STAT_FIREMASTERY = 329
const STAT_LIGHTNINGMASTERY = 330
const STAT_COLDMASTERY = 331
const STAT_POISONMASTERY = 332
const STAT_PSENEMYFIRERESREDUC = 333   // passive enemy fire resist reduction
const STAT_PSENEMYLIGHTNRESREDUC = 334 // passive enemy lightning resist reduction
const STAT_PSENEMYCOLDRESREDUC = 335   // passive enemy cold resist reduction
const STAT_PSENEMYPSNRESREDUC = 336    // passive enemy poison resist reduction
const STAT_CRITICALSTRIKE = 337
const STAT_DODGE = 338
const STAT_AVOID = 339
const STAT_EVADE = 340
const STAT_WARMTH = 341
const STAT_MELEEARMASTERY = 342 // melee attack rating mastery
const STAT_MELEEDAMAGEMASTERY = 343
const STAT_MELEECRITHITMASTERY = 344
const STAT_THROWNWEAPONARMASTERY = 345 // thrown weapon attack rating mastery
const STAT_THROWNWEAPONDMGMASTERY = 346
const STAT_THROWNCRITHITMASTERY = 347 // thrown weapon critical hit mastery
const STAT_WEAPONBLOCK = 348
const STAT_SUMMONRESIST = 349
const STAT_MODIFIERLISTSKILL = 350
const STAT_MODIFIERLISTLEVEL = 351
const STAT_LASTSENTLIFEPERCENT = 352
const STAT_SOURCEUNITTYPE = 353
const STAT_SOURCEUNITID = 354
const STAT_SHORTPARAMETER1 = 355
const STAT_QUESTITEMDIFFICULTY = 356
const STAT_PASSIVEMAGICDMGMASTERY = 357
const STAT_PASSIVEMAGICRESREDUC = 358
