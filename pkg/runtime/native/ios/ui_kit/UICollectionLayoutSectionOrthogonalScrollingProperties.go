//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionLayoutSectionOrthogonalScrollingProperties : objc.NSObject
*/

type UICollectionLayoutSectionOrthogonalScrollingProperties struct {
  *objc.NSObject

}

func (r *UICollectionLayoutSectionOrthogonalScrollingProperties) Bounce() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutSectionOrthogonalScrollingProperties) SetBounce() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutSectionOrthogonalScrollingProperties) DecelerationRate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutSectionOrthogonalScrollingProperties) SetDecelerationRate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

