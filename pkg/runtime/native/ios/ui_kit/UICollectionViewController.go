//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewController : UIKit.UIViewController
*/

type UICollectionViewController struct {
  *UIViewController

}

func (r *UICollectionViewController) SetCollectionView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) SetClearsSelectionOnViewWillAppear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) SetUseLayoutToLayoutNavigationTransitions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) InstallsStandardGestureForInteractiveMovement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) CollectionView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) ClearsSelectionOnViewWillAppear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) UseLayoutToLayoutNavigationTransitions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) CollectionViewLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) SetInstallsStandardGestureForInteractiveMovement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) InitWithCollectionViewLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

