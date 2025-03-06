
// A function to generate a random number between 1 and 10
func randomNumber() int {
	return rand.Intn(10) + 1
}

// A function to solve the problem
func solveProblem(problem string) string {
	// Split the problem into two parts: the left-hand side and the right-hand side
	parts := strings.Split(problem, "=")

	// Convert the right-hand side of the equation to a number
	rightHandSide, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return "Invalid input"
	}

	// Use the randomNumber function to generate a random number between 1 and 10
	randomNumber := randomNumber()

	// Solve the problem by multiplying the random number with the right-hand side of the equation
	solution := randomNumber * rightHandSide

	// Convert the solution to a string and return it
	return strconv.Itoa(solution)
}