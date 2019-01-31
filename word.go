package words

import "strings"
import "fmt"

func WordCount(s string) map[string]int { //HL
	fmt.Println("Khanaphoz")
	r := strings.Replace(s, ".", "", -1)
	r = strings.Replace(r, ",", "", -1)
	r = strings.ToLower(r)
	counts := map[string]int{}
	for _, word := range strings.Fields(r) {
		counts[word]++
	}
	return counts

	//	return map[string]int{"x":1}
}
