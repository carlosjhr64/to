/*
  Just a small utility package to do string conversions
  which exits with "software error" on failure.
*/
package to

import (
  "os"
  "fmt"
  "time"
  "strconv"
)

func oops(a, t string) {
  fmt.Fprintf(os.Stderr, "Could not parse \"%s\" as a %s.\n", a, t)
  os.Exit(70) // Let's call this one a software error... somewhere.
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
