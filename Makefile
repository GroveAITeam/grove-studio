.PHONY: default
default:
	echo "Please Specify The Packaged APP ... "

.PHONY: darwin_universal
darwin_universal:
	wails build --platform darwin/universal
	mv -f ./build/bin/grove-studio.app ./build/bin/grove-studio-darwin-universal.app
