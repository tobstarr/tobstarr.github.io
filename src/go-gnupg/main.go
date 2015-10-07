package gognupg

import (
	"bytes"
	"io"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func secring() (openpgp.EntityList, error) {
	f, err := os.Open(os.ExpandEnv("$HOME/.gnupg/secring.gpg"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return openpgp.ReadKeyRing(f)
}

func encrypt(b []byte, doArmor bool) ([]byte, error) {
	sr, err := secring()
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}

	aw, err := armor.Encode(buf, openpgp.PublicKeyType, nil)
	if err != nil {
		return nil, err
	}
	w, err := openpgp.Encrypt(aw, sr, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if _, err := w.Write(b); err != nil {
		return nil, err
	}
	for _, c := range []io.Closer{w, aw} {
		if err := c.Close(); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
