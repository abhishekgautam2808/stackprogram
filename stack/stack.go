package stack

type operation interface {
	push()
	pop()
}

type Stack []int

func (s Stack) Push(v int) []int {
	l := len(s)
	if l > 10 || l == 10 {
		return s
	}
	return append(s, v)

}

func (s Stack) Pop() []int {

	l := len(s)
	if l == 0 {
		return s
	}
	return s[:l-1]
}
