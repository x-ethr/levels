package levels_test

import (
	"bytes"
	"context"
	"log/slog"

	"github.com/x-ethr/levels"
)

func Example() {
	ctx := context.Background()
	level := levels.Trace
	slog.SetLogLoggerLevel(level)

	options := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			a = levels.Attributes(groups, a)

			return a
		},
		Level: level,
	}

	var output bytes.Buffer

	slog.SetDefault(slog.New(slog.NewTextHandler(&output, options)))

	slog.Log(ctx, levels.Trace, "Example Trace Log Message")
}
