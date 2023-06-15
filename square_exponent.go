//exponent by squaring in Go, as described in Wikipedia
//except instead of explicit tail recursion it's iteration
//deals with integers only, does NOT include negative exponents.

func int_exp(n int, exp int) int {
    y := 1
    for exp > 1 {
        if (exp % 2 == 0) {
            n *= n
            exp /= 2
        } else {
            y = y * n
            n = n * n
            exp = (exp - 1) / 2
        }
    }
    return n * y
}


