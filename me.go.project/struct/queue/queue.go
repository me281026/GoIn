package queue


//定义一个队列,interface{}表示可以接收任何类型
type Queue []interface{}

//Push方法是往队列里面添加元素,这些方法都是以指针的形式接收队列
func (q *Queue) Push(v interface{}) {
	*q = append(*q,v)
}

//Pop方法是将队列里先进来的元素推出去
func (q *Queue) Pop() interface{} {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

//IsEmpty方法是判断队列是否为空
func (q *Queue) Isempty() interface{} {
	return len(*q) == 0
}
