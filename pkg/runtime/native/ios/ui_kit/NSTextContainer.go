//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextContainer : objc.NSObject
*/

type NSTextContainer struct {
  *objc.NSObject

}

func (r *NSTextContainer) SetLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) Size() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) WidthTracksTextView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) InitWithSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) LineFragmentRectForProposedRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) TextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetLineBreakMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) MaximumNumberOfLines() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) HeightTracksTextView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetWidthTracksTextView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) ReplaceLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) LayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) ExclusionPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetExclusionPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) LineFragmentPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetLineFragmentPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetHeightTracksTextView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) LineBreakMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetMaximumNumberOfLines() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) IsSimpleRectangularTextContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

