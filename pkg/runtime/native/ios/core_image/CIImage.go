//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIImage : objc.NSObject
*/

type CIImage struct {
  *objc.NSObject

}

func (r *CIImage) ImageByCroppingToRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) YellowImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByClampingToRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) PixelBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithIOSurface() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) EmptyImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCGImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCGLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithIOSurface() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByColorMatchingWorkingSpaceToColorSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) Definition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithBitmapData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageBySamplingLinear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) RedImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) GrayImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCGImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByInsertingIntermediate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) CGImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingGaussianBlurWithSigma() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCGImageSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCGLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCVPixelBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageTransformForOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageTransformForCGOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByClampingToExtent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageBySettingAlphaOneInExtent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) WhiteImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) BlackImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) GreenImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) Properties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithMTLTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCVPixelBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithMTLTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingCGOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageBySamplingNearest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) BlueImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCVImageBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageBySettingProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByPremultiplyingAlpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ClearImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) Url() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByUnpremultiplyingAlpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) RegionOfInterestForImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) CyanImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) MagentaImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithCVImageBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByApplyingTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ColorSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) Extent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithBitmapData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageWithContentsOfURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) InitWithCGImageSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByCompositingOverImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImage) ImageByColorMatchingColorSpaceToWorkingSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

