package romantoint

func RomanToInt(s string) (result int) {
  romanToInt := map[byte] int {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
  }

  for i := 0; i < len(s); i++ {
    val, _ := romanToInt[s[i]]

    if i == len(s)-1 {
      result+=val
    } else {
      nextVal, _ := romanToInt[s[i+1]]

      if val > nextVal {
        result+=val
      } else {
        val -= nextVal
        result += val
        i+=2
      }
    }
  }

  return

}