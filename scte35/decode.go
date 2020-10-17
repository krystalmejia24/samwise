package scte35

import (
	"bytes"
	"fmt"

	"github.com/krystalmejia24/bitio"
)

const (
	BITFOUND, BITNOTFOUND = uint64(1), uint64(0)
)

func newReader(b []byte) (func(n byte) uint64, *bytes.Buffer) {
	buffer := bytes.NewBuffer(b)
	r := bitio.NewReader(buffer)
	return func(n byte) uint64 {
		bits, _ := r.ReadBits(n)
		return bits
	}, buffer
}

// NewScte35 returns a parsed scte35 marker
func NewScte35(signal []byte) (Scte35, error) {
	spliceInfo, e := parse(signal)
	if e != nil {
		return Scte35{}, e
	}

	return Scte35{
		SpliceInfo: spliceInfo,
	}, nil
}

func parse(signal []byte) (spliceInfo, error) {
	read, buffer := newReader(signal)
	fmt.Printf("bytes1\n%v\n", buffer.Bytes())
	tableID := read(8)
	sectionSyntaxIndicator := read(1)
	privateIndicator := read(1)
	reserved := read(2)
	sectionLength := read(12)
	protocolVersion := read(8)
	encryptedPacket := read(1)
	encryptionAlgorithm := read(6)
	ptsAdjustment := read(33)
	cwIndex := read(8)
	tier := read(12)
	spliceCommandLength := read(12)
	spliceCommandType := read(8)

	spliceEvent := spliceEvent{}
	switch SpliceCommandType(spliceCommandType) {
	case TimeSignal:
		spliceEvent.spliceTime = parseSpliceTime(read)
	case SpliceInsert:
		spliceEvent.spliceInsert, spliceEvent.spliceTime = parseSpliceInsert(read)
	}

	descriptorLoopLength := read(16)

	//TODO LOOP BASED ON LENGTH
	//length := descriptorLoopLength
	spliceDescriptors := []spliceDescriptor{}
	spliceDescriptors = append(spliceDescriptors, parseSpliceDescriptor(read))

	//spliceDescriptors := parseSpliceDescriptor(read)
	crc32 := read(32)

	return spliceInfo{
		tableID:                tableID,
		sectionSyntaxIndicator: sectionSyntaxIndicator,
		privateIndicator:       privateIndicator,
		reserved:               reserved,
		sectionLength:          sectionLength,
		protocolVersion:        protocolVersion,
		encryptedPacket:        encryptedPacket,
		encryptionAlgorithm:    encryptionAlgorithm,
		ptsAdjustment:          ptsAdjustment,
		cwIndex:                cwIndex,
		tier:                   tier,
		spliceCommandLength:    spliceCommandLength,
		spliceCommandType:      spliceCommandType,
		spliceEvent:            spliceEvent,
		descriptorLoopLength:   descriptorLoopLength,
		spliceDescriptors:      spliceDescriptors,
		crc32:                  crc32,
	}, nil
}

func parseSpliceTime(read func(n byte) uint64) spliceTime {
	s := spliceTime{}

	s.timeSpecifiedFlag = read(1)

	if s.timeSpecifiedFlag == BITFOUND {
		s.reserved = read(6)
		s.ptsTime = read(33)
	} else {
		s.reserved = read(7)
	}

	return s
}

func parseSpliceInsert(read func(n byte) uint64) (spliceInsert, spliceTime) {
	s, st := spliceInsert{}, spliceTime{}
	s.spliceEventID = read(32)
	s.spliceEventCancelIndicator = read(1)
	s.reserved1 = read(7)

	if s.spliceEventCancelIndicator == BITNOTFOUND {
		s.outOfNetworkIndicator = read(1)
		s.programSpliceFlag = read(1)
		s.durationFlag = read(1)
		s.spliceImmediateFlag = read(1)
		s.reserved2 = read(4)

		if s.programSpliceFlag == BITFOUND && s.spliceImmediateFlag == BITNOTFOUND {
			fmt.Println("found splice time in splice insert")
			st = parseSpliceTime(read)
		}

		if s.programSpliceFlag == BITNOTFOUND {
			s.componentCount = read(8)
			for i := 0; i < int(s.componentCount); i++ {
				fmt.Println("found splice time in splice insert as component")
				s.componentTag = read(8)
				s.component = append(s.component, parseSpliceTime(read))
			}
		}

		s.uniqueProgramID = read(16)
		s.availNum = read(8)
		s.availsExpected = read(8)
	}

	return s, st
}

func parseSpliceDescriptor(read func(n byte) uint64) spliceDescriptor {
	s := spliceDescriptor{}
	s.spliceDescriptorTag = read(8)
	s.descriptorLength = read(8)
	s.identifier = read(32)

	switch SpliceDescriptorTag(s.spliceDescriptorTag) {
	case AvailDescriptor:
		fmt.Println("AvailNotYetSupported")
	case DTMFDescriptor:
		fmt.Println("DTMFNotYetSupported")
	case SegmentationDescriptor:
		fmt.Println("SegmentationDescriptor")
		s.segmentationDescriptor = parseSegmentationDescriptor(read)
	case TimeDescriptor:
		fmt.Println("TimeNotYetSupported")
	case AudioDescriptor:
		fmt.Println("AudioNotYetSupported")
	default:
		i := SpliceDescriptorTag(s.spliceDescriptorTag)
		if ReservedTagMin <= i && i >= ReservedTagMax {
			fmt.Println("ReservedNotYetSupported")
		} else {
			fmt.Println("SpliceDescriptorNotFound")
		}
	}

	return s
}

func parseSegmentationDescriptor(read func(n byte) uint64) segmentationDescriptor {
	s := segmentationDescriptor{}

	s.segmentationEventID = read(32)
	s.segmentationEventCancelIndicator = read(1)
	s.reserved1 = read(7)

	if s.segmentationEventCancelIndicator == BITNOTFOUND {
		s.programSegmentationFlag = read(1)
		s.segmentationDurationFlag = read(1)
		s.deliveryNotRestrictedFlag = read(1)

		if s.deliveryNotRestrictedFlag == BITNOTFOUND {
			s.webDeliveryAllowedFlag = read(1)
			s.noRegionalBlackoutFlag = read(1)
			s.archiveAllowedFlag = read(1)
			s.deviceRestirictionsFlags = read(2)
		} else {
			s.reserved2 = read(5) //TODO test RESERVED2 SET IF DELIVERY NOT RESTRICTED FLAGS IS TRUE
		}

		if s.programSegmentationFlag == BITNOTFOUND {
			s.componentCount = read(8)
			for i := uint64(0); i < s.componentCount; i++ {
				c := descriptorComponent{}
				c.tag = read(8)
				c.reserved = read(7)
				c.ptsOffset = read(33)
				s.component = append(s.component, c)
			}

		}

		if s.segmentationDurationFlag == BITFOUND {
			s.segmentationDuration = read(40)
		}

		s.segmentationUpidType = read(8)
		s.segmentationUpidLength = read(8)
		s.segmentationUpid = read(uint8(s.segmentationUpidLength)) // value is determined by segmentationUpidType and segmentationUpidLength
		s.segmentationTypeID = read(8)                             // commands
		s.segmentNum = read(8)
		s.segmentsExpected = read(8)
		if s.segmentationTypeID == 0x34 || s.segmentationTypeID == 0x36 {
			s.subSegmentNum = read(8)
			s.subSegmentsExpected = read(8)
		}

	}

	return s
}
