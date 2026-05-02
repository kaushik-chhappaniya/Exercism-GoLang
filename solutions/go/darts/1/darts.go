package darts

func Score(x, y float64) int {
    var sq float64 =(x*x)+(y*y) 
    switch {
        case ((x*x)+(y*y))> (100) :{ 
            return 0
        }
        case sq <= 1 && (x >= -1 && x <= 1) && (y >= -1 && y <= 1)  :{
            return 10
        }
        case sq <= 25 && (x >= -5 && x <= 5) && (y >= -5 && y <= 5) :{
            return 5
        }
        case sq <= 100 && (x >= -10 && x <= 10) && (y >= -10 && y <= 10) : {
            return 1
        }
        default:
        return 0
    }
}
