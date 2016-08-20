NGROK_OSX_64=https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-darwin-amd64.zip
NGROK_WIN_64=https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-windows-amd64.zip
NGROK_LINIX_64=https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip
NGROK_LINUX_ARM=https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-arm.zip
NGROK_FREEBSD_64=https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-freebsd-amd64.zip

GOTTY_LINUX_64=https://github.com/yudai/gotty/releases/download/pre-release/gotty_linux_amd64.tar.gz

linux:
	rm -f data/*
	wget --progress=bar ${NGROK_LINIX_64}
	unzip ngrok-stable-linux-amd64.zip
	rm ngrok-stable-linux-amd64.zip
	wget --progress=bar ${GOTTY_LINUX_64}
	tar -zxf gotty_linux_amd64.tar.gz
	rm gotty_linux_amd64.tar.gz
	go-bindata -tags linux ngrok gotty
	go build

clean:
	rm ngrok gotty
	go clean
