package lasagnamaster

// PreparationTime return time given then use it else just use 2
func PreparationTime(list []string, time int) int {
    if time ==0 {
        return len(list)*2
    }
    return len(list)*time
}
// Quantities if its noodles then add 50 if sauce then add 0.2 else do nothing
func Quantities(list []string) (int,float64) {
	var qu int
    var sa float64
    for _,str := range list{
    switch
    {
        case str == "noodles": 
            qu+=50
		case str == "sauce":
        	sa+=0.2
        }
}
    return qu, sa
    }

// AddSecretIngredient function - copy last element of friend list to last element of my list
func AddSecretIngredient(frdList, myList []string){
    myList[len(myList)-1]= frdList[len(frdList)-1]
}
// ScaleRecipe as given list has scaling for 2 proportions divide the current for given proportions
func ScaleRecipe(list []float64, q int) []float64{
    arr := []float64{}
    for _,i:=range(list){
        arr = append(arr, float64(q)*i/2)
    }
    return arr
}

