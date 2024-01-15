//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSLayoutManager : objc.NSObject
*/

type NSLayoutManager struct {
  *objc.NSObject

}

func (r *NSLayoutManager) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetDrawsOutsideLineFragment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) AttachmentSizeForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LimitsLayoutForSuspiciousContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) CharacterRangeForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetTextStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) CharacterIndexForPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetAllowsNonContiguousLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) HasNonContiguousLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetLimitsLayoutForSuspiciousContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InvalidateDisplayForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureGlyphsForCharacterRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetAttachmentSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetUsesDefaultHyphenation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) BoundingRectForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) UnderlineGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) AddTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureGlyphsForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureLayoutForTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) RangeOfNominallySpacedGlyphsContainingIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetUsesFontLeading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ProcessEditingForTextStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) FirstUnlaidGlyphIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnumerateEnclosingRectsForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawGlyphsForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) IsValidGlyphIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetExtraLineFragmentRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawStrikethroughForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) UsesDefaultHyphenation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetGlyphs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) PropertyForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetLineFragmentRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) UsedRectForTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InvalidateDisplayForCharacterRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) FractionOfDistanceThroughGlyphForPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ShowsInvisibleCharacters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) UsesFontLeading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) NumberOfGlyphs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ExtraLineFragmentTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) RemoveTextContainerAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureLayoutForCharacterRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureLayoutForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetShowsInvisibleCharacters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextContainerChangedGeometry() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LineFragmentRectForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) NotShownAttributeForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphRangeForBoundingRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) AllowsNonContiguousLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ExtraLineFragmentUsedRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnsureLayoutForBoundingRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) CGGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphIndexForPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawBackgroundForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GetLineFragmentInsertionPointsForCharacterAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) FillBackgroundRectArray() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextContainers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InvalidateLayoutForCharacterRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphIndexForCharacterAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphRangeForCharacterRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TextContainerForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) EnumerateLineFragmentsForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InsertTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InvalidateGlyphsForCharacterRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GetGlyphsInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GetFirstUnlaidCharacterIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetShowsControlCharacters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) CharacterIndexForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) FirstUnlaidCharacterIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LineFragmentUsedRectForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) LocationForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) StrikethroughGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ShowsControlCharacters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) SetNotShownAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) GlyphRangeForTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ShowCGGlyphs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) DrawUnderlineForGlyphRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManager) ExtraLineFragmentRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

