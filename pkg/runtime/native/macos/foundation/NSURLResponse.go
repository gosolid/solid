//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSURLResponse : objc.NSObject
*/

type NSURLResponse struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSURLResponse) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) URL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) MIMEType() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) ExpectedContentLength() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) TextEncodingName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLResponse) SuggestedFilename() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

