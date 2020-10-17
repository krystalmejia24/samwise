package event

//ID3 holds configuration values for constructing an ID3 tag
type ID3 struct {
	OwnerID []byte
	Payload []byte
}

const (
	nilByte    byte = 0 //used as default for unused flags
	TagVersion byte = 4 //v2.4.0
)

func NewID3(owner string, payload string) *ID3 {
	return &ID3{
		OwnerID: []byte(owner),
		Payload: []byte(payload),
	}
}

func (i *ID3) CreateTag(owner string, payload string) []byte {
	id3 := append(getHeader(), getPrivFrame([]byte(owner), []byte(payload))...)

	return id3
}

func (i *ID3) GetTag() []byte {
	id3 := append(getHeader(), getPrivFrame(i.OwnerID, i.Payload)...)

	return id3
}

func getHeader() []byte {
	tempSize := []byte("D")
	header := append([]byte("ID3"), TagVersion, nilByte, nilByte, nilByte, nilByte, 8)
	return append(header, tempSize...)
}

func getPrivFrame(owner []byte, payload []byte) []byte {
	var data []byte
	data = append(owner, nilByte)
	data = append(data, payload...)

	header := append([]byte("PRIV"), nilByte, nilByte, nilByte)
	header = append(header, []byte(":")...)
	header = append(header, nilByte, nilByte)

	return append(header, data...)
}
