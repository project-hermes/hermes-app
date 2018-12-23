package macdive

import "encoding/xml"

type MacDive struct {
	XMLName xml.Name `xml:"dives"`
	Units   string   `xml:"units"`
	Schema  string   `xml:"schema"`
	Dives   []Dive   `xml:"dive"`
}

type Dive struct {
	Date                 string  `xml:"date"`
	Identifier           string  `xml:"identifier"`
	DiveNumber           uint    `xml:"diveNumber"`
	Rating               int     `xml:"rating"`
	RepetitiveDive       uint    `xml:"repetitiveDive"`
	Diver                string  `xml:"diver"`
	ComputerName         string  `xml:"computer"`
	ComputerSerialNumber string  `xml:"serial"`
	MaxDepth             float32 `xml:"maxDepth"`
	AverageDepth         float32 `xml:"averageDepth"`
	CNS                  float32 `xml:"cns"`
	DecoModel            string  `xml:"decoModel"`
	Duration             uint    `xml:"duration"`
	GasModel             string  `xml:"gasModel"`
	SurfaceInterval      uint    `xml:"surfaceInterval"`
	SampleInterval       uint    `xml:"sampleInterval"`
	TempAir              float32 `xml:"tempAir"`
	TempHigh             float32 `xml:"tempHigh"`
	TempLow              float32 `xml:"tempLow"`
	Visibility           string  `xml:"visibility"`
	Weight               float32 `xml:"weight"`
	Notes                string  `xml:"notes"`
	DiveMaster           string  `xml:"diveMaster"`
	DiveOperator         string  `xml:"diveOperator"`
	Skipper              string  `xml:"skipper"`
	Boat                 string  `xml:"boat"`
	Weather              string  `xml:"weather"`
	Current              string  `xml:"current"`
	SurfaceConditions    string  `xml:"surfaceConditions"`
	EntryType            string  `xml:"entryType"`

	Site    Site   `xml:"site"`
	Tags    string `xml:"tags"`
	Types   string `xml:"types"`
	Buddies string `xml:"buddies"`

	Gears  []Item   `xml:"gear>item"`
	Gases  []Gas    `xml:"gases>gas"`
	Sample []Sample `xml:"samples>sample"`
	Events string   `xml:"events"`
}

type Item struct {
	Type         string `xml:"type"`
	Manufacturer string `xml:"manufacturer"`
	Name         string `xml:"name"`
	Serial       string `xml:"serial"`
}

type Gas struct {
	PressureStart   float32 `xml:"pressureStart"`
	PressureEnd     float32 `xml:"pressureEnd"`
	Oxygen          uint    `xml:"oxygen"`
	Helium          uint    `xml:"helium"`
	Double          uint    `xml:"double"`
	TankSize        uint    `xml:"tankSize"`
	WorkingPressure uint    `xml:"workingPressure"`
	SupplyType      string  `xml:"supplyType"`
	Duration        uint    `xml:"duration"`
	TankName        string  `xml:"tankName"`
}

type Sample struct {
	Time        float32 `xml:"time"`
	Depth       float32 `xml:"depth"`
	Pressure    float32 `xml:"pressure"`
	Temperature float32 `xml:"temperature"`
	Ppo2        float32 `xml:"ppo2"`
	Ndt         float32 `xml:"ndt"`
}

type Site struct {
	Country     string  `xml:"country"`
	Location    string  `xml:"location"`
	Name        string  `xml:"name"`
	BodyOfWater string  `xml:"bodyOfWater"`
	WaterType   string  `xml:"waterType"`
	Difficulty  string  `xml:"difficulty"`
	Altitude    float32 `xml:"altitude"`
	Lat         float32 `xml:"lat"`
	Lon         float32 `xml:"lon"`
}
