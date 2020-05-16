# Netstatus

Netstatus is a simple internet connection monitor that sits in the status bar of your mac and shows wheather your connection is stable.

## Dependencies

Netstatus relies on the [AnyBar](https://github.com/tonsky/AnyBar) to display the connection status.
The easiest way to install AnyBar is using Homebrew:

```bash
brew cask install anybar
```

But you can also download the AnyBar binary directly from the [Releases](https://github.com/tonsky/AnyBar/releases) page.

## How to run

1. Download latest Netstatus binary from the [Releases](https://github.com/andrewkharook/netstatus/releases) page and put it in your PATH, e.g. `/usr/local/bin/`

2. Add either a cron job to run the netstatus repeatedly, or a [launchd](https://www.launchd.info) config, e.g.:

```xml
<!-- ~/Library/LaunchAgents/com.github.andrewkharook.netstatus.plist -->

<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key>
  <string>com.github.andrewkharook.netstatus</string>
  <key>Program</key>
  <string>/usr/local/bin/netstatus</string>
  <key>RunAtLoad</key>
  <true/>
  <key>StartInterval</key>
  <integer>60</integer>
</dict>
</plist>
```

3. Start Anybar and enjoy.
