//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITableViewCell : UIKit.UIView
*/

type UITableViewCell struct {
  *UIView

}

func (r *UITableViewCell) ContentView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SelectionStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) ShouldIndentWhileEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) IndentationWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetNeedsUpdateConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) ContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetAccessoryType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) IsEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) DetailTextLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SelectedBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) ReuseIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) EditingAccessoryType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) ImageView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetAutomaticallyUpdatesBackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetEditingAccessoryType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetSeparatorInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) BackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) AccessoryView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SeparatorInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) InitWithStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) DefaultBackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) AutomaticallyUpdatesContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) IsHighlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) EditingAccessoryView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetEditingAccessoryView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) DidTransitionToState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetShouldIndentWhileEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) DragStateDidChange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetSelectedBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetIndentationWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) FocusStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) PrepareForReuse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetHighlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) ConfigurationUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetAutomaticallyUpdatesContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetBackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetShowsReorderControl() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) UpdateConfigurationUsingState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) ConfigurationState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) TextLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetMultipleSelectionBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) ShowingDeleteConfirmation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetFocusStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) UserInteractionEnabledWhileDragging() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) WillTransitionToState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) EditingStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) ShowsReorderControl() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) AccessoryType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetIndentationLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) DefaultContentConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetSelectionStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetUserInteractionEnabledWhileDragging() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetConfigurationUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) IsSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) AutomaticallyUpdatesBackgroundConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) BackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) MultipleSelectionBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) SetAccessoryView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewCell) IndentationLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

