package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	ans := initial
	for i := 0; i < len(s); i++ {
		ans = fn(ans, s[i])
	}
	return ans
}
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	ans := initial
	for i := len(s)-1;i>=0; i-- {
		ans = fn(s[i], ans)
	}
	return ans
}

func (s IntList) Filter(fn func(int) bool) IntList {
    ans := make(IntList, 0)
    for i:=0; i<len(s); i++ {
        if fn(s[i]){
        	ans =append(ans,s[i])            
        }
    }
    return ans
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	ans := make(IntList, 0 )
    for i:=0; i< len(s); i++{
    	ans = append(ans, fn(s[i]))
    }
    return ans
}

func (s IntList) Reverse() IntList {
	ans := make(IntList, 0)
    for i:=0; i<len(s); i++ {
        ans = append(ans, s[len(s)-i-1])
    }
    return ans
}

func (s IntList) Append(lst IntList) IntList {
    ans :=make(IntList, 0)
    ans = append(s, lst...)
    return ans
}

func (s IntList) Concat(lists []IntList) IntList {
    ans :=make(IntList, 0)
    ans = append(ans,s...)
	for i:=0; i<len(lists); i++{
        ans = append(ans, lists[i]...)
    }
    return ans
}
