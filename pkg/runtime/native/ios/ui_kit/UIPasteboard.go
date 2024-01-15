//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPasteboard : objc.NSObject
*/

type UIPasteboard struct {
  *objc.NSObject

}

func (r *UIPasteboard) SetColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) HasImages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) ValueForPasteboardType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) AddItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) GeneralPasteboard() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) NumberOfItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) Items() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) PasteboardWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) RemovePasteboardWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) HasStrings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) Images() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) Color() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetPersistent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) PasteboardTypesForItemSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) DetectPatternsForPatterns() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) IsPersistent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) HasURLs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) ChangeCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) ItemProviders() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) String() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) URLs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetURLs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetColors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) DataForPasteboardType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) PasteboardTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) Strings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetImages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) Colors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetItemProviders() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) DetectValuesForPatterns() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetStrings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) HasColors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) PasteboardWithUniqueName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) ContainsPasteboardTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) ItemSetWithPasteboardTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) ValuesForPasteboardType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteboard) SetURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

