package main
import "fmt"
type Stack interface{
	push(item int)
	pop()int 
	peek()
	isEmpty()bool
} 

func (s *stack)isEmpty()bool{
	if s.head == nil{
		return true
	}
	return false
}
type node struct{
	data int 
	next *node
}
type stack struct{
	head *node
}
func (s *stack)push(item int){
	newNode := node{
		data: item,
		next: nil,
	}
	if s.head == nil{
		s.head = &newNode
	}else{
		temp := s.head
		newNode.next = temp
		s.head = &newNode
	}
	//printStack(s)
	fmt.Printf("PUSHED: %d\n",newNode.data)
	printStack(s)
	return 
}

func (s *stack)pop()int{
	if s.isEmpty(){
		fmt.Println("STACK UNDERFLOW!")
		return 0
	}
	temp := s.head
	s.head = s.head.next
	return temp.data
}

func (s *stack)peek(){
	temp := s.head
	for temp.next != nil{
		temp = temp.next
	}
	fmt.Println(temp.data)
}

func printStack(s *stack){
	if s.head == nil{
		return
	}
	temp := s.head
	for temp != nil{
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func new()Stack{
	return &stack{
		head: nil,
	}
}
func main(){
	st := new()
	st.push(1)
	st.push(2)
	st.push(3)
	st.push(4)
	st.push(5)
	fmt.Println("POP:",st.pop())
	fmt.Println("POP:",st.pop())
	fmt.Println("POP:",st.pop())
	fmt.Println("POP:",st.pop())
	fmt.Println("POP:",st.pop())
	fmt.Println("POP:",st.pop())
}