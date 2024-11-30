package gotam

import "fmt"

// private final ArrayList<Node> childNodes;

// private boolean active;
// private boolean hasBeenInitialized;

type BaseNode interface {
	Setup()
	Init()
	AddChildNode()
	Update()
}

type Node struct {
	parentNode         *Node // =nil for the root node
	isActive           bool
	hasBeenInitialized bool
	childNodes         []*Node // all child nodes
	Init               func()
}

func (n *Node) GetParent() *Node {
	return n.parentNode
}

func (n *Node) GetIsActive() bool {
	return n.isActive
}

func (n *Node) Setup(Init func()) error {
	if !n.isActive || !n.parentNode.isActive {
		return fmt.Errorf("this node has not been activated")
	}

	if !n.hasBeenInitialized {
		n.Init = Init
		n.hasBeenInitialized = true
	}

	n.isActive = true

	// initialize all childrens
	for _, child := range n.childNodes {
		err := child.Setup(child.Init)
		if err != nil {
			return err
		}

	}

	return nil
}

// add a new child node to this node
// when this node will be deactivated, all its childred
// will be deactivated too
func (n *Node) AddChildNode(child *Node, Init func()) error {
	n.childNodes = append(n.childNodes, child)
	child.parentNode = n

	if n.isActive {
		return child.Setup(Init)
	}
	return nil
}
