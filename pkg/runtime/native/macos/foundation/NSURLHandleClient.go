//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLHandleClient
*/

type NSURLHandleClient struct {

}

func (r *NSURLHandleClient) URLHandle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandleClient) URLHandleResourceDidBeginLoading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandleClient) URLHandleResourceDidFinishLoading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandleClient) URLHandleResourceDidCancelLoading() error {
  return fmt.Errorf("unimplemented")
}

