package main

const CACHE_SIZE int = 3

type Node struct {
	key int
	value int
	next, prev  *Node
}

var head *Node
var end *Node
var m map[int]*Node

func Set(key int, value int) {
	// TODO: add your code here!
	if old, ok := m[key]; ok {
		old.value = value
		remove(old)
		setHead(old)
	}else {
		created := new(Node)
		created.setNode(key, value)

		if(len(m) >= CACHE_SIZE){
			delete(m, end.key)
			remove(end)
			setHead(created)
		}else{
			setHead(created)
		}

		if m == nil {
			m = make(map[int]*Node)
		}
		m[key] = created
	}
}

func Get(key int) int {
	// TODO: add your code here!
	if n, ok := m[key]; ok {
		remove(n)
		setHead(n)
		return n.value

	}else{
		return -1
	}
}

func (node *Node) setNode(key int, value int ){
	node.key = key
	node.value = value
}

func remove(n *Node){
	if(n.prev != nil){
		n.prev.next = n.next
	}else{
		head = n.next
	}

	if(n.next != nil){
		n.next.prev = n.prev
	}else{
		end = n.prev
	}

}

func setHead(n *Node){
	n.next = head;
	n.prev = nil;

	if(head != nil){
		head.prev = n
	}

	head = n

	if(end == nil){
		end = head
	}

}