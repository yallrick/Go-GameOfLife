package main

import(
	"fmt"
)


// grand noir ⬛
// grand blanc ⬜
// noir ■
// blanc □

const(
	char_on  = rune('⬛')
	char_off = rune('⬜')
)

var game [][]bool

func main() {
	fmt.Println("Bonjour, c'est le jeu de la vie")

	game = gameConstructor(5, 5);

	initialize()


	for{
		draw()
		update()
		// TODO wait for 1 secondes
	}

}


func gameConstructor(n,m int) (g [][]bool){
	for i := 0; i < n ; i++ {
		var line []bool
		for j := 0; j < m ; j++  {
			line = append(line, false)
		}
		g = append(g, line)
	}
	return
}

func initialize(){
	// TODO initialize
}

func draw() {
	for i := 0; i < len(game); i++ {
		for j := 0; j < len(game[i]) ; j++  {
			fmt.Printf("%c", chooseCharCase(game[i][j]))
		}
		fmt.Print("\n")
	}
}

func chooseCharCase(isAlive bool) rune{
	if isAlive{
		return char_on
	}
	return char_off
}

func update(){
	for i := 0; i < len(game); i++ {
		for j := 0; j < len(game[i]) ; j++  {
			game[i][j] = destiny(i, j)
		}
	}
}


func destiny(i,j int) (d bool){
	c := countNeibourhoods(i,j)
	d = false
	if game[i][j] {
		if c == 2 || c == 3 {
			d = true
		}
	}else{
		if c == 3{
			d = true
		}
	}
	return
}

func countNeibourhoods(k,l int)(c int)  {
	c = 0

	for i := k - 1; i < k + 1; i++ {
		for j := l - 1; j < l + 1 ; j++  {
			// TODO countNeibourhoods
		}
	}

	return
}