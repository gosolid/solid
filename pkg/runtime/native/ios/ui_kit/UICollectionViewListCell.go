//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewListCell : UIKit.UICollectionViewCell
*/

type UICollectionViewListCell struct {
  *UICollectionViewCell

}

func (r *UICollectionViewListCell) DefaultContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) IndentationWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) IndentsAccessories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) SetAccessories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) SeparatorLayoutGuide() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) IndentationLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) SetIndentationLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) SetIndentationWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) SetIndentsAccessories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewListCell) Accessories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

