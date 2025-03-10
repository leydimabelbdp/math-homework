
package main
import ("math/rand" "time")
func main() {
    // Generate a random number between 1 and 100
    rand.Seed(time.Now().UnixNano())
    num := rand.Intn(100) + 1
    fmt.Println("Generated random number:", num)
}