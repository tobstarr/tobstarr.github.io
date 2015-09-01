# Setup Go

		tmp=$(mktemp -d /tmp/goXXXXXX)
		cd $tmp && curl -SsfL https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz | tar xfz - --strip=1
		mv $tmp ~/go
