package to

import "testing"
import "fmt"

func TestTO(test *testing.T) {
  var bad = test.Error

  q := Float("1.23")
  if q != 1.23 { bad("Float") }

  i := Int("123")
  if i != 123 { bad("Int") }

  t := Bool("true")
  if t != true { bad("Bool true") }

  f := Bool("false")
  if f != false { bad("Bool false") }

  d0 := Date("2016-01-14")
  d1 := fmt.Sprintf("%v", d0)
  if d1 != "2016-01-14 00:00:00 +0000 UTC" { bad("Date") }
}
