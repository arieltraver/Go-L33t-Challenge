func convert(s string, numRows int) string {
    zigStep := numRows + numRows - 2 //going up the ladder
    if zigStep == 0 { zigStep = 1 } //edge case with one row
    sideStep := 0
    newIndex := 0 //index of the string we're building
    oldIndex := 0 //index of the original string
    zigStepping := true
    rows := 0
    result := make([]byte, len(s)) //using a byte slice here
    for rows < numRows && oldIndex < len(s) && newIndex < len(s) {
        for oldIndex < len(s) && newIndex < len(s) {
            result[newIndex] = s[oldIndex] //fill in the char
            if(zigStep != 0 && zigStepping || sideStep == 0) {
                oldIndex += zigStep //step to the right
                zigStepping = false
            } else {
                oldIndex += sideStep //go down the zigzag
                zigStepping = true;
            }
            newIndex++;
        }
        rows++ //onto the next row
        oldIndex = rows //starting point
        zigStepping = true
        sideStep += 2
        zigStep -= 2
    }
    return string(result)
}