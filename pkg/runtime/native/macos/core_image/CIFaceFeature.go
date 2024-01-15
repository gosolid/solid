//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

/*
interface CoreImage.CIFaceFeature : CoreImage.CIFeature
*/

type CIFaceFeature struct {
  *CIFeature

}

