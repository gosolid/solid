//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionLayoutListConfiguration : objc.NSObject
*/

type UICollectionLayoutListConfiguration struct {
  *objc.NSObject

}

func (r *UICollectionLayoutListConfiguration) SeparatorConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetHeaderTopPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) InitWithAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) HeaderMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetShowsSeparators() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) Appearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) ShowsSeparators() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) LeadingSwipeActionsConfigurationProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetLeadingSwipeActionsConfigurationProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) TrailingSwipeActionsConfigurationProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetSeparatorConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) ItemSeparatorHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetItemSeparatorHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetTrailingSwipeActionsConfigurationProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetHeaderMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) FooterMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) SetFooterMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionLayoutListConfiguration) HeaderTopPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

