//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollisionBehavior : UIKit.UIDynamicBehavior
*/

type UICollisionBehavior struct {
  *UIDynamicBehavior

}

func (r *UICollisionBehavior) RemoveBoundaryWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) SetCollisionMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) TranslatesReferenceBoundsIntoBoundary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) BoundaryWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) AddItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) Items() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) CollisionDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) SetCollisionDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) InitWithItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) SetTranslatesReferenceBoundsIntoBoundary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) BoundaryIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) CollisionMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) SetTranslatesReferenceBoundsIntoBoundaryWithInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) AddBoundaryWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) RemoveAllBoundaries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollisionBehavior) RemoveItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

