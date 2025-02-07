package binarytree

type node struct {
	value int
	right *node
	left  *node
}

type tree struct {
	root *node
	elementsnumber int
}

func New() *tree {
	t := &tree{}

	return t
}

func (t *tree) Insert(newValue int) {
	defer func() {
		t.elementsnumber++
	}()
	
	if t.root == nil {
		t.root = &node{
			value: newValue,
			right: nil,
			left:  nil,
		}
		return
	}

	current := t.root

	for {
		if newValue >= current.value {
			if current.right == nil {
				current.right = &node{value: newValue}
				return
			}

			current = current.right

		} else { 
			if current.left == nil {
				current.left = &node{value: newValue}
				return
			}

			current = current.left
		}
	}
}

func (t *tree) Exists(newValue int) bool {
	if t.root == nil {
		return false
	}

	current := t.root

	for {
		if current == nil {
			return false
		}

		if newValue == current.value {
			return true
		}
		
		if newValue < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
}

func (t *tree) ElementsNumber() int {
	return t.elementsnumber
}

/* Допустим x это узел в двоичном дереве поиска,
если y является «левым поддеревом» для x,
то y.key < x.key. Если y является «правым поддеревом» для x, то y.key >= x.key
*/

/*
написать метод exists, который будет проверять наличие значения в дереве
написать тесты
*/