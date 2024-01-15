//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSMutableParagraphStyle : AppKit.NSParagraphStyle
*/

type NSMutableParagraphStyle struct {
  *NSParagraphStyle

}

func (r *NSMutableParagraphStyle) SetLineHeightMultiple() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) ParagraphSpacingBefore() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetParagraphSpacingBefore() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) LineBreakStrategy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) AddTabStop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) RemoveTabStop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetFirstLineHeadIndent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) MaximumLineHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) UsesDefaultHyphenation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) TighteningFactorForTruncation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) HeadIndent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) TailIndent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) LineHeightMultiple() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetHyphenationFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) Alignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) TabStops() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) FirstLineHeadIndent() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetUsesDefaultHyphenation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetAllowsDefaultTighteningForTruncation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) TextBlocks() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetTailIndent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) LineBreakMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetLineBreakMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetDefaultTabInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) LineSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetLineSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) ParagraphSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetParagraphSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetLineBreakStrategy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetMaximumLineHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetBaseWritingDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) DefaultTabInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetHeadIndent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetTighteningFactorForTruncation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetHeaderLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) HyphenationFactor() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetTabStops() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) AllowsDefaultTighteningForTruncation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetTextBlocks() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetParagraphStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) MinimumLineHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetMinimumLineHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) BaseWritingDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) TextLists() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) SetTextLists() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableParagraphStyle) HeaderLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

