package subsurface

import "encoding/xml"

type Subsurface struct {
	XMLName   xml.Name       `xml:"divelog"`
	Settings  DiveComputerId `xml:"settings>divecomputerid"`
	DiveSites []Site         `xml:"divesites>site"`
	Dives     []Dive         `xml:"dives>dive"`
}

type DiveComputerId struct {
	Model        string `xml:"model,attr"`
	DeviceId     string `xml:"deviceid,attr"`
	SerialNumber string `xml:"serial,attr"`
}

type Site struct {
	UUID string `xml:"uuid,attr"`
	Name string `xml:"name,attr"`
	GPS  string `xml:"gps,attr"`
}

type Dive struct {
	Number   uint   `xml:"number,attr"`
	Date     string `xml:"date,attr"`
	Time     string `xml:"time,attr"`
	Duration string `xml:"duration,attr"`

	Cylinder     Cylinder     `xml:"cylinder"`
	DiveComputer DiveComputer `xml:"divecomputer"`
}

type Cylinder struct {
	Size         string `xml:"size,attr"`
	WorkPressure string `xml:"workpressure,attr"`
	Description  string `xml:"description,attr"`
}

type DiveComputer struct {
	Model    string `xml:"model,attr"`
	DeviceId string `xml:"deviceid,attr"`

	Depth       Depth       `xml:"depth"`
	Temperature Temperature `xml:"temperature"`
	Events      []Event     `xml:"event"`
	Samples     []Sample    `xml:"sample"`
}

type Depth struct {
	Max  string `xml:"max,attr"`
	Mean string `xml:"mean,attr"`
}

type Temperature struct {
	Air   string `xml:"air,attr"`
	Water string `xml:"water,attr"`
}

type Event struct {
	Time     string `xml:"time,attr"`
	Type     string `xml:"type,attr"`
	Value    string `xml:"value,attr"`
	Name     string `xml:"name,attr"`
	Cylinder string `xml:"cylinder,attr"`
}

type Sample struct {
	Time  string `xml:"time,attr"`
	Depth string `xml:"depth,attr"`
	Temp  string `xml:"temp,attr"`
}
