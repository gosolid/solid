
#include "Chromium.h"
#include "AppDelegate.h"

#include <objc/message.h>
#include <stdlib.h>


@implementation AppDelegate

- (void)applicationDidFinishLaunching:(NSNotification*)notification {
	NSView *container = [[NSView alloc] initWithFrame:CGRectMake(0, 0, 640, 480)];
	container.autoresizingMask = NSViewWidthSizable | NSViewHeightSizable;

  // WebContentsPtr pWebContents = Chromium_CreateWebContents();
  // NSView* nativeView = (NSView*)Chromium_WebContents_GetNativeView(pWebContents);
  // nativeView.autoresizingMask = NSViewWidthSizable | NSViewHeightSizable;

	self.window = [[NSWindow alloc] initWithContentRect:container.frame
		styleMask:NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskTitled | NSWindowStyleMaskResizable
		backing:NSBackingStoreBuffered
		defer:false];
	// self.window.titlebarAppearsTransparent = true;
	// self.window.movableByWindowBackground = true;
	self.window.title = @"Grexie Cloud";
  // self.window.contentView = nativeView;
	// self.window.titlebarSeparatorStyle = NSTitlebarSeparatorStyleShadow;
  [self.window makeKeyAndOrderFront:nil];
	[self.window center];
}

@end

