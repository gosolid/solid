#define __OBJC2__ 1

#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>
#include <AppKit/AppKit.h>

@interface AppDelegate : NSObject <NSApplicationDelegate>
@property(strong, nonatomic) NSWindow *window;
@end