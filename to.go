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

const VERSION string = "1.0.0"

var Panic bool = false
func Oops(a *string, t string) {
  msg := fmt.Sprintf("Could not parse \"%s\" as a %s.", *a, t)
  if Panic {
    panic(msg)
  } else {
    fmt.Fprintln(os.Stderr, msg)
    os.Exit(70) // Let's call this one a software error... somewhere.
  }
}

func Date(a *string) time.Time {
  t, e := time.Parse("2006-01-02", *a)
  if e!= nil { Oops(a, "Date") }
  return t
}

func Int(a *string) int {
  i, e :=  strconv.Atoi(*a)
  if e != nil { Oops(a, "Int") }
  return i
}

func Float(a *string) float64 {
  f, e := strconv.ParseFloat(*a, 64)
  if e != nil { Oops(a, "Float") }
  return f
}

func Bool(a *string) bool {
  b := false
  if *a == "true" {
    b = true
  } else if *a != "false" {
    Oops(a, "Bool")
  }
  return b
}
