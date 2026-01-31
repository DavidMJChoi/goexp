package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

var cursorPos int = 0

const ACTION_CURSOR_UP int = 1
const ACTION_CURSOR_DOWN int = 2
const ACTION_CURSOR_RIGHT int = 3
const ACTION_CURSOR_LEFT int = 4

const ACTION_PLACE_PIECE int = 10

func main() {

	pieces := []string{
		"   ", "   ", "   ",
		"   ", "   ", "   ",
		"   ", "   ", "   ",
	}

	fmt.Print("\033[2J\033[H")
	drawBoard(0, pieces)
	for {

		action, exit := getKeyBoardInput()

		if exit {
			break
		}

		fmt.Print("\033[2J\033[H")
		drawBoard(action, pieces)
	}

	// time.Sleep(time.Second)

	// fmt.Print("\033[2J\033[H")
}

func getKeyBoardInput() (int, bool) {
	// 切换到原始输入模式
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var buf [3]byte
	for {
		n, err := os.Stdin.Read(buf[:])
		if err != nil {
			break
		}

		// 方向键通常是3字节序列：ESC [ A/B/C/D
		if n == 3 && buf[0] == 27 && buf[1] == '[' {
			switch buf[2] {
			case 'A':
				return ACTION_CURSOR_UP, false
			case 'B':
				return ACTION_CURSOR_DOWN, false
			case 'C':
				return ACTION_CURSOR_RIGHT, false
			case 'D':
				return ACTION_CURSOR_LEFT, false
			}
		} else if n == 1 {
			// ^C
			switch buf[0] {
			case 3:
				return -1, true
			case 32:
				return ACTION_PLACE_PIECE, false
			}

			fmt.Printf("普通键: %c\n", buf[0])
		} else {
			continue
		}
	}

	return -1, true
}

func drawBoard(action int, pieces []string) {
	/*

	    o │ o │ o
	   ───┼───┼───
	    o │ o │ o
	   ───┼───┼───
	    o │ o │ o

	*/

	switch action {
	case ACTION_CURSOR_UP:
		if cursorPos >= 3 {
			cursorPos -= 3
		} else {
			cursorPos += 6
		}

	case ACTION_CURSOR_DOWN:
		if cursorPos < 6 {
			cursorPos += 3
		} else {
			cursorPos -= 6
		}

	case ACTION_CURSOR_RIGHT:
		if cursorPos%3 == 2 {
			cursorPos -= 2
		} else {
			cursorPos += 1
		}

	case ACTION_CURSOR_LEFT:
		if cursorPos%3 == 0 {
			cursorPos += 2
		} else {
			cursorPos -= 1
		}
	case ACTION_PLACE_PIECE:
		pieces[cursorPos] = " O "

	default:
	}

	pieces[cursorPos] = "→" + pieces[cursorPos][1:]

	row1 := fmt.Sprintf("\n%s│%s│%s\n", pieces[0], pieces[1], pieces[2])
	row2 := fmt.Sprintf("%s│%s│%s\n", pieces[3], pieces[4], pieces[5])
	row3 := fmt.Sprintf("%s│%s│%s\n\n", pieces[6], pieces[7], pieces[8])

	fmt.Print(row1)
	fmt.Println("───┼───┼───")
	fmt.Print(row2)
	fmt.Println("───┼───┼───")
	fmt.Print(row3)

	pieces[cursorPos] = " " + pieces[cursorPos][3:]
}
