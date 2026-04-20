package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards(nums ...int) []int {
	var arr = []int{2,6,9}
    for _,v:=range(nums){
        if !(v == 2 || v == 6 || v == 9) {
        	arr = append(arr, v)
        }
    }
    return arr
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if index >= len(slice) || index < 0{
    	return -1
    }
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || (index < 0){
        return append(slice, value)
    } else {
        slice[index] = value
        return slice
    }
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    if values != nil {
		return append(values, slice...)
    }
    return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if index < 0 || index>= len(slice) {
        return slice
    }
	return append(slice[:index], slice[index+1:]...)
}
