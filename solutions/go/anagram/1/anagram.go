package anagram
import "strings"
type FreqMap map[rune]int
func Detect(subject string, candidates []string) []string {
    subjectFreq := make(FreqMap)
	// Make a map for subject/target word frequency and count the frequency
    subjectFreq.CountFrequency(strings.ToLower(subject))
    
    result := []string{}
    for _, candidate := range candidates {
        // check if the candidate and the subject are same or not; if same skip the iteration
        cand := strings.ToLower(candidate)
            if cand == strings.ToLower(subject) {
                continue
        }
        // Create a map of candidate's frequency
        candidateFreq := make(FreqMap)
        candidateFreq.CountFrequency(cand)
        // Compare the candidate and target word frequency maps
        if subjectFreq.CompareFrequency(candidateFreq) {
            result = append(result, candidate)
        }
    }
    return result    
}

func (freq FreqMap) CountFrequency(word string) {
    //  Used the method to avoid copy of map. Create frequency map of given word.
    for _, r:= range word {
        freq[r]++
    }
}

func (targetFreq FreqMap) CompareFrequency(candidateFreq map[rune]int) bool {
    // Compare if length of the target and candidate are same or not if not then return false.
    if len(targetFreq) != len(candidateFreq){
        return false
    }
    // Check if the letter from the candidate word are present in the target word. 
	for k,v := range targetFreq {
        if candidateFreq[k] != v {
            return false
        }
    }
    return true
}