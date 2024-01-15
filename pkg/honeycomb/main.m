#import "AppDelegate.h"
#import "Chromium.h"

int macosmain() {
  @autoreleasepool {
    Chromium_Initialize();
		// [[NSApplication sharedApplication] setDelegate:[[AppDelegate alloc] init]];
		[[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyRegular];
		[[NSApplication sharedApplication] activateIgnoringOtherApps:YES];
  }

	int argc = 0;
	const char** argv = {0};
  return NSApplicationMain(argc, argv);
}
