```go
func main() {

        var i int
        fmt.Println(i) // Output: 0
        for j := 0; j < 5; j++ {
                i++
        }
        fmt.Println(i) // Output: 5

        var k int
        fmt.Println(k) // Output: 0
        for j := 0; j < 5; j++ {
                k = k + 1
        }
        fmt.Println(k) // Output: 5

        // Corrected loop initialization to avoid confusion
        for l:=0; l< 5; l++{
                var m int
                m = m +1
                fmt.Println(m) //Output: 1 (for each iteration)
        }
}
```