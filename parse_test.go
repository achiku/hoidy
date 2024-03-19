package hoidy

import "testing"

func TestParseXML(t *testing.T) {
	cms, err := ParseXML("./data/company_test.xml")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cms.Corporations {
		t.Logf("%v", c)
	}
}
