//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIProgressView : UIKit.UIView
*/

type UIProgressView struct {
  *UIView

}

func (r *UIProgressView) SetProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) TrackTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) SetTrackTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) ProgressImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) TrackImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) ObservedProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) InitWithProgressViewStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) ProgressViewStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) Progress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) ProgressTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) SetProgressImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) SetProgressViewStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) SetProgressTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) SetTrackImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIProgressView) SetObservedProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

