package to // import "github.com/carlosjhr64/to"

Just a small utility package to do string conversions which exits with
"software error" on failure.
const VERSION Version = "0.2.0.alpha"
var Panic bool = false
func Bool(a string) bool
func Date(a string) time.Time
func Float(a string) float64
func Int(a string) int
func MNBC(v string) (int, int, int, string)
type Version string
