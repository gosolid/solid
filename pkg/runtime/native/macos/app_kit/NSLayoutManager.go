//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSLayoutManager : objc.NSObject
*/

type NSLayoutManager struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSLayoutManager) EnsureLayoutForTextContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetUsesDefaultHyphenation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ExtraLineFragmentRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InvalidateGlyphsForCharacterRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) BackgroundLayoutEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TypesetterBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetExtraLineFragmentRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetTypesetterBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextStorage() (*NSTextStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureGlyphsForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureLayoutForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetDrawsOutsideLineFragment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnumerateLineFragmentsForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetLimitsLayoutForSuspiciousContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) BoundingRectForGlyphRange() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) StrikethroughGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) RemoveTemporaryAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) AllowsNonContiguousLayout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphIndexForPoint() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) CharacterIndexForPoint() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawStrikethroughForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) UsesDefaultHyphenation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetAllowsNonContiguousLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) CGGlyphAtIndex() (uint16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphRangeForBoundingRect() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) FractionOfDistanceThroughGlyphForPoint() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GetLineFragmentInsertionPointsForCharacterAtIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureGlyphsForCharacterRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphRangeForTextContainer() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) FillBackgroundRectArray() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ShowsInvisibleCharacters() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetLayoutRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LayoutRectForTextBlock() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DefaultLineHeightForFont() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextContainerChangedGeometry() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureLayoutForBoundingRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) PropertyForGlyphAtIndex() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) FirstUnlaidCharacterIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) HasNonContiguousLayout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) RemoveTextContainerAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphIndexForCharacterAtIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) AddTemporaryAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureLayoutForCharacterRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextContainerForGlyphAtIndex() (*NSTextContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LineFragmentRectForGlyphAtIndex() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ExtraLineFragmentTextContainer() (*NSTextContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ReplaceTextStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InvalidateDisplayForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ShowsControlCharacters() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetUsesFontLeading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetLineFragmentRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetBoundsRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetTextContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphRangeForCharacterRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GetGlyphsInRange() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetShowsControlCharacters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetDefaultAttachmentScaling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) BoundsRectForTextBlock() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetShowsInvisibleCharacters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) UsesFontLeading() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ProcessEditingForTextStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetTemporaryAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) AddTextContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InvalidateLayoutForCharacterRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) Typesetter() (*NSTypesetter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetGlyphs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LineFragmentUsedRectForGlyphAtIndex() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetTextStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextContainerChangedTextView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DefaultAttachmentScaling() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) FirstUnlaidGlyphIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawBackgroundForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) AddTemporaryAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InsertTextContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) UsedRectForTextContainer() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LocationForGlyphAtIndex() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DefaultBaselineOffsetForFont() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InvalidateDisplayForCharacterRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) AttachmentSizeForGlyphAtIndex() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LimitsLayoutForSuspiciousContents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawUnderlineForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TemporaryAttributesAtCharacterIndex() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ShowCGGlyphs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) UnderlineGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetTypesetter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) IsValidGlyphIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ShowAttachmentCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextContainers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetAttachmentSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) RangeOfNominallySpacedGlyphsContainingIndex() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawGlyphsForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ExtraLineFragmentUsedRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetNotShownAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GetFirstUnlaidCharacterIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetBackgroundLayoutEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnumerateEnclosingRectsForGlyphRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TemporaryAttribute() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) NumberOfGlyphs() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) CharacterIndexForGlyphAtIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) NotShownAttributeForGlyphAtIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) CharacterRangeForGlyphRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayout() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

