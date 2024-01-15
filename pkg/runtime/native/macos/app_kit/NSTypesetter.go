//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTypesetter : objc.NSObject
*/

type NSTypesetter struct {
  *objc.NSObject

}

func (r *NSTypesetter) PrintingAdjustmentInLayoutManager() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SetLineFragmentPadding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) AttributesForExtraLineFragment() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) EndLineWithGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) ParagraphSpacingBeforeGlyphAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) BaselineOffsetInLayoutManager() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SharedSystemTypesetterForBehavior() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SetAttributedString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) CurrentTextContainer() (*NSTextContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) BeginParagraph() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) LayoutCharactersInRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SharedSystemTypesetter() (*NSTypesetter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) LayoutGlyphsInLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SetUsesFontLeading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SetBidiProcessingEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) AttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) ParagraphGlyphRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) ParagraphCharacterRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) TextContainers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) TypesetterBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) HyphenationFactor() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) GetLineFragmentRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) ParagraphSeparatorCharacterRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SetParagraphGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) LineSpacingAfterGlyphAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SetTypesetterBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SetHyphenationFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) BidiProcessingEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) LayoutManager() (*NSLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) DefaultTypesetterBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) TextTabForGlyphLocation() (*NSTextTab, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) UsesFontLeading() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) ParagraphSeparatorGlyphRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) CurrentParagraphStyle() (*NSParagraphStyle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) ParagraphSpacingAfterGlyphAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) LineFragmentPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) EndParagraph() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) BeginLineWithGlyphAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SetHardInvalidation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) SubstituteFontForFont() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTypesetter) LayoutParagraphAtPoint() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

