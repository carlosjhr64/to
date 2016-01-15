/*
  Just a small utility package to do string conversions
  which exits with "software error" on failure.
*/
package to

import (
  "os"
  "fmt"
  "time"
  "strings"
  "strconv"
)

type Version string
const VERSION Version = "0.1.0.alpha"

var PANIC bool = false
func oops(a, t string) {
  msg := fmt.Sprintf("Could not parse \"%s\" as a %s.", a, t)
  if PANIC {
    panic(msg)
  } else {
    fmt.Fprintln(os.Stderr, msg)
    os.Exit(70) // Let's call this one a software error... somewhere.
  }
}

func MNBC(v string) (int, int, int, string) {
  a := strings.SplitN(v, ".", 4)
  if len(a) < 3 { oops(v, "MNBC Version") }
  major := Int(a[0])
  minor := Int(a[1])
  build := Int(a[2])
  note := "" // or comment
  if len(a) == 4 { note = a[3] }
  return major, minor, build, note
}

func (v Version) MNBC() (int, int, int, string) {
  return MNBC(string(v))
}

func Date(a string) time.Time {
  t, e := time.Parse("2006-01-02", a)
  if e!= nil { oops(a, "Date") }
  return t
}

func Int(a string) int {
  i, e :=  strconv.Atoi(a)
  if e != nil { oops(a, "Int") }
  return i
}

func Float(a string) float64 {
  f, e := strconv.ParseFloat(a, 64)
  if e != nil { oops(a, "Float") }
  return f
}

func Bool(a string) bool {
  b := false
  if a == "true" {
    b = true
  } else if a != "false" {
    oops(a, "Bool")
  }
  return b
}
