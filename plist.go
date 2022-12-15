package plist

import (
	"go.k6.io/k6/js/modules"
	"howett.net/plist"
)

const (
	// Used by Decoder to represent an invalid property list.
	InvalidFormat int = plist.InvalidFormat

	// Used to indicate total abandon with regards to Encoder's output format.
	AutomaticFormat = plist.AutomaticFormat

	XMLFormat      = plist.XMLFormat
	BinaryFormat   = plist.BinaryFormat
	OpenStepFormat = plist.OpenStepFormat
	GNUStepFormat  = plist.GNUStepFormat
)

type PList struct {
	vu modules.VU
}

func (p *PList) Parse(input string) (data any, format int, err error) {
	format, err = plist.Unmarshal([]byte(input), &data)
	return
}

func (p *PList) Serialize(data any, format int) (string, error) {
	out, err := plist.Marshal(data, format)
	return string(out), err
}

func (p *PList) SerializeIndent(data any, format int, indent string) (string, error) {
	out, err := plist.MarshalIndent(data, format, indent)
	return string(out), err
}

func (p *PList) GetFormatName(format int) string {
	if name, ok := plist.FormatNames[format]; ok {
		return name
	}
	return plist.FormatNames[plist.InvalidFormat]
}

func newPlist(vu modules.VU) *PList {
	return &PList{
		vu: vu,
	}
}
