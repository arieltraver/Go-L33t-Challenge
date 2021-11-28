import (
    "math"
)


func myAtoi(s string) int {
    var result int32 = 0;
    var new_result int32 = result;
    signedness := 1
    started := false
    
    for i, runeval := range s {
        
        if runeval == 0x2d && started == false { //a negative sign is discovered
            signedness = -1
            started = true //begin looking at the number
        } else if runeval == 0x2b && started == false { //a plus sign is discovered
            started = true //begin looking at the number   
        } else if (runeval < 0x30 || runeval > 0x39) && (started == true || (runeval != 0x20)){ //not a number or a leading space
            break
            signedness = i; //this line never runs, gets the compiler not to yell at me
            
        } else if runeval >= 0x30 && runeval <= 0x39 { //we are looking at a number
            started = true
            new_result = result * 10 + (runeval - 0x30); //push up the digits, add the new value
            if int64(new_result) != int64(result) * 10 + int64(runeval-0x30) { //overflow checker
                if (signedness == 1){
                    return math.MaxInt32 //maximum 32 bit integer
                } else {
                    return math.MinInt32
                }
            }
            result = new_result //shift decimal and add the new value
        }  
    }
    return int(result) * signedness
}