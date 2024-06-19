package list

type Node struct {
    Key  string
    Value interface{}
    Prev *Node
    Next *Node
}

type DoublyLinkedList struct {
    head *Node
    tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
    return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) MoveToFront(node *Node) {
    if dll.head == node {
        return
    }
    dll.remove(node)
    dll.PushFront(node)
}

func (dll *DoublyLinkedList) PushFront(node *Node) {
    if dll.head == nil {
        dll.head = node
        dll.tail = node
        return
    }
    node.Next = dll.head
    dll.head.Prev = node
    dll.head = node
}

func (dll *DoublyLinkedList) RemoveTail() *Node {
    if dll.tail == nil {
        return nil
    }
    tail := dll.tail
    dll.remove(tail)
    return tail
}

func (dll *DoublyLinkedList) remove(node *Node) {
    if node.Prev != nil {
        node.Prev.Next = node.Next
    } else {
        dll.head = node.Next
    }
    if node.Next != nil {
        node.Next.Prev = node.Prev
    } else {
        dll.tail = node.Prev
    }
}
