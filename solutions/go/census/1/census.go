// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
        Name:name,
        Age:age,
        Address:address,
    }
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	// check map is not blank using len() func 
	if r.Name=="" || r.Address == nil || len(r.Address) ==0 {
        return false 
    }
    // check if street key is not blank
    if _,ok:=r.Address["street"];!ok {
    	return false
    }
		// Check any other key in map isnt blank
        for _,v :=range (r.Address) {
            if v == "" {
                return false
            }
        }
    
    return true
    
}

// Delete deletes a resident's information. Assign all to default values
func (r *Resident) Delete() {
	r.Name = ""
    r.Age = 0
    r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	sum :=0
    if len(residents) <=0{
        return 0
    }
    for _,r :=range residents{
        if r.HasRequiredInfo(){
            sum++
        }
    }
    return sum
}
