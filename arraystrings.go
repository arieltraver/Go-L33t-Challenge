/* so essentially we want to avoid wasting space by reassigning new strings
//options:
    -- store every letter in a static array. compare the arrays. what size though? not very space efficient
    -- store every letter in a linked list. compare the lists char by char. better because linked lists are mutable and space efficient
    -- no need for storage, just iterate through both lists char by char. make a 'next' function which helps you ignore */

//import("fmt")
//okay so the solution apparently was not all that fast? going to look into this one.

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    
    char1 := 0
    char2 := 0
    i1 := 0
    i2 := 0
    increment := func() int {
        //fmt.Println("i1 is", i1, ", c1 is", char1, ", i2 is", i2, ", c2 is", char2)
        if (char1 >= len(word1[i1]) - 1) {
            char1 = 0;
            i1 ++;
        } else {
            char1++;
        }
        if (char2 >= len(word2[i2]) - 1) {
            char2 = 0;
            i2 ++;
        } else {
            char2++;
        }
        if (i1 >= len(word1)) {
            return -1
        }
        if (i2 >= len(word2)) {
            return -1
        }
        return 0;
    }
    
    if (word1[0][0] != word2[0][0]) { return false; } //inc start
    
    for (increment() == 0) {
        if word1[i1][char1] != word2[i2][char2] {
            return false;
        }
    }
    //make sure length is equal -- both string lengths reached
    if (i1 < len(word1) || i2 < len(word2)) {
        return false;
    } else {
        return true;
    }
}
