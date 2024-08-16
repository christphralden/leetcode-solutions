// Time: 11ms (92.16%) | Memory: 6.72 (80.39%)
package main

import "fmt"

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	return checkPalindrome(s, left, right, true)
}

func checkPalindrome(s string, left, right int, safe bool) bool {
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			if safe {
				if checkPalindrome(s, left, right-1, false) {
					return true
				}
				if checkPalindrome(s, left+1, right, false) {
					return true
				}
				return false
			}
			return false
		}
	}
	return true
}

func main(){
  // tc1 := "abca"  
  // tc2 := "abc"  
  tc3 := "ebcbbececabbacecbbcbe"
  tc4 := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"  

  // fmt.Println(validPalindrome(tc1))
  // fmt.Println(validPalindrome(tc2))
  fmt.Println(validPalindrome(tc3))
  fmt.Println(validPalindrome(tc4))
}
