package runrelayinterval

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	vcfg "vk/pointconfig"
)

func webInterface2Struct(data interface{}) (back vcfg.RunRelIntervalStruct) {
	// WEB struct
	web := webPointStruct{}
	for part, v := range data.(map[string]interface{}) { // list add configuration parts

		fmt.Println("PART ", part)

		d := webPointArr{}                     // array for the configuration part records
		for _, v1 := range v.([]interface{}) { // fill part record array
			rec := webPoint{} // storage for a record data
			for k2, v2 := range v1.(map[string]interface{}) {

				fmt.Printf("FIELD %q VALUE %v TYPE %T\n", k2, v2, v2)

				switch strings.ToUpper(k2) {
				case "GPIO":
					rec.Gpio = v2.(string)
				case "STATE":
					rec.State = v2.(string)
				case "SECONDS": //"SECONDS":
					str := v2.(string)
					// remove nanoseconds
					nn, _ := strconv.Atoi(str)

					rec.Seconds = strconv.Itoa(nn / 1000000000) //v2.(string)
				default:
					log.Fatal(fmt.Sprintf("Unknow WEB interface record field \"%s\"", k2))
				}
			}
			d = append(d, rec)
		}

		switch strings.ToUpper(part) {
		case "START":
			web.Start = d
		case "BASE":
			web.Base = d
		case "FINISH":
			web.Finish = d
		default:
			log.Fatal(fmt.Sprintf("Unknow WEB interface part \"%s\"", part))
		}
	}

	// from the WEB structure to the regular one
	back = vcfg.RunRelIntervalStruct{}
	back.Start = web.Start.webArray2Regular()
	back.Base = web.Base.webArray2Regular()
	back.Finish = web.Finish.webArray2Regular()

	return
}

func (d webPointArr) webArray2Regular() (newStr vcfg.RunRelIntervalArray) {

	newStr = vcfg.RunRelIntervalArray{}

	for _, v := range d {
		newR := vcfg.RunRelInterval{}
		newR.Gpio, _ = strconv.Atoi(v.Gpio)
		newR.State, _ = strconv.Atoi(v.State)

		t, _ := time.ParseDuration(v.Seconds + "s")
		newR.Seconds = t.Round(time.Second)

		newStr = append(newStr, newR)
	}

	return
}
