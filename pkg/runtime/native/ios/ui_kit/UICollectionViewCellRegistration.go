//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewCellRegistration : objc.NSObject
*/

type UICollectionViewCellRegistration struct {
  *objc.NSObject

}

func (r *UICollectionViewCellRegistration) RegistrationWithCellNib() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCellRegistration) CellClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCellRegistration) CellNib() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCellRegistration) ConfigurationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCellRegistration) RegistrationWithCellClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

