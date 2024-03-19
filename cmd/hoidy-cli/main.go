package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/achiku/hoidy"
)

func main() {
	pt := flag.String("file", "", "XML file path")
	flag.Parse()

	if *pt == "" {
		log.Fatal("file parameter is required")
	}

	cms, err := hoidy.ParseXML(*pt)
	if err != nil {
		log.Fatal(err)
	}
	for _, cm := range cms.Corporations {
		fmt.Printf(
			"insert into corporation (corporate_number, name, furigana, prefecture_name, city_name) values ('%s', '%s', '%s', '%s', '%s');\n",
			cm.CorporateNumber, cm.Name, cm.Furigana, cm.PrefectureName, cm.CityName,
		)
	}
}
