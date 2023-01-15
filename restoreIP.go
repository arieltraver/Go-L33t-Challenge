import ("strings"
"strconv")
//8(4) --> acceptable section lengths: 1, 2, 3
//6(3) --> acceptable section lengths: 1, 2, 3
    // 0_12_345
//1_1_3_3
//10(4) --> acceptable values: 2, 3
//instead of building the string itself, build a list of lengths
//once u hit the end, use the numbers to add the string to 
//at the very end convert all the numbers into actual strings.
//check 

func restoreIpAddresses(s string) []string {
    l := len(s)
    if (l > 12 || l < 4) { return []string{} }
    result := make([]string, 0, 32)
    var helper func(int, int, int)
    helper = func (lengths int, periodsLeft int, current int) {
        //lengths: a 3 digit number, representing section lengths
        if periodsLeft == 0 {
            num, _ := strconv.Atoi(s[current:l])
            if num <= 255 && (s[current] != '0' || l - current == 1) {
                //no more periods left and last chunk is valid
                a := lengths / 100
                b := lengths % 100 / 10 + a
                c := lengths % 10 + b
                result = append(result, s[:a] + "." + s[a:b] +
                "." + s[b:c] + "." + s[c:])
            }
        } else {
            groups := periodsLeft + 1
            lengthLeft := l - current
            max := lengthLeft / groups + lengthLeft - groups //max len of group
            min := lengthLeft - (periodsLeft * 3) //min len of group
            if (min < 1) { min = 1} //no less than 1 digit
            if (max > 3) { max = 3} //no more than 3 digits
            if s[current] == '0' {
                max = 1
            }
            for i := min; i <= max; i ++ {
                //if the next chunk has 3 digits, make sure it's below 255
                if (i < 3) {
                    helper(lengths * 10 + i, periodsLeft - 1, current + i)
                } else {
                    num, _ := strconv.Atoi(s[current:current+i])
                    if num <= 255 {
                        helper(lengths * 10 + i, periodsLeft - 1, current + i)
                    }
                }
            }
        }
    }
    helper(0, 3, 0)
    return result
}
