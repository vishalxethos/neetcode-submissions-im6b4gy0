// func isAnagram(s string, t string) bool {
//     if len(s) != len(t) {
//         return false
//     }

//     counts := make(map[rune]int)

//     for _, char := range s {
//         counts[char]++
//     }

//     for _, char := range t {
//         counts[char]--
//         if counts[char] < 0 {
//             return false
//         }
//     }

//     return true
// }

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sHash := make(map[byte]int)
    tHash := make(map[byte]int)

    for i := 0; i < len(s) ; i++ {
        sHash[s[i]]++
        tHash[t[i]]++
    }

    // fmt.Println(sHash,tHash)
    
    
    for key,value := range sHash {
        if tHash[key] != value {
            return false
        }

    }

    return true
}