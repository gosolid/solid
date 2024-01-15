//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITextDropProposal : UIKit.UIDropProposal
*/

type UITextDropProposal struct {
  *UIDropProposal

}

func (r *UITextDropProposal) DropAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDropProposal) SetDropAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDropProposal) DropProgressMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDropProposal) SetDropProgressMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDropProposal) UseFastSameViewOperations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDropProposal) SetUseFastSameViewOperations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDropProposal) DropPerformer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDropProposal) SetDropPerformer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

