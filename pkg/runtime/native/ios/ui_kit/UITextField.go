//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITextField : UIKit.UIControl
*/

type UITextField struct {
  *UIControl

}

func (r *UITextField) AttributedText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) TypingAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetTextColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetPlaceholder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetInputAccessoryView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) EditingRectForBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) RightViewRectForBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) BorderStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetInputView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) DrawPlaceholderInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetTextAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) DefaultTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetLeftViewMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetRightViewMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) InputAccessoryView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetAttributedText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) MinimumFontSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) ClearsOnInsertion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) Background() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) DisabledBackground() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetDisabledBackground() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetClearButtonMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) LeftViewMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) AllowsEditingTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) RightViewMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetAttributedPlaceholder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) AdjustsFontSizeToFitWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetBackground() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) BorderRectForBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) TextColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) ClearButtonMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) TextRectForBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) LeftViewRectForBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) TextAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetBorderStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) Placeholder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetClearsOnInsertion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) PlaceholderRectForBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) DrawTextInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) AttributedPlaceholder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) IsEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) LeftView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetRightView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) ClearButtonRectForBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) Text() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) ClearsOnBeginEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetMinimumFontSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetDefaultTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetAdjustsFontSizeToFitWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetAllowsEditingTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetTypingAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetLeftView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) RightView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) InputView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) Font() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextField) SetClearsOnBeginEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

