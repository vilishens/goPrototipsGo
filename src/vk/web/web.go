package web

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	vparams "vk/params"
)

func GoWeb(chGoOn chan bool, chDone chan int, chErr chan error) {

	/* šitas vajadzīgs, ja būs jāizmanto iepriekšējo autentikāciju
	os.Mkdir(backendfile, 0755)
	defer os.Remove(backendfile)
	*/

	setTmpls()
	setMux()

	locGo := make(chan bool)
	go startListen(locGo)

	tmpGo := <-locGo
	chGoOn <- tmpGo
}

func startListen(chGoOn chan bool) {
	//	prefix := vomni.WebPrefix //v_cli.Param(v_cli.CliFileServPrefix)
	//	f_static := vutils.FileAbsPath(vomni.WebStaticPath, "")
	lPort := vparams.Params.PortWEBInternal //vparam.Params.WebPort
	netAddr := ""                           // v_cli.Param(v_cli.CliNetAddr)

	listenAddr := ":" + strconv.Itoa(lPort)
	fmt.Println("Listen...", listenAddr)

	//mux := http.NewServeMux() // ???????????
	if netAddr != "" {
		netType := "udp"
		addrOut := "ip-address-found.txt"
		l, err := net.Listen(netType, netAddr)
		if err != nil {
			err = fmt.Errorf("Error! %s", err.Error())
			panic(err)
		}

		tmpStr := l.Addr().Network() + ":" + l.Addr().String()

		err = ioutil.WriteFile(addrOut, []byte(tmpStr), 0644)
		if err != nil {
			err = fmt.Errorf("Error! %s", err.Error())
			panic(err)
		}

		chGoOn <- true
		s := &http.Server{}
		panic(s.Serve(l))
	} else {
		chGoOn <- true
		panic(http.ListenAndServe(listenAddr, nil))
	}
}
