// Time: 0ms (100%) | Memory: 3 (69.39%)
func isPalindrome(s string) bool {
  lower := strings.ToLower(s)
  length := len(lower)

  j := length-1
  i := 0
  x := []byte{} // trying bytes instead of comparing strings
  y := []byte{}
  for i < length && j>0{
    x[0] = lower[i]
    y[0] = lower[j]
    if x[0] == ' ' || !((x[0] >= 97 && x[0] <= 122) || (x[0] >= 48 && x[0] <= 57)) {
      i++
      continue
    }
    if y[0] == ' ' || !((y[0] >= 97 && y[0] <= 122) || (y[0] >= 48 && y[0] <= 57)) {
      j--
      continue
    }

    if x[0] != y[0]{
        return false
    }
    i++
    j--
  } 

  return true
}
