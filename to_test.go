package to

import "testing"
import "fmt"

func TestTO(test *testing.T) {
  var bad = test.Error
  var v string

  v = "1.23"
  q := Float(&v)
  if q != 1.23 { bad("Float") }

  v = "123"
  i := Int(&v)
  if i != 123 { bad("Int") }

  v = "true"
  t := Bool(&v)
  if t != true { bad("Bool true") }

  v = "false"
  f := Bool(&v)
  if f != false { bad("Bool false") }

  v = "2016-01-14"
  d0 := Date(&v)
  d1 := fmt.Sprintf("%v", d0)
  if d1 != "2016-01-14 00:00:00 +0000 UTC" { bad("Date") }
}
