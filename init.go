package k6setenv

import (
	"go.k6.io/k6/js/modules"
)

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/setenv", new(Environment))
	modules.Register("k6/x/pinger", new(LgcPinger))
}
