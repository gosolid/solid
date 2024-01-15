//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSCell : objc.NSObject
*/

type NSCell struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
  *NSUserInterfaceItemIdentification
  *NSAccessibilityElement
  *NSAccessibility
}

func (r *NSCell) TakeDoubleValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) FloatValue() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) SelectWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) EditWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) MenuForEvent() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) TakeObjectValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) Highlight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) ImageRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSCell) IsHighlighted() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) KeyEquivalent() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetFormatter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetDoubleValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetAllowsUndo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) TitleRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSCell) CalcDrawInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) ContinueTracking() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetIntegerValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetCellAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) HighlightColorWithFrame() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetSelectable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) UsesSingleLineMode() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) TakeStringValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) PrefersTrackingUntilMouseUp() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetControlView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) IsScrollable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetUpFieldEditorAttributes() (*NSText, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetStringValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SendsActionOnEndEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) AllowsUndo() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetUsesSingleLineMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetControlSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) MouseDownFlags() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetBaseWritingDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) TruncatesLastVisibleLine() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) CellAttribute() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) TakeFloatValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) DrawingRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSCell) DrawWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) TrackMouse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) Formatter() (*foundation.NSFormatter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) Alignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetWraps() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) TakeIntValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) GetPeriodicDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetBezeled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) TakeIntegerValueFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) DrawInteriorWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) StartTrackingAt() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) DraggingImageComponentsWithFrame() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) IsBordered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) DefaultMenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) StringValue() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetFloatValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetRepresentedObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) CellSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCell) Tag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) BaseWritingDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetTruncatesLastVisibleLine() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) Wraps() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetIntValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) DoubleValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) LineBreakMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) ControlView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) ObjectValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) HasValidObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) ControlSize() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetUserInterfaceLayoutDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) ResetCursorRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) StopTracking() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) EndEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetContinuous() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) IsBezeled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) IntValue() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) CellSizeForBounds() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) IsSelectable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetSendsActionOnEndEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) SetLineBreakMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) InitImageCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) FieldEditorForView() (*NSTextView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetHighlighted() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) IntegerValue() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) SendActionOn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) Font() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) RepresentedObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCell) IsOpaque() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) SetScrollable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCell) UserInterfaceLayoutDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) IsContinuous() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCell) State() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCell) Compare() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

