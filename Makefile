.PHONY: default
default:
	echo "Please Specify The Packaged APP ... "

.PHONY: darwin_amd64
darwin_amd64:
	wails build --platform darwin/amd64
	mv -f ./build/bin/smart-engine.app ./build/bin/smart-engine-darwin-amd64.app

.PHONY: darwin_arm64
darwin_arm64:
	wails build --platform darwin/arm64
	mv -f ./build/bin/smart-engine.app ./build/bin/smart-engine-darwin-arm64.app

.PHONY: darwin_universal
darwin_universal:
	wails build --platform darwin/universal
	mv -f ./build/bin/smart-engine.app ./build/bin/smart-engine-darwin-universal.app

.PHONY: windows_amd64
windows_amd64:
	wails build --platform windows/amd64 -skipbindings -o smart-engine-windows-amd64.exe

.PHONY: windows_arm64
windows_arm64:
	wails build --platform windows/amd64 -skipbindings -o smart-engine-windows-arm64.exe

.PHONY: linux_amd64
linux_amd64:
	wails build --clean --platform linux/amd64 -o smart-engine-linux-amd64.exe

.PHONY: linux_arm64
linux_arm64:
	wails build --clean --platform linux/amd64 -o smart-eneine-linux-arm64.exe

.PHONY: make_all
make_all:
	make darwin_amd64
	make darwin_arm64
	make darwin_universal
	make windows_amd64
	make windows_arm64
	make linux_amd64
	make linux_arm64
