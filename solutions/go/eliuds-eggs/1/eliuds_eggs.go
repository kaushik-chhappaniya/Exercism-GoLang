package eliudseggs

func EggCount(displayValue int) int {
	sum :=0
    for i:=displayValue; i>0;i--{
        sum += i%2
        i /= 2
    }
    return sum
}
