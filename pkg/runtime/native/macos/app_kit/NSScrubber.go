//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSScrubber : AppKit.NSView
*/

type NSScrubber struct {
  *NSView

}

func (r *NSScrubber) ScrollItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) RegisterNib() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) ScrubberLayout() (*NSScrubberLayout, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetFloatsSelectionViews() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SelectionBackgroundStyle() (*NSScrubberSelectionStyle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) ItemViewForItemAtIndex() (*NSScrubberItemView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) NumberOfItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) ShowsArrowButtons() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) ReloadData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) PerformSequentialBatchUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) ReloadItemsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) MoveItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) RegisterClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetScrubberLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetItemAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SelectionOverlayStyle() (*NSScrubberSelectionStyle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) MakeItemWithIdentifier() (*NSScrubberItemView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetContinuous() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetBackgroundView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) InsertItemsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) HighlightedIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) ItemAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) FloatsSelectionViews() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetSelectionBackgroundStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) BackgroundView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SelectedIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) IsContinuous() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetSelectionOverlayStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetShowsArrowButtons() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) Mode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) ShowsAdditionalContentIndicators() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetShowsAdditionalContentIndicators() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) RemoveItemsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubber) SetSelectedIndex() error {
  return fmt.Errorf("unimplemented")
}

