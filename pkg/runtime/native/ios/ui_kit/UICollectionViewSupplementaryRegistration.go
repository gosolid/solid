//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewSupplementaryRegistration : objc.NSObject
*/

type UICollectionViewSupplementaryRegistration struct {
  *objc.NSObject

}

func (r *UICollectionViewSupplementaryRegistration) RegistrationWithSupplementaryNib() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewSupplementaryRegistration) SupplementaryClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewSupplementaryRegistration) SupplementaryNib() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewSupplementaryRegistration) ElementKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewSupplementaryRegistration) ConfigurationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewSupplementaryRegistration) RegistrationWithSupplementaryClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

