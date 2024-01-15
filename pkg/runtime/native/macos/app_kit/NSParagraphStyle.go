//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSParagraphStyle : objc.NSObject
*/

type NSParagraphStyle struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSMutableCopying
  *foundation.NSSecureCoding
}

func (r *NSParagraphStyle) DefaultWritingDirectionForLanguage() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) ParagraphSpacingBefore() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) DefaultTabInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) ParagraphSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) Alignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) MinimumLineHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) BaseWritingDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) TextLists() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) LineBreakStrategy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) LineHeightMultiple() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) HyphenationFactor() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) DefaultParagraphStyle() (*NSParagraphStyle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) LineSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) TailIndent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) FirstLineHeadIndent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) MaximumLineHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) LineBreakMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) UsesDefaultHyphenation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) TabStops() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) AllowsDefaultTighteningForTruncation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) TighteningFactorForTruncation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) HeadIndent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) TextBlocks() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSParagraphStyle) HeaderLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

