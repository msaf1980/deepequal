deepequal
=========

A version of Go's reflect.DeepEqual that includes reasons for failing inequality checks.

This is A simply copy of the DeepEqual function from the Go standard library, enhanced to give a reason for a comparison failure and interpeter float NaN == NaN as true.

TODO: Can't compare structs with unexported fileds (panic is a result). `reflect.DeepEqual` use for them some internal methods.

Usage:

```
import "github.com/msaf180/deepequal"

func main() {
    x := map[string]int{ "a":1, "b", 5}    
    y := map[string]int{ "a":1, "b", 8}    

    equal, reason := deepequal.Compare(x, y)
    if !equal {
        println(reason)
    }
}
```

