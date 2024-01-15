//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLResponse : objc.NSObject
*/

type NSURLResponse struct {
  *objc.NSObject

}

func (r *NSURLResponse) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) MIMEType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) ExpectedContentLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) TextEncodingName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) SuggestedFilename() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

