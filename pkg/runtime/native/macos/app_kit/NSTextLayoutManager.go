//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextLayoutManager : objc.NSObject
*/

type NSTextLayoutManager struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *NSTextSelectionDataSource
}

func (r *NSTextLayoutManager) EnumerateTextLayoutFragmentsFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetUsesFontLeading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextSelections() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) LinkRenderingAttributes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) EnsureLayoutForRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextLayoutFragmentForPosition() (*NSTextLayoutFragment, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetRenderingAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextViewportLayoutController() (*NSTextViewportLayoutController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) LayoutQueue() (*foundation.NSOperationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetTextSelectionNavigation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) InvalidateRenderingAttributesForTextRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) RenderingAttributesForLink() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) EnumerateTextSegmentsInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) ReplaceContentsInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) AddRenderingAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) ReplaceTextContentManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) EnsureLayoutForBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) InvalidateLayoutForRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextContainer() (*NSTextContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetTextSelections() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetLimitsLayoutForSuspiciousContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) RenderingAttributesValidator() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) LimitsLayoutForSuspiciousContents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) UsesHyphenation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetLayoutQueue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextSelectionNavigation() (*NSTextSelectionNavigation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) EnumerateRenderingAttributesFromLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextContentManager() (*NSTextContentManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) UsageBoundsForTextContainer() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetRenderingAttributesValidator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) UsesFontLeading() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) RemoveRenderingAttribute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetUsesHyphenation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetTextContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextLayoutFragmentForLocation() (*NSTextLayoutFragment, error) {
  return nil, fmt.Errorf("unimplemented")
}

