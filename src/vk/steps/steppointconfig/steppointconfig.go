package steppointconfig

import (
	"time"
	vomni "vk/omnibus"
	vpointcfg "vk/pointconfig"
	vstep "vk/steps/step"
)

var isRunning bool

type thisStep vstep.StepVars

var ThisStep thisStep

func init() {
	ThisStep.Name = vomni.StepNamePointConfig
	ThisStep.Err = make(chan error)
	ThisStep.GoOn = make(chan bool)
	ThisStep.Done = make(chan int)
}

func (s *thisStep) stepDo() {
	// put code here to do what is necessary

	chErr := make(chan error)
	chDone := make(chan int)
	chGoOn := make(chan bool)

	go vpointcfg.GetPointCfg(chGoOn, chDone, chErr) // put the right call here

	for {
		select {
		case err := <-chErr:
			s.Err <- err
		case done := <-chDone:
			s.Done <- done
		case <-chGoOn:
			s.GoOn <- true
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
