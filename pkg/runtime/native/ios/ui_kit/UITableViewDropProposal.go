//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITableViewDropProposal : UIKit.UIDropProposal
*/

type UITableViewDropProposal struct {
  *UIDropProposal

}

func (r *UITableViewDropProposal) InitWithDropOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDropProposal) Intent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

