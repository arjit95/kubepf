package iostreams

import "io"

// IOStreams provides the standard names for iostreams.
type IOStreams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
	Logger io.Writer
}
