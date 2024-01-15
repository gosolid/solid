//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewCell : UIKit.UICollectionReusableView
*/

type UICollectionViewCell struct {
  *UICollectionReusableView

}

func (r *UICollectionViewCell) SetSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SelectedBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) DragStateDidChange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) DefaultBackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) ConfigurationUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) ContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) IsSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetConfigurationUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) AutomaticallyUpdatesBackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) BackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetNeedsUpdateConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) ConfigurationState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) ContentView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) BackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetBackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetAutomaticallyUpdatesBackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetSelectedBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) UpdateConfigurationUsingState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) AutomaticallyUpdatesContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetAutomaticallyUpdatesContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) IsHighlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCell) SetHighlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

