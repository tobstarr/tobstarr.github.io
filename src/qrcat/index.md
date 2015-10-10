# qrcat

Send yourself texts from your computer to your smartphone.

## Usage

	$ echo sosolongbecauseitissecretpassword | qrcat
	$ cat credentials.file | qrcat


## Code

__Download__ [qrcat.go](qrcat.go)

{{ require "src/qrcat/qrcat.go" | code }}

## Code as QR Code

	curl -s tobstarr.com/qrcat/qrcat.go | qrcat

<img src="/qrcat/qrcat.png" />
