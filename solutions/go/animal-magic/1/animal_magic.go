package chance
import "math/rand"
// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
 return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    min := 0.0
    max:= 12.0
    result := min + rand.Float64()*(max-min)
	 return result
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    var arr = []string{"ant","beaver","cat","dog","elephant","fox","giraffe","hedgehog"}
	rand.Shuffle(len(arr), func (i,j int) {
        arr[i], arr[j] = arr[j], arr[i]
    })
    return arr
}
