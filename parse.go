package hoidy

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// Corporation company info
// https://www.houjin-bangou.nta.go.jp/download/
type Corporation struct {
	SequenceNumber  string `xml:"sequenceNumber"`
	CorporateNumber string `xml:"corporateNumber"`
	Process         string `xml:"process"`
	// UpdateDate      time.Time `xml:"updateDate"`
	UpdateDate string `xml:"updateDate"`
	// ChangeDate      time.Time `xml:"changeDate"`
	ChangeDate     string `xml:"changeDate"`
	Name           string `xml:"name"`
	PrefectureName string `xml:"prefectureName"`
	CityName       string `xml:"cityName"`
	StreetNumber   string `xml:"streetNumber"`
	Furigana       string `xml:"furigana"`
}

type Corporations struct {
	XMLName      xml.Name      `xml:"corporations"`
	Corporations []Corporation `xml:"corporation"`
}

// ParseXML houjin ID XML parser
func ParseXML(pt string) (*Corporations, error) {
	f, err := os.Open(pt)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var cms Corporations
	if err := xml.Unmarshal(c, &cms); err != nil {
		return nil, err
	}
	return &cms, nil
}
