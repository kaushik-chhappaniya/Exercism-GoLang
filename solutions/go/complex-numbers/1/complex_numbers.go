package complexnumbers
import 	"math/cmplx"

// Define the Number type here.
type Number struct {
    a float64
    b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
    return Number{
        a: n1.a + n2.a,
        b: n1.b + n2.b,
        }
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
        a: n1.a - n2.a,
        b: n1.b - n2.b,
        }
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
        	a: ((n1.a * n2.a) - (n1.b * n2.b)),
        	b: ((n1.a * n2.b) + (n1.b * n2.a)),
    	}
}

func (n Number) Times(factor float64) Number {
 return Number {
     a: factor * n.a,
     b: factor * n.b,
 }   
}

func (n1 Number) Divide(n2 Number) Number {
    cn1 := complex(n1.a, n1.b)
    cn2 := complex(n2.a, n2.b)
    division := cn1/cn2
    
	return Number {
        a: real(division),
        b: imag(division),
    }
}

func (n Number) Conjugate() Number {
	return Number{
        a: n.a,
        b: (-1*n.b),
    }
}

func (n Number) Abs() float64 {
	cn1 := complex(n.a, n.b)
    return cmplx.Abs(cn1)
}

func (n Number) Exp() Number {
	cn1 := complex(n.a, n.b)
    res := cmplx.Exp(cn1)
    return Number {
        a: real(res),
        b: imag(res),
    }
}
