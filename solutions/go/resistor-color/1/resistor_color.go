package resistorcolor

// Colors returns the list of all colors.
var ColorsArr = []string{"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
    }
func Colors() []string {
	
    return []string{"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
    }
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i,c := range ColorsArr {
        if c == color {
            return i
        }
    }
    return -1
}
