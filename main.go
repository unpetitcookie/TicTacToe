package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"github.com/fatih/color"
)

var board [9]string
var currentPlayer string

func main() {
	clearScreen()
	initializeBoard()
	currentPlayer = "X"
	color.Cyan("Bienvenue dans le jeu du morpion!\n")

	for {
		printBoard()
		position := getPlayerInput()
		makeMove(position)

		if winConditions() {
			printBoard()
			winner := color.New(color.FgGreen, color.Bold).Sprintf("Le joueur %s a gagné!", currentPlayer)
			fmt.Println(winner)
			break
		}

		if checkDraw() {
			printBoard()
			fmt.Println("Egalité!")
			break
		}

		switchPlayer()
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func initializeBoard() {
	for i := 0; i < len(board); i++ {
		board[i] = strconv.Itoa(i + 1)
	}
}

func printBoard() {
	fmt.Println("-------------")
	for i := 0; i < len(board); i += 3 {
		fmt.Printf("| %s | %s | %s |\n", formatCell(board[i]), formatCell(board[i+1]), formatCell(board[i+2]))
		fmt.Println("-------------")
	}
}

func formatCell(cell string) string {
	if cell == "X" {
		return color.RedString(cell)
	} else if cell == "O" {
		return color.BlueString(cell)
	}
	return cell
}

func getPlayerInput() int {
	var position int
	for {
		fmt.Printf("Joueur %s, entrez une position (1-9) : ", currentPlayer)
		fmt.Scanln(&position)
		clearScreen()
		if position < 1 || position > 9 {
			fmt.Println("Position invalide. Veuillez réessayer.")
		} else if board[position-1] != strconv.Itoa(position) {
			fmt.Println("Position déjà prise. Veuillez réessayer.")
		} else {
			break
		}
	}
	return position
}

func makeMove(position int) {
	board[position-1] = currentPlayer
}

func winConditions() bool {
	winConditions := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // lignes
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // colonnes
		{0, 4, 8}, {2, 4, 6},            // diagonales
	}

	for _, condition := range winConditions {
		if board[condition[0]] == currentPlayer && board[condition[1]] == currentPlayer && board[condition[2]] == currentPlayer {
			return true
		}
	}

	return false
}

func checkDraw() bool {
	for _, cell := range board {
		if cell != "X" && cell != "O" {
			return false
		}
	}
	return true
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}
