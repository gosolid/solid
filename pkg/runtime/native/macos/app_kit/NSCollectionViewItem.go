//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionViewItem : AppKit.NSViewController
*/

type NSCollectionViewItem struct {
  *NSViewController
  *foundation.NSCopying
  *NSCollectionViewElement
}

func (r *NSCollectionViewItem) SetHighlightState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) ImageView() (*NSImageView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) TextField() (*NSTextField, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) SetTextField() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) IsSelected() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) SetSelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) SetImageView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) DraggingImageComponents() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) CollectionView() (*NSCollectionView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewItem) HighlightState() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

