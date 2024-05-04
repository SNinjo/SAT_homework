package queue

type Queue struct {
	slice []interface{}
}

func New() *Queue {
	q := new(Queue)
	q.slice = make([]interface{}, 0)
	return q
}

func (q Queue) IsEmpty() bool {
	return len(q.slice) == 0
}

func (q *Queue) Put(data interface{}) {
	q.slice = append(q.slice, data)
}

func (q *Queue) Get() interface{} {
	data := q.slice[0]
	q.slice = q.slice[1:]
	return data
}
