//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLComponents : objc.NSObject
*/

type NSURLComponents struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSURLComponents) Port() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPercentEncodedPassword() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) PercentEncodedHost() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) PercentEncodedQuery() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) QueryItems() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) PercentEncodedQueryItems() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) ComponentsWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPercentEncodedUser() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) URLRelativeToURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPercentEncodedHost() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPercentEncodedQueryItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetHost() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) PercentEncodedPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) RangeOfFragment() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPassword() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) PercentEncodedUser() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetEncodedHost() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) RangeOfPassword() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) String() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPort() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPercentEncodedPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetQueryItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) User() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) Query() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) RangeOfQuery() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPercentEncodedQuery() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) PercentEncodedFragment() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) Host() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) Path() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) PercentEncodedPassword() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetPercentEncodedFragment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetScheme() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) Fragment() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) EncodedHost() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) ComponentsWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetQuery() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) RangeOfScheme() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) RangeOfPath() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) Scheme() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) Password() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) URL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) RangeOfHost() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) RangeOfPort() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetUser() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) SetFragment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLComponents) RangeOfUser() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

