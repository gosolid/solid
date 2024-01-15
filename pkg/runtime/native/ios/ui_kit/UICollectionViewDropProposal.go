//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewDropProposal : UIKit.UIDropProposal
*/

type UICollectionViewDropProposal struct {
  *UIDropProposal

}

func (r *UICollectionViewDropProposal) InitWithDropOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDropProposal) Intent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

