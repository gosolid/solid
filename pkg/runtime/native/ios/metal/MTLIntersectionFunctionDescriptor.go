//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
interface Metal.MTLIntersectionFunctionDescriptor : Metal.MTLFunctionDescriptor
*/

type MTLIntersectionFunctionDescriptor struct {
  *MTLFunctionDescriptor

}

