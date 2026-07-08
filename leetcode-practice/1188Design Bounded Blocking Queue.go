// 定义队列结构
type BoundedBlockingQueue struct {
    ch chan int
}

// 构造函数
func Constructor(capacity int) BoundedBlockingQueue {
    return BoundedBlockingQueue{
        ch : make(chan int, capacity),
    }
}

// 入队方法（如果是别人调用，且队列满了，底层的 channel 会自动让调用方的 goroutine 阻塞）
func (q *BoundedBlockingQueue) Enqueue(element int) {
    q.ch <- element
}

// 出队方法
func (q *BoundedBlockingQueue) Dequeue() int {
    return <- q.ch 
}

func (q *BoundedBlockingQueue) Size() int {
    return len(q.ch)
}

// 首先我还是没有懂题目意图， 题目是像让我准备一个数据结构对象，OOP编程中使用
// 其次我默认的情况下在type中初始化分配了内存，其实这个工作应该在实例化的时候做，另一个我默认的肌肉记忆就是对值操作了而不是对指针操作
// 然后我操作指针又错了，这里可以直接q.ch直接指明对象 但是我却写了&ch

//==================graceful增加shutdown()机制============================


// BoundedBlockingQueue 带有优雅退出机制的版本
type BoundedBlockingQueue struct {
	ch   chan int           // 数据平面
	done chan struct{}      // 控制平面
}

func Constructor(capacity int) BoundedBlockingQueue {
	return BoundedBlockingQueue{
		ch:   make(chan int, capacity),
		done: make(chan struct{}), 
	}
}

func (q *BoundedBlockingQueue) Enqueue(element int) bool {
	select {
	case <-q.done: //一旦外部 close(q.done)，就会触发这个，直接拒绝新的入队请求，防止变僵尸协程
		return false
	case q.ch <- element:
		return true
	}
}

func (q *BoundedBlockingQueue) Dequeue() (int, bool) { //同理dequeue这边也要做优雅退出，不能直接写阻塞等待queue中有数据，不然也成僵尸go routine了
	select {
	case <-q.done: //收到shutdown信号后执行清空当前队列
		select {
		case val := <-q.ch:
			return val, true 
		default:
			return 0, false  
		}
	case val := <-q.ch:
		return val, true
	}
}

func (q *BoundedBlockingQueue) Size() int {
	return len(q.ch)
}

func (q *BoundedBlockingQueue) Shutdown() {
	// 所有监听 <-q.done 的 select 都会立刻收到信号
	close(q.done)
}

// 开始的时候不清楚工程上用chan struct{} 做控制平面的习惯
// 后来不理解dequeue为什么也要做优雅退出：同样是解决消费者go routine阻塞变僵尸协程的问题
