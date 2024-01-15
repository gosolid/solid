//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextView : AppKit.NSText
*/

type NSTextView struct {
  *NSText
  *NSColorChanging
  *NSMenuItemValidation
  *NSUserInterfaceValidations
  *NSTextInputClient
  *NSTextLayoutOrientationProvider
  *NSDraggingSource
  *NSStandardKeyBindingResponding
  *NSTextInput
  *NSAccessibilityNavigableStaticText
  *NSTextContent
}

func (r *NSTextView) LowerBaseline() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) OrderFrontListPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) UpdateDragTypeRegistration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) ShouldDrawInsertionPoint() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextView) TextStorage() (*NSTextStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) InvalidateTextContainerOrigin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) AlignJustified() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) SetNeedsDisplayInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) ClickedOnLink() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) SetUsesAdaptiveColorMappingForDarkAppearance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) InsertText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) TightenKerning() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) ChangeLayoutOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) StronglyReferencesTextStorage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextView) SetBaseWritingDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) UseStandardKerning() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) TurnOffLigatures() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) TextContainer() (*NSTextContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) PerformFindPanelAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) LayoutManager() (*NSLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) SetAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) PerformValidatedReplacementInRange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextView) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) InitUsingTextLayoutManager() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) OrderFrontTablePanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) ChangeColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) SelectionRangeForProposedRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTextView) CharacterIndexForInsertionAtPoint() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextView) TurnOffKerning() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) StopSpeaking() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) SetTextContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) SetTextContainerInset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) TextContainerOrigin() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSTextView) TextViewUsingTextLayoutManager() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) SetConstrainedFrameSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) LoosenKerning() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) ToggleTraditionalCharacterShape() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) TextLayoutManager() (*NSTextLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) TextContentStorage() (*NSTextContentStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) DrawInsertionPointInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) TextContainerInset() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSTextView) OrderFrontSpacingPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) RulerView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) UpdateRuler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) UsesAdaptiveColorMappingForDarkAppearance() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextView) ReplaceTextContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) RaiseBaseline() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) UpdateFontPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) Outline() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) ChangeDocumentBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) DrawViewBackgroundInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) SetLayoutOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) OrderFrontLinkPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) StartSpeaking() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextView) UseStandardLigatures() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) UseAllLigatures() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextView) ChangeAttributes() error {
  return fmt.Errorf("unimplemented")
}

