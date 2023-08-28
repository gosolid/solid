//go:build darwin

package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func sudome() {
	if exe, err := os.Executable(); err != nil {
		panic(err)
	} else {
		cmd := exec.Command("osascript", "-e", fmt.Sprintf("do shell script \"sudo %s %s\" with prompt \"Solid\" with administrator privileges", exe, strings.Join(os.Args[1:], " ")))
		cmd.Start()
		os.Exit(0)
	}
}

const launchdPlist = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>com.grexie.Solid</string>
		<key>Program</key>
		<string>/usr/local/libexec/solid</string>
		<key>ProgramArguments</key>
		<array>
			<string>/usr/local/bin/solid</string>
			<string>--daemon</string>
		</array>
		<key>KeepAlive</key>
		<true/>
		<key>StartInterval</key>
    <integer>2</integer>
		<key>StandardErrorPath</key>
    <string>/var/log/solid.log</string>

    <key>StandardOutPath</key>
    <string>/var/log/solid.log</string>
	</dict>
</plist>
`

func installDaemon() error {
	if os.Getuid() != 0 {
		sudome()
		return nil
	}

	if exe, err := os.Executable(); err != nil {
		return err
	} else if err := os.MkdirAll("/usr/local/bin", 0755); err != nil {
		return err
	} else if err := os.Symlink(exe, "/usr/local/bin/solid"); err != nil {
		return err
	} else if err := os.WriteFile("/Library/LaunchDaemons/com.grexie.Solid.plist", []byte(launchdPlist), 0755); err != nil {
		return err
	} else if err := exec.Command("launchctl", "load", "/Library/LaunchDaemons/com.grexie.Solid.plist").Run(); err != nil {
		return err
	} else {
		return nil
	}
}

func uninstallDaemon() error {
	if os.Getuid() != 0 {
		sudome()
		return nil
	}

	exec.Command("launchctl", "remove", "/Library/LaunchDaemons/com.grexie.Solid.plist").Run()
	exec.Command("launchctl", "unload", "/Library/LaunchDaemons/com.grexie.Solid.plist").Run()
	exec.Command("launchctl", "disable", "/Library/LaunchDaemons/com.grexie.Solid.plist").Run()
	os.Remove("/Library/LaunchDaemons/com.grexie.Solid.plist")
	os.Remove("/usr/local/bin/solid")
	return nil
}

func startDaemon() error {
	if os.Getuid() != 0 {
		sudome()
		return nil
	}

	exec.Command("launchctl", "load", "/Library/LaunchDaemons/com.grexie.Solid.plist").Run()
	return nil
}

func stopDaemon() error {
	if os.Getuid() != 0 {
		sudome()
		return nil
	}

	exec.Command("launchctl", "unload", "/Library/LaunchDaemons/com.grexie.Solid.plist").Run()
	return nil
}
