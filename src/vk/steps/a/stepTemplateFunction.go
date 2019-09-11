package step<NAME>

import (
	"fmt"
	"time"
	vomni "vk/omnibus"
	"vk/steps/step"
)

type thisStep step.StepVars
var ThisStep thisStep

func init() {
	ThisStep.Name =vomni.<NAME>
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	chErr := make(chan error)
	chDone := make(chan int)
	chGoOn := make(chan bool)
//	go vudp.Server(chGoOn, chDone, chErr) // put the right call here

//!!!!!	for {
		select {
		case err := <-chErr:
			s.Err <- err
		case done := <-chDone:
			s.Done <- done
		case <-chGoOn:
			s.GoOn <- true
		}
//!!!!!	}
}

func (s *thisStep) StepName() string {
	return ThisStep.Name
}

func (s *thisStep) StepPre(chDone chan int, chGoOn chan bool, chErr chan error) {
	// do if something is required before the step execution

	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan int, chGoOn chan bool, chErr chan error) {

	// do what you would like
	go s.stepDo()

//!!!!!	for {
		select {
		case locErr := <-s.Err:
			vomni.StepErr <- locErr
		case locDone := <-s.Done:
			chDone <- locDone
		case locGoOn := <-s.GoOn:
			chGoOn <- locGoOn
		}
		time.Sleep(vomni.StepExecDelay)
//!!!!!	}
}

func (s *thisStep) StepPost() {
	// may be something needs to be done before leave the step
	// if not just send Done flag
	s.Done <- vomni.DoneStop
}
