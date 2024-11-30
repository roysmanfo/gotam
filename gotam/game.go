package gotam

import "github.com/go-gl/glfw/v3.3/glfw"

type Game struct {
	root   *Node
	Window *glfw.Window
}

func CreateGameManager(window *glfw.Window) Game {
	root := &Node{
		parentNode:         nil,
		isActive:           true,
		hasBeenInitialized: false,
		childNodes:         make([]*Node, 0),
		Init:               func() {}}
	gameManager := Game{root: root, Window: window}

	return gameManager
}
