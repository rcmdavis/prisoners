package main

// Play Iterated Prisoner's Dilemma and calculate payoff
func playIPD(agentStrategy, opponentStrategy map[string]string, rounds int) int {
	agentHistory := []string{}
	opponentHistory := []string{}
	agentScore := 0

	for i := 0; i < rounds; i++ {
		agentMove := getNextMove(agentStrategy, opponentHistory)
		opponentMove := getNextMove(opponentStrategy, agentHistory)

		agentScore += calculatePayoff(agentMove, opponentMove)

		agentHistory = append(agentHistory, agentMove)
		opponentHistory = append(opponentHistory, opponentMove)
	}

	return agentScore
}

// Function to get the next move based on the strategy and opponent's history
func getNextMove(strategy map[string]string, opponentHistory []string) string {
	if len(opponentHistory) == 0 {
		return strategy[""] // Use the first move defined in the strategy
	}
	lastMove := opponentHistory[len(opponentHistory)-1]
	if move, exists := strategy[lastMove]; exists {
		return move
	}
	return strategy[""] // Default to the first move if no match is found
}

// Calculate payoff for a single round
func calculatePayoff(agentMove, opponentMove string) int {
	if agentMove == "C" && opponentMove == "C" {
		return 3 // Mutual cooperation
	} else if agentMove == "C" && opponentMove == "D" {
		return 0 // Agent cooperates, opponent defects
	} else if agentMove == "D" && opponentMove == "C" {
		return 5 // Agent defects, opponent cooperates
	} else {
		return 1 // Mutual defection
	}
}
