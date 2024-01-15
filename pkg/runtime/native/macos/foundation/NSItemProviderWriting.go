//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSItemProviderWriting
*/

type NSItemProviderWriting struct {

}

func (r *NSItemProviderWriting) ItemProviderVisibilityForRepresentationWithTypeIdentifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSItemProviderWriting) LoadDataWithTypeIdentifier() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProviderWriting) WritableTypeIdentifiersForItemProvider() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

