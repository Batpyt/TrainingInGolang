package main

import "fmt"

func main() {
	q := NewQueue()
	q.Append(1)
	fmt.Println(q)
	fmt.Println(q.Out())
	fmt.Println(q)
	q.Append(2)
	fmt.Println(q)
	q.Append(3)
	fmt.Println(q)
	fmt.Println(q.Out())
	fmt.Println(q)
	fmt.Println(q.Out())
	fmt.Println(q.Out())
	fmt.Println(q.Out())
}

type Queue struct {
	in, out []int
}

func NewQueue() Queue {
	return Queue{
		in:  nil,
		out: nil,
	}
}

func (q *Queue) Append(num int) {
	q.in = append(q.in, num)
}

//123
//in: 123
//out: 321 in:[]
//Out() out:32 in[]

func (q *Queue) Out() int {
	if len(q.out) <= 0 {
		if len(q.in) <= 0 {
			return -1
		}
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}

	v := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return v
}
