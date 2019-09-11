package stepConfig

import (
	"time"
	vcfg "vk/cfg" //	"vk/start"
	vomni "vk/omnibus"
	"vk/steps/step"
)

var isRunning bool

type thisStep step.StepVars

var ThisStep thisStep

func init() {
	ThisStep.Name = vomni.StepNameConfig
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	chErr := make(chan error)
	chDone := make(chan bool)

	go vcfg.Cfg(chDone, chErr)

	for {
		select {
		case err := <-chErr:
			s.Err <- err
		case done := <-chDone:
			s.GoOn <- done
		}
	}
}

func (s *thisStep) StepName() string {
	return ThisStep.Name
}

func (s *thisStep) StepPre(chDone chan int, chGoOn chan bool, chErr chan error) {
	// do if something is required before the step execution

	chGoOn <- true
}

func (s *thisStep) StepExec(chDone chan int, chGoOn chan bool, chErr chan error) {

	defer func() { isRunning = false }()

	// do what you would like
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

func (s *thisStep) StepPost(done chan bool) {
	// may be something needs to be done before leave the step
	// if not just send Done flag
	time.Sleep(vomni.DelayStepExec)

	if isRunning {
		s.Done <- vomni.DonePostStop
	}

	done <- true
}
