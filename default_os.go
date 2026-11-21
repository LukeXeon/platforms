//go:build !android

package platforms

import "runtime"

// defaultOS is the operating system of the platform's own images.
const defaultOS = runtime.GOOS
