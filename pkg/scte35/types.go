package scte35

// Scte35 marker
type Scte35 struct {
	SpliceInfo spliceInfo
}

type spliceInfo struct {
	tableID                uint64
	sectionSyntaxIndicator uint64
	privateIndicator       uint64
	reserved               uint64
	sectionLength          uint64
	protocolVersion        uint64
	encryptedPacket        uint64
	encryptionAlgorithm    uint64
	ptsAdjustment          uint64
	cwIndex                uint64
	tier                   uint64
	spliceCommandLength    uint64
	spliceCommandType      uint64
	spliceEvent            spliceEvent
	descriptorLoopLength   uint64
	spliceDescriptors      []spliceDescriptor
	crc32                  uint64
}

// SpliceCommandType specifies command in the marker used to splice
type SpliceCommandType uint64

const (
	// SpliceNull is a Null Splice command type
	SpliceNull SpliceCommandType = 0x00
	// SpliceSchedule is a splice schedule command type
	SpliceSchedule SpliceCommandType = 0x04
	// SpliceInsert is a splice insert command type
	SpliceInsert SpliceCommandType = 0x05
	// TimeSignal is a splice signal command type
	TimeSignal SpliceCommandType = 0x06
	// BandwidthReservation is a command type that represents a reservation of bandwidth
	BandwidthReservation SpliceCommandType = 0x07
	// PrivateCommand is a command type that represents private command data
	PrivateCommand SpliceCommandType = 0xFF
)

type spliceEvent struct {
	spliceTime   spliceTime
	spliceInsert spliceInsert
}

type spliceTime struct {
	timeSpecifiedFlag uint64
	reserved          uint64
	ptsTime           uint64
}

type spliceInsert struct {
	spliceEventID              uint64
	spliceEventCancelIndicator uint64
	reserved1                  uint64
	outOfNetworkIndicator      uint64
	programSpliceFlag          uint64
	durationFlag               uint64
	spliceImmediateFlag        uint64
	reserved2                  uint64
	componentCount             uint64
	componentTag               uint64
	component                  []spliceTime
	uniqueProgramID            uint64
	availNum                   uint64
	availsExpected             uint64
}

type spliceDescriptor struct {
	spliceDescriptorTag    uint64 // 0x02 = segmentation descriptor descriptor
	descriptorLength       uint64
	identifier             uint64
	privateByte            uint64
	segmentationDescriptor segmentationDescriptor
}

// SpliceDescriptorTag specifies descriptor carried in the marker
type SpliceDescriptorTag uint64

const (
	// AvailDescriptor is an implementation of a splice_descriptor. It provides an optional extension to the
	// splice_insert() command that allows an authorization identifier to be sent for an avail
	AvailDescriptor SpliceDescriptorTag = 0x00
	// DTMFDescriptor provides an optional extension to the splice_insert() command that allows a receiver device to
	// generate a legacy analog DTMF sequence
	DTMFDescriptor SpliceDescriptorTag = 0x01
	// SegmentationDescriptor provides an optional extension to the time_signal() and splice_insert() commands that
	// allows for segmentation messages to be sent in a time/video accurate method
	SegmentationDescriptor SpliceDescriptorTag = 0x02
	// TimeDescriptor provides an optional extension to the splice_insert(), splice_null() and time_signal() commands
	//that allows a programmer’s wall clock time to be sent to a client
	TimeDescriptor SpliceDescriptorTag = 0x03
	// AudioDescriptor should be used when programmers and/or MVPDs do not support dynamic signaling (e.g., signaling of
	// audio language changes) and with legacy audio formats that do not support dynamic signaling
	AudioDescriptor SpliceDescriptorTag = 0x04
	// ReservedTagMin is the lower bound of the reserved tag
	ReservedTagMin SpliceDescriptorTag = 0x05
	// ReservedTagMax is the upper bound of the reserved tag
	ReservedTagMax SpliceDescriptorTag = 0xFF
)

type availDescriptor struct {
	spliceDescriptorTag uint64
	descriptorLength    uint64
	identifier          uint64
	providerAvailID     uint64
}

type dtmfDescriptor struct {
	spliceDescriptorTag uint64
	descriptorLength    uint64
	identifier          uint64
	preroll             uint64
	dtmfCount           uint64
	reserved            uint64
	dtmfChar            uint64
}

type timeDescriptor struct {
	spliceDescriptorTag uint64
	descriptorLength    uint64
	identifier          uint64
	TAIseconds          uint64
	TAIns               uint64
	utcOffset           uint64
}

type audioDescriptor struct {
	spliceDescriptorTag uint64
	descriptorLength    uint64
	identifier          uint64
	audioCount          uint64
	reserved            uint64
	componentTag        uint64
	ISOcode             uint64
	bitStreamMode       uint64
	numChannels         uint64
	fullSrvcAudio       uint64
}

type segmentationDescriptor struct {
	segmentationEventID              uint64
	segmentationEventCancelIndicator uint64
	reserved1                        uint64
	programSegmentationFlag          uint64
	segmentationDurationFlag         uint64
	deliveryNotRestrictedFlag        uint64
	webDeliveryAllowedFlag           uint64
	noRegionalBlackoutFlag           uint64
	archiveAllowedFlag               uint64
	deviceRestirictionsFlags         uint64
	reserved2                        uint64
	componentCount                   uint64
	component                        []descriptorComponent
	ptsOffset                        uint64
	segmentationDuration             uint64
	segmentationUpidType             uint64
	segmentationUpidLength           uint64
	segmentationUpid                 uint64 // value is determined by segmentationUpidType and segmentationUpidLength
	segmentationTypeID               uint64 // commands
	segmentNum                       uint64
	segmentsExpected                 uint64
	subSegmentNum                    uint64
	subSegmentsExpected              uint64
}

type descriptorComponent struct {
	tag       uint64
	ptsOffset uint64
	reserved  uint64
}

type breakDuration struct {
	autoReturn uint64
	reserved   uint64
	duration   uint64
}
