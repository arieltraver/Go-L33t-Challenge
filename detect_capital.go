//import ("strings")

func detectCapitalUse(word string) bool {
    if len(word) <= 1 { return true; }
    capitals := 0
    lowercase := 0
    for _, char := range(word[1:]) {
        if char >= 97 {
            lowercase += 1
            if capitals > 0 { return false }
        } else {
            capitals += 1
            if lowercase > 0 { return false }
        }
    }
    return (word[0] < 97 || lowercase > 0)
}
