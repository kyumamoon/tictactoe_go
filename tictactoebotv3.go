package main

import "fmt"

func showwinner(u string) {
	var winner string
	switch u {
	case "p":
		winner = "Player 1"
	case "b":
		winner = "Player 2"
	}
	fmt.Printf("Winner is %s!", winner)
}

func turns(t *int64, g *[9]int64, m *int64) {
	plotmove := func(a int64, u string) {
		if g[a-1] == 1 || g[a-1] == 2 {
			fmt.Printf("Invalid Spot, Try Again!\n")
			turns(t, g, m)
		} else {
			if u == "p" {
				g[a-1] = 1
			} else {
				g[a-1] = 2
			}
		}
		printgrid(g)
	}

	checkwinner := func(u string) {
		columns := [9]int{0, 3, 6, 1, 4, 7, 2, 5, 8}
		diag1 := [3]int64{0, 4, 8}
		diag2 := [3]int64{2, 4, 6}

		var v int
		var pos int
		var meter int
		var vv int64
		var winner int

		if u == "p" {
			vv = 1
		} else {
			vv = 2
		}

		for i := 0; i < 2; i++ {
			for i := 0; i < 9; i++ {
				if v == 0 {
					if g[i] == vv {
						meter++
					}
				} else {
					if g[columns[i]] == vv {
						meter++
					}
				}
				pos++
				if pos == 3 {
					if meter == 3 {
						v = 0
						showwinner(u)
						winner++
						break
					} else {
						meter = 0
						pos = 0
					}
				}
			}
			v++
		}
		if v == 2 {
			for i := 0; i < 2; i++ {
				for i := 0; i < 3; i++ {
					if v == 2 {
						if diag1[i] == vv {
							meter++
						}
					} else if v == 3 {
						if diag2[i] == vv {
							meter++
						}
					}
					pos++
					if pos == 3 {
						if meter == 3 {
							v = 0
							showwinner(u)
							winner++
							break
						} else {
							meter = 0
							pos = 0
						}
					}
				}
				v++
			}
		}

		if winner == 0 {
			if *t == 0 {
				*t = 1
			} else {
				*t = 0
			}
			if *m != 8 {
				*m++
				fmt.Printf("TURNS: %d", *m)
				turns(t, g, m)
			} else {
				fmt.Printf("TIE!")
			}
		}
	}
	validate := func(answ int64, user string) {
		if answ > 9 || answ < 1 {
			fmt.Printf("Invalid Spot, Try Again!\n")
			turns(t, g, m)
		} else {
			plotmove(answ, user)
			checkwinner(user)
		}
	}

	usersturn := func(t *int64, user string) {
		fmt.Printf("\nPlayer %d's Turn\n\nType your move as a number (EX: 1 will be the top left corner of the grid).\n", *t+1)
		var answer int64
		fmt.Scanln(&answer)
		validate(answer, user)
	}

	if *t == 0 {
		usersturn(t, "p")
	} else {
		usersturn(t, "b")
	}
}

func printgrid(g *[9]int64) {
	p := 0
	for i := 0; i < 3; i++ {
		fmt.Printf("%d | %d | %d\n", g[p], g[p+1], g[p+2]) // 0,1,2
		p += 3
	}
}

func start() {
	grid := [9]int64{}
	var turn int64 = 0 // 0: player, 1: bot
	var maxturns int64

	fmt.Printf("Tic Tac Toe Golang\n")
	printgrid(&grid)
	turns(&turn, &grid, &maxturns)
}

func main() {
	start()
}
