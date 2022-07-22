package lntemp

import (
	"fmt"
	"io"
	"os"

	"github.com/lightningnetwork/lnd/lntest"
)

const (
	// NeutrinoBackendName is the name of the neutrino backend.
	NeutrinoBackendName = "neutrino"

	// TODO(yy): delete.
	DefaultTimeout = lntest.DefaultTimeout
)

// CopyFile copies the file src to dest.
func CopyFile(dest, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.Create(dest)
	if err != nil {
		return err
	}

	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}

	return d.Close()
}

// errNumNotMatched is a helper method to return a nicely formatted error.
func errNumNotMatched(name string, subject string,
	want, got, total, old int) error {

	return fmt.Errorf("%s: assert %s failed: want %d, got: %d, total: "+
		"%d, previously had: %d", name, subject, want, got, total, old)
}
