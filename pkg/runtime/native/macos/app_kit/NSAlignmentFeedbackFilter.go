//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSAlignmentFeedbackFilter : objc.NSObject
*/

type NSAlignmentFeedbackFilter struct {
  *objc.NSObject

}

func (r *NSAlignmentFeedbackFilter) AlignmentFeedbackTokenForVerticalMovementInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlignmentFeedbackFilter) PerformFeedback() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlignmentFeedbackFilter) InputEventMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAlignmentFeedbackFilter) UpdateWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlignmentFeedbackFilter) UpdateWithPanRecognizer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlignmentFeedbackFilter) AlignmentFeedbackTokenForMovementInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlignmentFeedbackFilter) AlignmentFeedbackTokenForHorizontalMovementInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

