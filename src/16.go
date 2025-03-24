func main() {
    // Define a variable to hold your mathematical expression
    var expression string = "5 + 3 * 2 - 4 / 2"

    // Calculate the result of the expression
    var result float64
    result, _ = calculate(expression)

    // Print the result
    fmt.Printf("The result of the expression is: %f\n", result)
}

// Function to calculate a mathematical expression given as a string
func calculate(expression string) (float64, error) {
    // Split the expression into individual components
    exprList := strings.Fields(expression)

    var num1 float64 = 0.0
    for _, e := range exprList {
        if strings.Contains(e, "+") || strings.Contains(e, "-") {
            break
        }
        num1 += calculateComponent(e)
    }

    return num1, nil

}

// Function to calculate a mathematical component given as a string
func calculateComponent(expression string) float64 {
    switch e := expression; e {
    case "+":
        return 5.0 + 3.0
    case "-":
        return 2.0 - 4.0 / 2.0
    default:
        return 1.0 // Return a value of 1 for any other component
    }
}
