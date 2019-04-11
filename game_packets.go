package d2prox

const GsWalkToLocation = 0x01
const GsWalkToEntity = 0x02
const GsRunToLocation = 0x03
const GsRunToEntity = 0x04
const GsPickupItem = 0x16
const GsDropItem = 0x17
const GsNPCInit = 0x2F
const GsPing = 0x6d

var gsClientPacketLengths = map[int]int{
	GsWalkToLocation: 5,  // 0x01
	GsWalkToEntity:   9,  // 0x02
	GsRunToLocation:  5,  // 0x03
	GsRunToEntity:    9,  // 0x04
	GsPickupItem:     13, // 0x16
	GsDropItem:       13, // 0x16
	GsNPCInit:        9,  // 0x2F
	GsPing:           13, // 0x65
}

func GsClientPacketLength(buffer PacketBuffer, offset, length int) (int, error) {
	msgID := buffer.Byte(offset)
	plen, known := gsClientPacketLengths[msgID]
	if known {
		return plen, nil
	}
	return length - offset, nil
}

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

const GsCreateClientPlayer = 0x59
const GsGameClose = 0xb0

var gsServerPacketLengths = map[int]int{
	GsGameLoading:       1,
	GsGameFlags:         8,
	GsLoadSuccess:       2,
	GsLoadAct:           12,
	GsLoadComplete:      1,
	GsUnloadComplete:    1,
	GsGameExitSuccess:   1,
	GsMapReveal:         6,
	GsMapHide:           6,
	GsAssignLevelWarp:   12,
	GsRemoveObject:      6,
	GsHandshake:         6,
	GsNPCHit:            9,
	GsPlayerStop:        13,
	GsObjectState:       12,
	GsPlayerMove:        16,
	GsPlayerToTarget:    16,
	GsReportKill:        8,
	GsReassignPlayer:    11,
	GsUnknown1:          -1, // variable
	GsUnknown2:          -1, // variable
	GsPlayerHpMp:        15,
	GsGoldToInv8:        2,
	GsAddExp8:           2,
	GsAddExp16:          3,
	GsAddExp32:          5,
	GsSetAttr8:          3,
	GsSetAttr16:         4,
	GsSetAttr32:         6,
	GsAttributeUpdate:   10,
	GsUpdateItemOSkill:  12,
	GsUpdateItemSkill:   12,
	GsSetSkill:          13,
	GsGameChat:          -1, // variable
	GsNPCInfo:           40,
	GsQuestInfo:         103,
	GsGameQuestInfo:     97,
	GsNPCTransaction:    15,
	GsPlaySound:         8,
	GsUpdateItemStats:   -1, // variable
	GsUseStackableItem:  8,
	GsUnknown3:          13,
	GsClearCursor:       6,
	GsRelator1:          11,
	GsRelator2:          11,
	GsUnitSkillOnTarget: 16,
	GsUnitCastSkill:     17,
	GsMercForHire:       7,
	GsClearMercList:     1,
	GsQuestSpecial:      15,
	GsAssignObject:      14,
	GsPlayerQuestLog:    42,
	GsDarkness:          10,

	GsCreateClientPlayer: 0x1A,
	GsGameClose:          1,
}

func GsServerPacketLength(buffer PacketBuffer, offset, length int) (int, error) {
	msgID := buffer.Byte(offset)
	plen, known := gsServerPacketLengths[msgID]
	if known {
		return plen, nil
	}
	return length - offset, nil
}
