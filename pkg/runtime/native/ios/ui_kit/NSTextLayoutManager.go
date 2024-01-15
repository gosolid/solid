//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextLayoutManager : objc.NSObject
*/

type NSTextLayoutManager struct {
  *objc.NSObject

}

func (r *NSTextLayoutManager) EnsureLayoutForRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) RemoveRenderingAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetLayoutQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) EnsureLayoutForBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextLayoutFragmentForLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetLimitsLayoutForSuspiciousContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) LayoutQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) ReplaceTextContentManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetUsesFontLeading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) UsesHyphenation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextViewportLayoutController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextSelections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetTextSelections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) EnumerateTextLayoutFragmentsFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) ReplaceContentsInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetTextSelectionNavigation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) LinkRenderingAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextLayoutFragmentForPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) InvalidateLayoutForRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) EnumerateRenderingAttributesFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetUsesHyphenation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextContentManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) RenderingAttributesValidator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) InvalidateRenderingAttributesForTextRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) UsageBoundsForTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetRenderingAttributesValidator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) AddRenderingAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) RenderingAttributesForLink() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) UsesFontLeading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) SetRenderingAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) LimitsLayoutForSuspiciousContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) TextSelectionNavigation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutManager) EnumerateTextSegmentsInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

