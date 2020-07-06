package datastructure

type Queue []interface{}

func NewQueue() *Queue{
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue)Enqueue(val interface{}){
	*q = append(*q, val)
}

func (q *Queue)Dequeue()interface{}{
	front := (*q)[0]
	*q = (*q)[1:]
	return front
}

func (q *Queue) Size() int {
	return len(*q)
}