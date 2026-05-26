package binarysearch

func SearchInts(list []int, key int) int {
    low := 0
    high := len(list)-1
	for low <= high {
		// get middle number
        mid := (high+low)/2
        // if key found at middle return mid
        if list[mid] == key {
            return mid
        }
        // else if key is smaller than middle number then set high to middle index
        if key < list[mid] {
            high = mid - 1
        } else {
        // else if key is larger than middle number then set low to middle index
            low = mid + 1
        }
    } 
	return -1
}
