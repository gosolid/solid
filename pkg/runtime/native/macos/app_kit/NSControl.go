//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSControl : AppKit.NSView
*/

type NSControl struct {
  *NSView

}

func (r *NSControl) IgnoresMultiClick() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControl) IntValue() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetUsesSingleLineMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) RefusesFirstResponder() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetIntegerValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) BaseWritingDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) LineBreakMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) Alignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) IsContinuous() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControl) IsHighlighted() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControl) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) TakeObjectValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) MouseDown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) TakeDoubleValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetContinuous() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetAttributedStringValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SendAction() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControl) ObjectValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetAllowsExpansionToolTips() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) IntegerValue() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetFloatValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetRefusesFirstResponder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) FloatValue() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) AllowsExpansionToolTips() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControl) SizeThatFits() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSControl) PerformClick() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) TakeStringValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSControl) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControl) StringValue() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSControl) AttributedStringValue() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSControl) UsesSingleLineMode() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControl) SizeToFit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) TakeFloatValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetLineBreakMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) ControlSize() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetControlSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SendActionOn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) TakeIntegerValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetIgnoresMultiClick() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetHighlighted() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) Formatter() (*foundation.NSFormatter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetFormatter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetStringValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) DoubleValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) ExpansionFrameWithFrame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSControl) DrawWithExpansionFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetIntValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetDoubleValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) Tag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSControl) Font() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSControl) SetBaseWritingDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) TakeIntValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControl) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

