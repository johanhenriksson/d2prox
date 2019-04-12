package d2prox

import "fmt"

//
// client -> server
//

const GsWalk = 0x01
const GsWalkToEntity = 0x02
const GsRun = 0x03
const GsRunToEntity = 0x04
const GsShiftLeftSkill = 0x05
const GsLeftSkillOnUnit = 0x06
const GsShiftLeftSkillOnUnit = 0x07
const GsShiftLeftSkillHold = 0x08
const GsLeftSkillOnUnitHold = 0x09
const GsShiftLeftSkillOnUnitHold = 0x0A
const GsUnknown12 = 0x0B
const GsRightSkill = 0x0C
const GsRightSkillOnUnit = 0x0D
const GsShiftRightSkillOnUnit = 0x0E
const GsRightSkillHold = 0x0F
const GsRightSkillOnUnitHold = 0x10
const GsShiftRightSkillOnUnitHold = 0x11
const GsSetInfernoState = 0x12
const GsInteractEntity = 0x13
const GsOverheadChat = 0x14
const GsChatMessage = 0x15
const GsPickupItem = 0x16
const GsDropItem = 0x17
const GsBufferItem = 0x18
const GsRemoveBufferItem = 0x19
const GsEquipItem = 0x1A
const GsSwap2HandItem = 0x1B
const GsRemoveBodyItem = 0x1C
const GsSwapCursorItemWithBody = 0x1D
const GsSwapItemsWith2HandItem = 0x1E
const GsSwapCursorBufferItems = 0x1F
const GsActivateBufferItem = 0x20
const GsStackItems = 0x21
const GsUnstackItems = 0x22
const GsItemToBelt = 0x23
const GsItemFromBelt = 0x24
const GsSwitchBeltItem = 0x25
const GsUseBeltItem = 0x26
const GsIdentifyItem = 0x27
const GsSocketItem = 0x28
const GsScrollToBook = 0x29
const GsItemToCube = 0x2A
const GsInitateEntityChat = 0x2F
const GsTerminateEntityChat = 0x30
const GsQuestMessage = 0x31
const GsBuyItem = 0x32
const GsSellItem = 0x33
const GsIdentifyItemsWithNPC = 0x34
const GsRepair = 0x35
const GsHireMerc = 0x36
const GsIdentifyFromGamble = 0x37
const GsEntityAction = 0x38
const GsPurchaseLife = 0x39
const GsAddStatPoint = 0x3A
const GsAddSkillPoint = 0x3B
const GsSelectSkill = 0x3C
const GsHighlightDoor = 0x3D
const GsActivateScrollOfInifuss = 0x3E
const GsPlayAudio = 0x3F
const GsRequestQuestData = 0x40
const GsResurrect = 0x41
const GsStaffInOrfice = 0x44
const GsChangeTPLocation = 0x45
const GsHaveMercInteract = 0x46
const GsMoveMerc = 0x47
const GsRemoveBusyState = 0x48
const GsWaypointInteract = 0x49
const GsRequestEntityUpdate = 0x4B
const GsTransmorgify = 0x4C
const GsPlayNPCMessage = 0x4D
const GsClickButton = 0x4F
const GsDropGold = 0x50
const GsBindHotkeySkill = 0x51
const GsUnknown13 = 0x52
const GsTurnStaminaOn = 0x53
const GsTurnStaminaOff = 0x54
const GsQuestCompleted = 0x58
const GsMakeEntityMove = 0x59
const GsSquelchHostile = 0x5D
const GsInviteParty = 0x5E
const GsUpdatePlayerPos = 0x5F
const GsSwapWeapons = 0x60
const GsPickupMercItem = 0x61
const GsResurrectMerc = 0x62
const GsShiftLeftClickItemToBelt = 0x63
const GsHackDetection1 = 0x64
const GsHackDetection2 = 0x65
const GsWardenResponse = 0x66
const GsGameLogon = 0x68
const GsLeaveGame = 0x69
const GsJoinGame = 0x6B
const GsUploadSave = 0x6C
const GsPing = 0x6D

var gsClientPacketLengths = map[int]int{
	GsWalk:                      5,
	GsWalkToEntity:              9,
	GsRun:                       5,
	GsRunToEntity:               9,
	GsShiftLeftSkill:            5,
	GsLeftSkillOnUnit:           9,
	GsShiftLeftSkillOnUnit:      9,
	GsShiftLeftSkillHold:        5,
	GsLeftSkillOnUnitHold:       9,
	GsShiftLeftSkillOnUnitHold:  9,
	GsUnknown12:                 1,
	GsRightSkill:                5,
	GsRightSkillOnUnit:          9,
	GsShiftRightSkillOnUnit:     9,
	GsRightSkillHold:            5,
	GsRightSkillOnUnitHold:      9,
	GsShiftRightSkillOnUnitHold: 9,
	GsSetInfernoState:           1,
	GsInteractEntity:            9,
	GsOverheadChat:              -1, // variable
	GsChatMessage:               -1, // variable
	GsPickupItem:                13,
	GsDropItem:                  5,
	GsBufferItem:                17,
	GsRemoveBufferItem:          5,
	GsEquipItem:                 9,
	GsSwap2HandItem:             9,
	GsRemoveBodyItem:            3,
	GsSwapCursorItemWithBody:    9,
	GsSwapItemsWith2HandItem:    9,
	GsSwapCursorBufferItems:     17,
	GsActivateBufferItem:        13,
	GsStackItems:                9,
	GsUnstackItems:              5,
	GsItemToBelt:                9,
	GsItemFromBelt:              5,
	GsSwitchBeltItem:            9,
	GsUseBeltItem:               13,
	GsIdentifyItem:              9,
	GsSocketItem:                9,
	GsScrollToBook:              9,
	GsItemToCube:                9,
	GsInitateEntityChat:         9,
	GsTerminateEntityChat:       9,
	GsQuestMessage:              9,
	GsBuyItem:                   17,
	GsSellItem:                  17,
	GsIdentifyItemsWithNPC:      5,
	GsRepair:                    17,
	GsHireMerc:                  9,
	GsIdentifyFromGamble:        5,
	GsEntityAction:              13,
	GsPurchaseLife:              5,
	GsAddStatPoint:              3,
	GsAddSkillPoint:             3,
	GsSelectSkill:               9,
	GsHighlightDoor:             5,
	GsActivateScrollOfInifuss:   5,
	GsPlayAudio:                 3,
	GsRequestQuestData:          1,
	GsResurrect:                 1,
	GsStaffInOrfice:             17,
	GsChangeTPLocation:          9,
	GsHaveMercInteract:          13,
	GsMoveMerc:                  13,
	GsRemoveBusyState:           1,
	GsWaypointInteract:          9,
	GsRequestEntityUpdate:       9,
	GsTransmorgify:              5,
	GsPlayNPCMessage:            3,
	GsClickButton:               7,
	GsDropGold:                  9,
	GsBindHotkeySkill:           9,
	GsUnknown13:                 5,
	GsTurnStaminaOn:             1,
	GsTurnStaminaOff:            1,
	GsQuestCompleted:            3,
	GsMakeEntityMove:            17,
	GsSquelchHostile:            7,
	GsInviteParty:               6,
	GsUpdatePlayerPos:           5,
	GsSwapWeapons:               1,
	GsPickupMercItem:            3,
	GsResurrectMerc:             5,
	GsShiftLeftClickItemToBelt:  5,
	GsHackDetection1:            9,
	GsHackDetection2:            1, // unknown or unused
	GsWardenResponse:            1, // unknown or unused
	GsLeaveGame:                 1,
	GsGameLogon:                 37,
	GsJoinGame:                  1,
	GsUploadSave:                -1, // variable
	GsPing:                      13, // 0x65
}

var gsClientPacketLengthFuncs = map[int]PacketLengthFunc{
	GsOverheadChat: func(buffer PacketBuffer, offset, length int) (int, error) {
		// todo: implement properly
		return length - offset, nil
	},
	GsChatMessage: func(buffer PacketBuffer, offset, length int) (int, error) {
		// todo: implement properly
		return length - offset, nil
	},
	GsUploadSave: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Byte(offset+1) + 6, nil
	},
}

func GsClientPacketLength(buffer PacketBuffer, offset, length int) (int, error) {
	msgID := buffer.Byte(offset)
	plen, known := gsClientPacketLengths[msgID]
	if !known {
		return 0, fmt.Errorf("Unknown GS packet (C->S): 0x%x", msgID)
	}

	if plen < 0 {
		// variable length packet
		lengthFunc, exists := gsClientPacketLengthFuncs[msgID]
		if !exists {
			return 0, fmt.Errorf("No length function for C->S packet 0x%X", msgID)
		}
		return lengthFunc(buffer, offset, length)
	}

	return plen, nil
}

var gsClientPacketNames = map[int]string{
	GsWalk:                      "Walk",
	GsWalkToEntity:              "WalkToEntity",
	GsRun:                       "Run",
	GsRunToEntity:               "RunToEntity",
	GsShiftLeftSkill:            "ShiftLeftSkill",
	GsLeftSkillOnUnit:           "LeftSkillOnUnit",
	GsShiftLeftSkillOnUnit:      "ShiftLeftSkillOnUnit",
	GsShiftLeftSkillHold:        "ShiftLeftSkillHold",
	GsLeftSkillOnUnitHold:       "LeftSkillOnUnitHold",
	GsShiftLeftSkillOnUnitHold:  "ShiftLeftSkillOnUnitHold",
	GsRightSkill:                "RightSkill",
	GsRightSkillOnUnit:          "RightSkillOnUnit",
	GsShiftRightSkillOnUnit:     "ShiftRightSkillOnUnit",
	GsRightSkillHold:            "RightSkillHold",
	GsRightSkillOnUnitHold:      "RightSkillOnUnitHold",
	GsShiftRightSkillOnUnitHold: "ShiftRightSkillOnUnitHold",
	GsSetInfernoState:           "SetInfernoState",
	GsInteractEntity:            "InteractEntity",
	GsOverheadChat:              "OverheadChat", // variable
	GsChatMessage:               "Chat",         // variable
	GsPickupItem:                "PickupItem",
	GsDropItem:                  "DropItem",
	GsBufferItem:                "BufferItem",
	GsRemoveBufferItem:          "RemoveBufferItem",
	GsEquipItem:                 "EquipItem",
	GsSwap2HandItem:             "Swap2HandItem",
	GsRemoveBodyItem:            "RemoveBodyItem",
	GsSwapCursorItemWithBody:    "SwapCursorItemWithBody",
	GsSwapItemsWith2HandItem:    "SwapItemsWith2HandItem",
	GsSwapCursorBufferItems:     "SwapCursorBufferItems",
	GsActivateBufferItem:        "ActivateBufferItem",
	GsStackItems:                "StackItems",
	GsUnstackItems:              "UnstackItems",
	GsItemToBelt:                "ItemToBelt",
	GsItemFromBelt:              "ItemFromBelt",
	GsSwitchBeltItem:            "SwitchBeltItem",
	GsUseBeltItem:               "UseBeltItem",
	GsIdentifyItem:              "IdentifyItem",
	GsSocketItem:                "SocketItem",
	GsScrollToBook:              "ScrollToBook",
	GsItemToCube:                "ItemToCube",
	GsInitateEntityChat:         "InitiateEntityChat",
	GsTerminateEntityChat:       "TerminateEntityChat",
	GsQuestMessage:              "QuestMessage",
	GsBuyItem:                   "BuyItem",
	GsSellItem:                  "SellItem",
	GsIdentifyItemsWithNPC:      "IdentifyItemsWithNPC",
	GsRepair:                    "Repair",
	GsHireMerc:                  "HireMerc",
	GsIdentifyFromGamble:        "IdentifyFromGable",
	GsEntityAction:              "EntityAction",
	GsPurchaseLife:              "PurchaseLife",
	GsAddStatPoint:              "AddStatPoint",
	GsAddSkillPoint:             "AddSkillPoint",
	GsSelectSkill:               "SelectSkill",
	GsHighlightDoor:             "HighlightDoor",
	GsActivateScrollOfInifuss:   "ActivateScrollOfInifuss",
	GsPlayAudio:                 "PlayAudio",
	GsRequestQuestData:          "RequestQuestData",
	GsResurrect:                 "Resurrect",
	GsStaffInOrfice:             "StaffInOrfice",
	GsChangeTPLocation:          "ChangeTPLocation",
	GsHaveMercInteract:          "HaveMercInteract",
	GsMoveMerc:                  "MoveMerc",
	GsRemoveBusyState:           "RemoveBusyState",
	GsWaypointInteract:          "WaypointInteract",
	GsRequestEntityUpdate:       "RequestEntityUpdate",
	GsTransmorgify:              "Transmorgify",
	GsPlayNPCMessage:            "PlayNPCMessage",
	GsClickButton:               "ClickButton",
	GsDropGold:                  "DropGold",
	GsBindHotkeySkill:           "BindHotkeySkill",
	GsTurnStaminaOn:             "TurnStaminaOn",
	GsTurnStaminaOff:            "TurnStaminaOff",
	GsQuestCompleted:            "QuestCompleted",
	GsMakeEntityMove:            "MakeEntityMove",
	GsSquelchHostile:            "Squelch/Hostile",
	GsInviteParty:               "InviteParty",
	GsUpdatePlayerPos:           "UpdatePlayerPos",
	GsSwapWeapons:               "SwapWeapons",
	GsPickupMercItem:            "PickupMercItem",
	GsResurrectMerc:             "ResurrectMerc",
	GsShiftLeftClickItemToBelt:  "ShiftClickItemToBelt",
	GsHackDetection1:            "HackDetection1",
	GsHackDetection2:            "HackDetection2",
	GsWardenResponse:            "WardenResponse",
	GsLeaveGame:                 "LeaveGame",
	GsGameLogon:                 "GameLogon",
	GsJoinGame:                  "JoinGame",
	GsUploadSave:                "UploadSave",
	GsPing:                      "Ping",
}

func GsClientPacketName(packet Packet) string {
	msgID := int(packet.GsMsgID())
	name, exists := gsClientPacketNames[msgID]
	if !exists {
		return fmt.Sprintf("Unknown (0x%X)", byte(msgID))
	}
	return name
}

//
// server -> client
//

const GsGameLoading = 0x00
const GsGameFlags = 0x01
const GsLoadSuccess = 0x02
const GsLoadAct = 0x03
const GsLoadComplete = 0x04
const GsUnloadComplete = 0x05
const GsGameExitSuccess = 0x06
const GsMapReveal = 0x07
const GsMapHide = 0x08
const GsAssignLevelWarp = 0x09
const GsRemoveObject = 0x0A
const GsHandshake = 0x0B
const GsNPCHit = 0x0C
const GsPlayerStop = 0x0D
const GsObjectState = 0x0E
const GsPlayerMove = 0x0F
const GsPlayerToTarget = 0x10
const GsReportKill = 0x11
const GsReassignPlayer = 0x15
const GsUnknown1 = 0x16
const GsUnknown2 = 0x17
const GsPlayerHpMp = 0x18
const GsGoldToInv8 = 0x19
const GsAddExp8 = 0x1A
const GsAddExp16 = 0x1B
const GsAddExp32 = 0x1C
const GsSetAttr8 = 0x1D
const GsSetAttr16 = 0x1E
const GsSetAttr32 = 0x1F
const GsAttributeUpdate = 0x20
const GsUpdateItemOSkill = 0x21
const GsUpdateItemSkill = 0x22
const GsSetSkill = 0x23
const GsGameChat = 0x26
const GsNPCInfo = 0x27
const GsQuestInfo = 0x28
const GsGameQuestInfo = 0x29
const GsNPCTransaction = 0x2A
const GsPlaySound = 0x2C
const GsUpdateItemStats = 0x3E
const GsUseStackableItem = 0x3F
const GsUnknown3 = 0x40
const GsClearCursor = 0x42
const GsRelator1 = 0x47
const GsRelator2 = 0x48
const GsUnitSkillOnTarget = 0x4C
const GsUnitCastSkill = 0x4D
const GsMercForHire = 0x4E
const GsClearMercList = 0x4F
const GsQuestSpecial = 0x50
const GsAssignObject = 0x51
const GsPlayerQuestLog = 0x52
const GsDarkness = 0x53
const GsNPCEnchants = 0x57
const GsUnknown4 = 0x58
const GsAssignPlayer = 0x59
const GsEventMessages = 0x5A
const GsPlayerInGame = 0x5B
const GsPlayerLeft = 0x5C
const GsQuestItemState = 0x5D
const GsGameQuestsAvailable = 0x5E
const GsUnknown5 = 0x5F
const GsTownPortalState = 0x60
const GsCanGoToAct = 0x61
const GsUnknown6 = 0x62
const GsWaypointMenu = 0x63
const GsPlayerKillCount = 0x65
const GsNPCMove = 0x67
const GsNPCMoveToTarget = 0x68
const GsNPCState = 0x69
const GsUnknown7 = 0x6A
const GsNPCAction = 0x6B
const GsNPCAttack = 0x6C
const GsNPCStop = 0x6D
const GsUnknown8 = 0x73
const GsPlayerCorpseAssign = 0x74
const GsPlayerPartyInfo = 0x75
const GsPlayerInProximity = 0x76
const GsButtonActions = 0x77
const GsTradeAccepted = 0x78
const GsGoldInTrade = 0x79
const GsPetAction = 0x7A
const GsAssignSkillHotkey = 0x7B
const GsUseScroll = 0x7C
const GsSetItemFlags = 0x7D
const GsCMNCOF = 0x7E
const GsAllyPartyInfo = 0x7F
const GsAssignMerc = 0x81
const GsPortalOwnership = 0x82
const GsSpecialQuestEvent = 0x89
const GsNPCWantsInteract = 0x8A
const GsPlayerRelationship = 0x8B
const GsRelationshipUpdate = 0x8C
const GsAssignPlayerToParty = 0x8D
const GsCorpseAssign = 0x8E
const GsPong = 0x8F
const GsPartyAutomapInfo = 0x90
const GsSetNPCGossip = 0x91
const GsRemoveUnitDisplay = 0x92
const GsUnknown9 = 0x93
const GsBaseSkillLevels = 0x94
const GsLifeManaUpdate = 0x95
const GsWalkVerify = 0x96
const GsWeaponSwitch = 0x97
const GsUpdateNPCUnknownField40 = 0x98
const GsCastSkillOnUnit = 0x99
const GsCastSkillOnLocation = 0x9A
const GsMercReviveCost = 0x9B
const GsItemActionWorld = 0x9C
const GsItemActionOwned = 0x9D
const GsSetMercStat8 = 0x9E
const GsSetMercStat16 = 0x9F
const GsSetMercStat32 = 0xA0
const GsAddMercExp8 = 0xA1
const GsAddMercExp16 = 0xA2
const GsSkillAuraStat = 0xA3
const GsNextBaalWaveClassID = 0xA4
const GsStateSkillMove = 0xA5
const GsUnknown10 = 0xA6
const GsDelayedState = 0xA7
const GsSetState = 0xA8
const GsEndState = 0xA9
const GsAddUnit = 0xAA
const GsNPCHeal = 0xAB
const GsAssignNPC = 0xAC
const GsWardenRequest = 0xAE
const GsConnectionInfo = 0xAF
const GsTerminated = 0xB0
const GsUnknown11 = 0xB2
const GsDownloadSave = 0xB3
const GsTimeout = 0xB4

var gsServerPacketLengths = map[int]int{
	GsGameLoading:             1,
	GsGameFlags:               8,
	GsLoadSuccess:             1,
	GsLoadAct:                 12,
	GsLoadComplete:            1,
	GsUnloadComplete:          1,
	GsGameExitSuccess:         1,
	GsMapReveal:               6,
	GsMapHide:                 6,
	GsAssignLevelWarp:         12,
	GsRemoveObject:            6,
	GsHandshake:               6,
	GsNPCHit:                  9,
	GsPlayerStop:              13,
	GsObjectState:             12,
	GsPlayerMove:              16,
	GsPlayerToTarget:          16,
	GsReportKill:              8,
	GsReassignPlayer:          11,
	GsUnknown1:                -1, // variable
	GsUnknown2:                12,
	GsPlayerHpMp:              15,
	GsGoldToInv8:              2,
	GsAddExp8:                 2,
	GsAddExp16:                3,
	GsAddExp32:                5,
	GsSetAttr8:                3,
	GsSetAttr16:               4,
	GsSetAttr32:               6,
	GsAttributeUpdate:         10,
	GsUpdateItemOSkill:        12,
	GsUpdateItemSkill:         12,
	GsSetSkill:                13,
	GsGameChat:                -1, // variable
	GsNPCInfo:                 40,
	GsQuestInfo:               103,
	GsGameQuestInfo:           97,
	GsNPCTransaction:          15,
	GsPlaySound:               8,
	GsUpdateItemStats:         -1, // variable
	GsUseStackableItem:        8,
	GsUnknown3:                13,
	GsClearCursor:             6,
	GsRelator1:                11,
	GsRelator2:                11,
	GsUnitSkillOnTarget:       16,
	GsUnitCastSkill:           17,
	GsMercForHire:             7,
	GsClearMercList:           1,
	GsQuestSpecial:            15,
	GsAssignObject:            14,
	GsPlayerQuestLog:          42,
	GsDarkness:                10,
	GsNPCEnchants:             14,
	GsUnknown4:                7,
	GsAssignPlayer:            26,
	GsEventMessages:           40,
	GsPlayerInGame:            -1, // variable
	GsPlayerLeft:              5,
	GsQuestItemState:          6,
	GsGameQuestsAvailable:     38,
	GsUnknown5:                5,
	GsTownPortalState:         7,
	GsCanGoToAct:              2,
	GsUnknown6:                7,
	GsWaypointMenu:            21,
	GsPlayerKillCount:         7,
	GsNPCMove:                 16,
	GsNPCMoveToTarget:         21,
	GsNPCState:                12,
	GsUnknown7:                12,
	GsNPCAction:               16,
	GsNPCAttack:               16,
	GsNPCStop:                 10,
	GsUnknown8:                32,
	GsPlayerCorpseAssign:      10,
	GsPlayerPartyInfo:         13,
	GsPlayerInProximity:       6,
	GsButtonActions:           2,
	GsTradeAccepted:           21,
	GsGoldInTrade:             6,
	GsPetAction:               13,
	GsAssignSkillHotkey:       8,
	GsUseScroll:               6,
	GsSetItemFlags:            18,
	GsCMNCOF:                  5,
	GsAllyPartyInfo:           10,
	GsAssignMerc:              20,
	GsPortalOwnership:         29,
	GsSpecialQuestEvent:       2,
	GsNPCWantsInteract:        6,
	GsPlayerRelationship:      6,
	GsRelationshipUpdate:      11,
	GsAssignPlayerToParty:     7,
	GsCorpseAssign:            10,
	GsPong:                    33,
	GsPartyAutomapInfo:        13,
	GsSetNPCGossip:            26,
	GsRemoveUnitDisplay:       6,
	GsUnknown9:                8,
	GsBaseSkillLevels:         -1, // variable
	GsLifeManaUpdate:          13,
	GsWalkVerify:              9,
	GsWeaponSwitch:            1,
	GsUpdateNPCUnknownField40: 7,
	GsCastSkillOnUnit:         16,
	GsCastSkillOnLocation:     17,
	GsMercReviveCost:          7,
	GsItemActionWorld:         -1, // variable
	GsItemActionOwned:         -1, // variable
	GsSetMercStat8:            7,
	GsSetMercStat16:           8,
	GsSetMercStat32:           10,
	GsAddMercExp8:             7,
	GsAddMercExp16:            8,
	GsSkillAuraStat:           24,
	GsNextBaalWaveClassID:     3,
	GsStateSkillMove:          8,
	GsUnknown10:               -1, // variable
	GsDelayedState:            7,
	GsSetState:                -1, // variable
	GsEndState:                7,
	GsAddUnit:                 -1, // variable
	GsNPCHeal:                 7,
	GsAssignNPC:               -1, // variable
	GsWardenRequest:           -1, // variable
	GsConnectionInfo:          2,
	GsTerminated:              1,
	GsUnknown11:               53,
	GsDownloadSave:            -1, // variable
	GsTimeout:                 5,
}

var gsServerPacketLengthFuncs = map[int]PacketLengthFunc{
	GsUnknown1: func(buffer PacketBuffer, offset, length int) (int, error) {
		count := buffer.Byte(offset + 3)
		return count*9 + 4, nil
	},
	GsGameChat: func(buffer PacketBuffer, offset, length int) (int, error) {
		nickEnd := buffer.IndexOf(0x00, offset+10)
		if nickEnd == -1 {
			return 0, fmt.Errorf("Unable to parse chat message packet")
		}
		msgEnd := buffer.IndexOf(0x00, nickEnd+1)
		if msgEnd == -1 {
			return 0, fmt.Errorf("Unable to parse chat message packet")
		}
		return msgEnd + 1 - offset, nil
	},
	GsUpdateItemStats: func(buffer PacketBuffer, offset, length int) (int, error) {
		// unknown so far :/
		return length - offset, nil
	},
	GsPlayerInGame: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Uint16(offset + 1), nil
	},
	GsBaseSkillLevels: func(buffer PacketBuffer, offset, length int) (int, error) {
		count := buffer.Byte(offset + 1)
		return count*3 + 6, nil
	},
	GsItemActionWorld: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Byte(offset + 2), nil
	},
	GsItemActionOwned: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Byte(offset + 2), nil
	},
	GsUnknown10: func(buffer PacketBuffer, offset, length int) (int, error) {
		// unknown so far :/
		return length - offset, nil
	},
	GsSetState: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Byte(offset + 6), nil
	},
	GsAddUnit: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Byte(offset + 6), nil
	},
	GsAssignNPC: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Byte(offset + 12), nil
	},
	GsWardenRequest: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Uint16(offset+1) + 1, nil
	},
	GsDownloadSave: func(buffer PacketBuffer, offset, length int) (int, error) {
		return buffer.Byte(offset+1) + 7, nil
	},
}

func GsServerPacketLength(buffer PacketBuffer, offset, length int) (int, error) {
	msgID := buffer.Byte(offset)
	plen, known := gsServerPacketLengths[msgID]
	if !known {
		return 0, fmt.Errorf("Unknown GS packet (S->C): 0x%x", msgID)
	}

	if plen < 0 {
		// variable length packet
		lengthFunc, exists := gsServerPacketLengthFuncs[msgID]
		if !exists {
			return 0, fmt.Errorf("No length function for S->C packet 0x%X", msgID)
		}
		return lengthFunc(buffer, offset, length)
	}

	return plen, nil
}

var gsServerPacketNames = map[int]string{
	GsGameLoading:             "GameLoading",
	GsGameFlags:               "GameFlags",
	GsLoadSuccess:             "LoadSuccess",
	GsLoadAct:                 "LoadAct",
	GsLoadComplete:            "LoadComplete",
	GsUnloadComplete:          "UnloadComplete",
	GsGameExitSuccess:         "GameExitSuccess",
	GsMapReveal:               "MapReveal",
	GsMapHide:                 "MapHide",
	GsAssignLevelWarp:         "AssignLevelWarp",
	GsRemoveObject:            "RemoveObject",
	GsHandshake:               "Handshake",
	GsNPCHit:                  "NPCHit",
	GsPlayerStop:              "PlayerStop",
	GsObjectState:             "ObjectState",
	GsPlayerMove:              "PlayerMove",
	GsPlayerToTarget:          "PlayerMoveToTarget",
	GsReportKill:              "ReportKill",
	GsReassignPlayer:          "ReassignPlayer",
	GsPlayerHpMp:              "PlayerHealthMana",
	GsGoldToInv8:              "GoldToInv8",
	GsAddExp8:                 "AddExp8",
	GsAddExp16:                "AddExp16",
	GsAddExp32:                "AddExp32",
	GsSetAttr8:                "SetAttr8",
	GsSetAttr16:               "SetAttr16",
	GsSetAttr32:               "SetAttr32",
	GsAttributeUpdate:         "AttributeUpdate",
	GsUpdateItemOSkill:        "UpdateItemOSKill",
	GsUpdateItemSkill:         "UpdateItemSkill",
	GsSetSkill:                "SetSkill",
	GsGameChat:                "GameChat",
	GsNPCInfo:                 "NPCInfo",
	GsQuestInfo:               "QuestInfo",
	GsGameQuestInfo:           "GameQuestInfo",
	GsNPCTransaction:          "NPCTransaction",
	GsPlaySound:               "PlaySound",
	GsUpdateItemStats:         "UpdateItemStats",
	GsUseStackableItem:        "UseStackableItem",
	GsClearCursor:             "ClearCursor",
	GsRelator1:                "Relator1",
	GsRelator2:                "Relator2",
	GsUnitSkillOnTarget:       "UnitSkillOnTarget",
	GsUnitCastSkill:           "UnitCastSkill",
	GsMercForHire:             "MercForHire",
	GsClearMercList:           "ClearMercList",
	GsQuestSpecial:            "QuestSpecial",
	GsAssignObject:            "AssignObject",
	GsPlayerQuestLog:          "PlayerQuestLog",
	GsDarkness:                "Darkness",
	GsNPCEnchants:             "NPCEnchants",
	GsAssignPlayer:            "AssignPlayer",
	GsEventMessages:           "EventMessages",
	GsPlayerInGame:            "PlayerInGame",
	GsPlayerLeft:              "PlayerLeft",
	GsQuestItemState:          "QuestItemState",
	GsGameQuestsAvailable:     "GameQuestsAvailable",
	GsTownPortalState:         "TownPortalState",
	GsCanGoToAct:              "CanGoToAct",
	GsWaypointMenu:            "WaypointMenu",
	GsPlayerKillCount:         "PlayerKillCount",
	GsNPCMove:                 "NPCMove",
	GsNPCMoveToTarget:         "NPCMoveToTarget",
	GsNPCState:                "NPCState",
	GsUnknown7:                "Unknown",
	GsNPCAction:               "NPCAction",
	GsNPCAttack:               "NPCAttack",
	GsNPCStop:                 "NPCStop",
	GsPlayerCorpseAssign:      "PlayerCorpseAssign",
	GsPlayerPartyInfo:         "PlayerPartyInfo",
	GsPlayerInProximity:       "PlayerInProximity",
	GsButtonActions:           "ButtonActions",
	GsTradeAccepted:           "TradeAccepted",
	GsGoldInTrade:             "GoldInTrade",
	GsPetAction:               "PetAction",
	GsAssignSkillHotkey:       "AssignSkillHotkey",
	GsUseScroll:               "UseScroll",
	GsSetItemFlags:            "SetItemFlags",
	GsCMNCOF:                  "CMNCOF",
	GsAllyPartyInfo:           "AllyPartyInfo",
	GsAssignMerc:              "AssignMerc",
	GsPortalOwnership:         "PortalOwnership",
	GsSpecialQuestEvent:       "SpecialQuestEvent",
	GsNPCWantsInteract:        "NPCWantsToInteract",
	GsPlayerRelationship:      "PlayerRelationship",
	GsRelationshipUpdate:      "RelationshipUpdate",
	GsAssignPlayerToParty:     "AssignPlayerToParty",
	GsCorpseAssign:            "CorpseAssign",
	GsPong:                    "Pong",
	GsPartyAutomapInfo:        "PartyAutomapInfo",
	GsSetNPCGossip:            "SetNPCGossip",
	GsRemoveUnitDisplay:       "RemoveUnitDisplay",
	GsBaseSkillLevels:         "BaseSkillLevels",
	GsLifeManaUpdate:          "LifeManaUpdate",
	GsWalkVerify:              "WalkVerify",
	GsWeaponSwitch:            "WeaponSwitch",
	GsUpdateNPCUnknownField40: "UpdateNPCUnknownField40",
	GsCastSkillOnUnit:         "CastSkillOnUnit",
	GsCastSkillOnLocation:     "CastSkillOnLocation",
	GsMercReviveCost:          "MercReviveCost",
	GsItemActionWorld:         "ItemActionWorld",
	GsItemActionOwned:         "ItemActionOwned",
	GsSetMercStat8:            "SetMercStat8",
	GsSetMercStat16:           "SetMercStat16",
	GsSetMercStat32:           "SetMercStat32",
	GsAddMercExp8:             "AddMercExp8",
	GsAddMercExp16:            "AddMercExp16",
	GsSkillAuraStat:           "SkillAuraStat",
	GsNextBaalWaveClassID:     "NextBaalWaveNPCClass",
	GsStateSkillMove:          "StateSkillMove",
	GsDelayedState:            "DelayedState",
	GsSetState:                "SetState",
	GsEndState:                "EndState",
	GsAddUnit:                 "AddUnit",
	GsNPCHeal:                 "NPCHeal",
	GsAssignNPC:               "AssignNPC",
	GsWardenRequest:           "WardenRequest",
	GsConnectionInfo:          "ConnectionInfo",
	GsTerminated:              "Terminated",
	GsDownloadSave:            "DownloadSave",
	GsTimeout:                 "Timeout",
}

func GsServerPacketName(packet Packet) string {
	msgID := int(packet.GsMsgID())
	name, exists := gsServerPacketNames[msgID]
	if !exists {
		return fmt.Sprintf("Unknown (0x%X)", byte(msgID))
	}
	return name
}
