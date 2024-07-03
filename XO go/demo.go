package main

import "fmt"

func main() {
	Board := [3][3]string{}
	moves := 0
	player := "x"
	for {
		fmt.Println("player", player)
		var row int
		fmt.Println("please enter row number 0,1 or2")
		fmt.Scanln(&row)
		if row > 2 {
			fmt.Print("invalid input")
			continue
		}
		var column int
		fmt.Println("please enter column number0,1 or2")
		fmt.Scanln(&column)
		fmt.Println("row", row, "column", column)
		if column > 2 {
			fmt.Print("invalid input")
			continue
		}

		if Board[row][column] == "" {
			Board[row][column] = player
		} else {
			fmt.Print("you enter invalid position")
			continue //علشان مطلعش برا خالص ارجع اكمل لو المكان دا من فاضي
		}

		fmt.Println(Board[0])
		fmt.Println(Board[1])
		fmt.Println(Board[2])
		//بمشي علي الاراي اشوف حد كسب ولا لا
		for i := 0; i < 3; i++ {
			if Board[i][0] == player && Board[i][0] == Board[i][1] && Board[i][0] == Board[i][2] {
				fmt.Println("winnerx")
				return //هتعمل خروج خال من المين
			}

			if Board[0][i] == player && Board[0][i] == Board[1][i] && Board[0][i] == Board[2][i] {
				fmt.Println("winner")
				return //هتعمل خروج خال من المين

			}
		}
		if Board[0][0] == player && Board[0][0] == Board[1][1] && Board[0][0] == Board[2][2] || Board[0][2] == player &&
			Board[0][2] == Board[1][1] && Board[0][2] == Board[2][0] {
			fmt.Println("winner")
			return
		}
		if moves == 9 {
			fmt.Println("it's a draw")
			return
		}
		if player == "x" {
			player = "o"

		} else {
			player = "x"
		}
	}

}
