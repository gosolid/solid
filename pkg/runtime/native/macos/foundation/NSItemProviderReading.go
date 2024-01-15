//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSItemProviderReading
*/

type NSItemProviderReading struct {

}

func (r *NSItemProviderReading) ObjectWithItemProviderData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSItemProviderReading) ReadableTypeIdentifiersForItemProvider() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

