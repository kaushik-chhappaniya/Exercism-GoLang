package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	mp:=make(map[string]int)
    mp["quarter_of_a_dozen"]=3
mp["half_of_a_dozen"]=6
mp["dozen"]=12
mp["small_gross"]=120
mp["gross"]=144
mp["great_gross"]=1728
    return mp
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if val, ok := units[unit]; ok {
		bill[item] = bill[item]+val
        return true
    } else {
        return false
        }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	

    if val1, ok1 := units[unit]; !ok1 {
       	 return false
        } else {
        	if val, ok := bill[item]; !ok{
                return false
            } else {
           		 if (val - val1 < 0) {
                	return false
            	} else if (val-val1==0) {
                	delete(bill,item)
                		return true
            	} else {
                    bill[item] = val-val1
                    return true
        		}
    		}
    	}     
    }


// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if q,ok := bill[item]; ok {
        return q,true
    } else {
    return 0,ok
	}	
}
