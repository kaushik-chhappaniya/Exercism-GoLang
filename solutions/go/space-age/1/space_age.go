package spaceage

type Planet string

func Age(seconds float64, planet Planet) float64 {
	daySec := float64(365.25*24*60*60)
    
    yrs := seconds/daySec
    switch planet {
        case "Mercury":
        	return yrs/0.2408467
        case "Venus":
        	return yrs/0.61519726
        case "Earth":
        	return yrs/1.0000
        case "Mars":
        	return yrs/1.8808158
        case "Jupiter":
        	return yrs/11.862615
      	case "Saturn":
        	return yrs/29.447498  
        case "Uranus":
        	return yrs/84.016846
        case "Neptune":
        	return yrs/164.79132
        default:
        	return -1.00
    }
    
}

