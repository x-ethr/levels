package levels_test

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"

	"github.com/x-ethr/levels"
)

func Test(t *testing.T) {
	ctx := context.Background()

	t.Run("Trace", func(t *testing.T) {
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

		if strings.Contains("TRACE", output.String()) {
			t.Logf("Success: %s", output.String())
		}
	})
}
