package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	vomni "vk/omnibus"

	"github.com/gorilla/mux"
)

func pagePointListActionSubtype(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	todo := strings.ToUpper(vars["todo"])
	point := vars["point"]
	subtype := vars["subtype"]

	//	var data interface{}

	fmt.Println("SVIRIDOVS")

	/*
		tmplStr := "pointcfg"
		data = pointCfg(point)

		refl := reflect.ValueOf(data)

		zType := refl.FieldByName("Type")

		switch zType.Int() {
		case vomni.PointTypeRelayOnOffInterval:
			tmplStr = "cfgrelayonoffinterval"
		default:
			tmplStr = "pointcfg"
		}

		err = tmpls.ExecuteTemplate(w, tmplStr, point)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	*/
	switch todo {
	case "CFG":

		err := error(nil)
		tmplStr := ""

		switch subtype {

		case strconv.Itoa(vomni.CfgTypeRelayInterval):
			tmplStr = "cfgrelayinterval"
		default:
			err := fmt.Errorf("Don't have code to handle configuration %q", point)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//rescanPoint(point)
		//tmplStr := "pointcfg"

		//cfg, _ := strconv.Atoi(subtype)

		fmt.Printf("Kods %q subtype %q\n", strconv.Itoa(vomni.CfgTypeRelayInterval), point)

		//data := pointData(point)

		err = tmpls.ExecuteTemplate(w, tmplStr, point)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	case "LOADCFG", "SAVECFG":
		/*
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err.Error())
			}
			err = json.Unmarshal(body, &data)
			if err != nil {
				panic(err.Error())
			}
		*/
	case "FREEZE", "UNFREEZE", "LOADDEFAULTCFG", "LOADSAVEDCFG":
	default:
		log.Fatal(fmt.Sprintf("===> Don't know what to do with %q (point %q with subtype %q )", todo, point, subtype))
	}

	responseOK(w)
	//	xrun.ReceivedWebMsg(point, todo, data)
}
