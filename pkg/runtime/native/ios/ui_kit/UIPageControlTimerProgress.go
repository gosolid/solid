//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPageControlTimerProgress : UIKit.UIPageControlProgress
*/

type UIPageControlTimerProgress struct {
  *UIPageControlProgress

}

func (r *UIPageControlTimerProgress) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) InitWithPreferredDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) PauseTimer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) DurationForPage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) ResumeTimer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) ResetsToInitialPageAfterEnd() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) SetPreferredDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) SetDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) SetResetsToInitialPageAfterEnd() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) IsRunning() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlTimerProgress) PreferredDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

