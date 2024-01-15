//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSExtensionRequestHandling
*/

type NSExtensionRequestHandling struct {

}

func (r *NSExtensionRequestHandling) BeginRequestWithExtensionContext() error {
  return fmt.Errorf("unimplemented")
}

