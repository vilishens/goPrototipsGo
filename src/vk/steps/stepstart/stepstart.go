package stepstart

import (
	"time"
	vcli "vk/cli"
	vomni "vk/omnibus"
	"vk/steps/step"
)

var isRunning bool

type thisStep step.StepVars

var ThisStep thisStep

func init() {
	ThisStep.Name = vomni.StepNameStart
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	err := vcli.Init()

	if nil != err {
		s.Err <- err
		return
	}

	s.GoOn <- true
}

func (s *thisStep) StepName() string {
	return ThisStep.Name
}

func (s *thisStep) StepPost(done chan bool) {
	// may be something needs to be done before leave the step
	// if not just send Done flag
	time.Sleep(vomni.DelayStepExec)

	if isRunning {
		s.Done <- vomni.DonePostStop
	}

	done <- true
}

func (s *thisStep) StepPre(chDone chan int, chGoOn chan bool, chErr chan error) {
	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan int, chGoOn chan bool, chErr chan error) {

	defer func() { isRunning = false }()

	go s.stepDo()

	stop := false
	for !stop {
		select {
		case locErr := <-s.Err:
			vomni.StepErr <- locErr
			stop = true
		case locDone := <-s.Done:
			if locDone != vomni.DonePostStop {
				chDone <- locDone
			}
			stop = true
		case locGoOn := <-s.GoOn:
			isRunning = true
			chGoOn <- locGoOn
		}
		time.Sleep(vomni.DelayStepExec)
	}
}
