# Auto-clicker

It's a small and simple auto-mouse-mover for every distribution.

## Installation

### For MacOS:

```bash
Xcode Command Line Tools (And Privacy setting: #277 )

xcode-select --install
```
### For Windows:

```bash
MinGW-w64 (Use recommended)

Download the Mingw, then set system environment variables C:\mingw64\bin to the Path. Set environment variables to run GCC from command line.

Or the other GCC (But you should compile the "libpng" with yourself when use the bitmap.)
```
### For everything else:

```bash
GCC

X11 with the XTest extension (the Xtst library)


"Bitmap":

libpng (Just used by bitmap)

"Event":

xcb, xkb, libxkbcommon

"Clipboard":

xsel xclip
```
### Ubuntu:

```bash
# gcc
sudo apt install gcc libc6-dev

# x11
sudo apt install libx11-dev xorg-dev libxtst-dev

# Bitmap
sudo apt install libpng++-dev

# Hook
sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev libxkbcommon-dev

# Clipboard
sudo apt install xsel xclip
```
### Fedora:

```bash
sudo dnf install libXtst-devel

# Bitmap
sudo dnf install libpng-devel

# Hook
sudo dnf install libxkbcommon-devel libxkbcommon-x11-devel xorg-x11-xkb-utils-devel

# Clipboard
sudo dnf install xsel xclip
```
