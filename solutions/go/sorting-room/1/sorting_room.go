package sorting
import (
    "fmt"
    "strconv"
)
// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %0.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %0.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    
   _, ok := fnb.(FancyNumber)
    if !ok {
        return 0
    }
     n,_ := strconv.Atoi(fnb.Value())
    return n
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
     _, ok := fnb.(FancyNumber)
    if !ok {
        return fmt.Sprintf("This is a fancy box containing the number 0.0")
    }
     n,_ := strconv.Atoi(fnb.Value())
    return fmt.Sprintf("This is a fancy box containing the number %0.1f",float64(n))
	
}

// DescribeAnything should return a string describing whatever it contains.

func DescribeAnything(val any) string {
    switch v := val.(type) {
    case FancyNumberBox:
        return DescribeFancyNumberBox(v)
    case FancyNumber:
        return DescribeFancyNumberBox(v)
    case int:
        return DescribeNumber(float64(v))
    case  float64:
        return DescribeNumber(v)
    case NumberBox:
        return DescribeNumberBox(v)
    case string:
        return "Return to sender"
    default :
        return "Return to sender"
    }
}
