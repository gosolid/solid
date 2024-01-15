//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSParagraphStyle : objc.NSObject
*/

type NSParagraphStyle struct {
  *objc.NSObject

}

func (r *NSParagraphStyle) Alignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) MinimumLineHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) MaximumLineHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) AllowsDefaultTighteningForTruncation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) LineBreakStrategy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) DefaultParagraphStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) LineSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) LineBreakMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) LineHeightMultiple() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) UsesDefaultHyphenation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) TailIndent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) FirstLineHeadIndent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) BaseWritingDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) HyphenationFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) TabStops() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) DefaultWritingDirectionForLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) HeadIndent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) DefaultTabInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) TextLists() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) ParagraphSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) ParagraphSpacingBefore() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

