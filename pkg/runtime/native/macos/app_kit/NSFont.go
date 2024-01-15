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
interface AppKit.NSFont : objc.NSObject
*/

type NSFont struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSFont) FamilyName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) Leading() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) IsFixedPitch() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFont) VerticalFont() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) TitleBarFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) AdvancementForCGGlyph() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSFont) Ascender() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) UnderlineThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) LabelFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) MonospacedDigitSystemFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) MenuBarFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) SmallSystemFontSize() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) FontName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) NumberOfGlyphs() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) MessageFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) GetBoundingRects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFont) SystemFontSizeForControlSize() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) BoundingRectForCGGlyph() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSFont) MaximumAdvancement() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSFont) ItalicAngle() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) CapHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) SetUserFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFont) SetUserFixedPitchFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFont) Matrix() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) FontDescriptor() (*NSFontDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) DisplayName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) TextTransform() (*foundation.NSAffineTransform, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) MonospacedSystemFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) BoldSystemFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) PaletteFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) ControlContentFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) PointSize() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) UnderlinePosition() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) FontWithDescriptor() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) SystemFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) XHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) IsVertical() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFont) MenuFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) MostCompatibleStringEncoding() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) SetInContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFont) LabelFontSize() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) CoveredCharacterSet() (*foundation.NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) BoundingRectForFont() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSFont) UserFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) SystemFontSize() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) GetAdvancements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFont) Set() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFont) FontWithName() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) FontWithSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) Descender() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFont) UserFixedPitchFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFont) ToolTipsFontOfSize() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

