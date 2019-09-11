package step

type StepVars struct {
	Name string
	Err  chan error
	GoOn chan bool
	Done chan int
}

type Step interface {
	StepExec(chDone chan int, chGoOn chan bool, chErr chan error)
	StepName() string
	StepPost(chan bool)
	StepPre(chDone chan int, chGoOn chan bool, chErr chan error)
}

func Execute(s Step, chMainDone chan int, chMainGoOn chan bool, chMainErr chan error) {
	chErr := make(chan error)
	chDone := make(chan int)
	chGoOn := make(chan bool)

	go s.StepPre(chDone, chGoOn, chErr)
	select {
	case done := <-chDone:
		chMainDone <- done
		return
	case err := <-chErr:
		chMainErr <- err
		return
	case <-chGoOn:
	}

	go s.StepExec(chDone, chGoOn, chErr)

	//	for {
	select {
	case done := <-chDone:
		chMainDone <- done
		return
	case err := <-chErr:
		chMainErr <- err
		return
	case <-chGoOn:
		chMainGoOn <- true
	}
	//}
}
