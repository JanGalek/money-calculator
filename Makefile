.PHONY: build

build-linux:
	fyne-cross linux

build-android:
	fyne-cross android

build-mac:
	fyne-cross darwin

build-windows:
	fyne-cross windows

build:
	build-linux
	build-android
	build-mac
	build-windows