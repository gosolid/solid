//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIButton : UIKit.UIControl
*/

type UIButton struct {
  *UIControl

}

func (r *UIButton) CurrentTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SystemButtonWithImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) Menu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) IsHeld() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetPointerInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) CurrentTitleShadowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) UpdateConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) PointerStyleProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) CurrentBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) CurrentImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetNeedsUpdateConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) IsPointerInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SystemButtonWithPrimaryAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) TitleShadowColorForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) TitleLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) ConfigurationUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetAutomaticallyUpdatesConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SubtitleLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) ButtonWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetPreferredMenuElementOrder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) CurrentTitleColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) ImageView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) TitleForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) Role() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetRole() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetPreferredSymbolConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) PreferredSymbolConfigurationForImageInState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) IsHovered() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) BackgroundImageForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetConfigurationUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) TitleColorForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetTitleShadowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) TintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) ChangesSelectionAsPrimaryAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetPointerStyleProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) CurrentPreferredSymbolConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) CurrentAttributedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) ButtonWithType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) AttributedTitleForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) PreferredMenuElementOrder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetChangesSelectionAsPrimaryAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) ButtonType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetMenu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetAttributedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) ImageForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) AutomaticallyUpdatesConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIButton) SetTitleColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

