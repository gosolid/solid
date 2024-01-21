//go:build darwin && !ios

package main

/*
#cgo darwin CFLAGS: -x objective-c -mmacosx-version-min=10.14
#cgo darwin LDFLAGS: -lobjc -framework AppKit -framework CoreFoundation -framework CoreGraphics -framework WebKit -mmacosx-version-min=10.14
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>
#include <AppKit/AppKit.h>
#include <WebKit/WebKit.h>


@interface AppDelegate : NSObject <NSApplicationDelegate>
@property(strong, nonatomic) NSWindow *window;
@property(strong, nonatomic) WKWebView *webView;
@property(strong, nonatomic) NSScrollView *scrollView;
@end

@implementation AppDelegate

- (void)applicationDidFinishLaunching:(NSNotification*)notification {
	NSView *container = [[NSView alloc] initWithFrame:CGRectMake(0, 0, 640, 480)];
	container.autoresizingMask = NSViewWidthSizable | NSViewHeightSizable;


	self.webView = [[WKWebView alloc] initWithFrame:container.bounds];
	self.webView.autoresizingMask = NSViewWidthSizable | NSViewHeightSizable;
	[self.webView setValue:@false forKey:@"drawsBackground"];
	[self.webView setValue:@true forKey:@"drawsTransparentBackground"];
	[container addSubview:self.webView];

	self.webView.translatesAutoresizingMaskIntoConstraints = false;
	[NSLayoutConstraint activateConstraints:@[
			[self.webView.leadingAnchor constraintEqualToAnchor:container.safeAreaLayoutGuide.leadingAnchor],
			[self.webView.trailingAnchor constraintEqualToAnchor:container.safeAreaLayoutGuide.trailingAnchor],
			[self.webView.topAnchor constraintEqualToAnchor:container.safeAreaLayoutGuide.topAnchor],
			[self.webView.bottomAnchor constraintEqualToAnchor:container.safeAreaLayoutGuide.bottomAnchor]
	]];

	self.window = [[NSWindow alloc] initWithContentRect:container.frame
		styleMask:NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskTitled | NSWindowStyleMaskResizable
		backing:NSBackingStoreBuffered
		defer:false];
	self.window.titlebarAppearsTransparent = true;
	self.window.movableByWindowBackground = true;
	self.window.title = @"Grexie Cloud";
	self.window.contentView = container;
	self.window.titlebarSeparatorStyle = NSTitlebarSeparatorStyleShadow;
  [self.window makeKeyAndOrderFront:nil];
	[self.window center];

	NSURL *url = [NSURL URLWithString:@"http://localhost:3086/"];
	NSURLRequest *urlReq = [NSURLRequest requestWithURL:url];

	[self.webView loadRequest:urlReq];

	container.additionalSafeAreaInsets = NSEdgeInsetsMake(
		 0,
		 0,
		 0,
		 0
	);

}

@end

int macosmain() {
  @autoreleasepool {
		[[NSApplication sharedApplication] setDelegate:[[AppDelegate alloc] init]];
		[[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyRegular];
		[[NSApplication sharedApplication] activateIgnoringOtherApps:YES];
  }

	int argc = 0;
	const char** argv = {0};
  return NSApplicationMain(argc, argv);
}

*/
import "C"

import (
	"runtime"
)

func main() {
	runtime.LockOSThread()

	// startWorker()

	C.macosmain()
}
