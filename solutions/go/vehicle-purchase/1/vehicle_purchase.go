package purchase
import (
    "slices"
    "fmt"
    )
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    list := []string{"truck","car"}
	// return kind == "car" || "truck"
    return slices.Contains(list, kind)
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    arr := []string{option1,option2}
    slices.Sort(arr)
	return arr[0] + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
        case age < 3.0:
        	fmt.Println("price:",originalPrice*80/100)
        	return originalPrice*80/100
        case age >= 3.0 && age <10.0 :
        fmt.Println("price:",originalPrice*70/100)
        	return originalPrice*70/100
        case age >= 10.0:
        fmt.Println("price:",originalPrice*50/100)
        	return originalPrice*50/100
        default:
        	return originalPrice*50/100
    }
}
