package main

import "fmt"

func permutation(s string) []string {
	var (
		result []string
		helper func([]byte)
		mapper = make(map[string]struct{})
		exist  struct{}
	)
	bytes := []byte(s)
	helper = func(sub []byte) {
		if len(sub) == 0 {
			tmpResult := string(bytes)
			if _, ok := mapper[tmpResult]; !ok {
				result = append(result, string(bytes))
				mapper[tmpResult] = exist
			}
			return
		}
		for i := 0; i < len(sub); i++ {
			sub[0], sub[i] = sub[i], sub[0]
			helper(sub[1:])
			sub[0], sub[i] = sub[i], sub[0]
		}
	}
	helper(bytes)
	return result

}

func main() {
	fmt.Println(permutation("a"))
	fmt.Println(permutation("aab"))
	fmt.Println(permutation("abc"))
}
