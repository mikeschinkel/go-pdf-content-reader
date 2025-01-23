package pdf

import (
	"io"
)

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error("ERROR: Failed to close", "error", "err")
	}
}
