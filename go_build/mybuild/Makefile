.PHONY: clean

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
VERSION=1.0.0

TARGET=mybuild

	ifeq ("windows", "$(GOOS)")
		TARGET=mybuild.exe
	endif

${TARGET}:
	env GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-X main.version=$(VERSION)" -o ${TARGET} .

gopher:
	env GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-X main.version=$(VERSION)" -tags=gopher -o ${TARGET} .

# 以下是測試執行檔。請自行在各平台測試。
# 如要在 MacOS 執行 windows 程式，可安裝 wine。
# brew install --cask wine-stable
# 使用 docker 執行 linux 程式。

darwin:${TARGET}
	./${TARGET}

windows:${TARGET}
	wine64 mybuild.exe

linux:${TARGET}
	docker run -v `pwd`:/myapp ubuntu:18.04 /myapp/mybuild

clean:
	-rm ${TARGET}
	-rm ${TARGET}.exe