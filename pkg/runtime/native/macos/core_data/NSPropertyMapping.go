//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSPropertyMapping : objc.NSObject
*/

type NSPropertyMapping struct {
  *objc.NSObject

}

func (r *NSPropertyMapping) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyMapping) ValueExpression() (*foundation.NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyMapping) SetValueExpression() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyMapping) UserInfo() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyMapping) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyMapping) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

