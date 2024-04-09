package batcher // 包名为 batcher

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

// 定义一个错误变量 ErrFull，表示通道已满
var ErrFull = errors.New("channel is full")

// Option 接口用于定义配置选项
type Option interface {
	apply(*options) // apply 方法用于将选项应用到配置中
}

// options 结构体用于保存批处理器的配置选项
type options struct {
	size     int           // 批处理器的大小
	buffer   int           // 缓冲区大小
	worker   int           // 工作线程数
	interval time.Duration // 批处理间隔
}

// check 方法用于检查配置选项，如果选项值不合法则设置为默认值
func (o options) check() {
	if o.size <= 0 {
		o.size = 100
	}
	if o.buffer <= 0 {
		o.buffer = 100
	}
	if o.worker <= 0 {
		o.worker = 5
	}
	if o.interval <= 0 {
		o.interval = time.Second
	}
}

// funcOption 结构体用于定义函数选项
type funcOption struct {
	f func(*options) // f 函数用于配置选项
}

// apply 方法用于将选项应用到配置中
func (fo *funcOption) apply(o *options) {
	fo.f(o)
}

// newOption 函数用于创建新的函数选项
func newOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// WithSize 函数用于设置批处理器的大小选项
func WithSize(s int) Option {
	return newOption(func(o *options) {
		o.size = s
	})
}

// WithBuffer 函数用于设置批处理器的缓冲区大小选项
func WithBuffer(b int) Option {
	return newOption(func(o *options) {
		o.buffer = b
	})
}

// WithWorker 函数用于设置批处理器的工作线程数选项
func WithWorker(w int) Option {
	return newOption(func(o *options) {
		o.worker = w
	})
}

// WithInterval 函数用于设置批处理器的批处理间隔选项
func WithInterval(i time.Duration) Option {
	return newOption(func(o *options) {
		o.interval = i
	})
}

// msg 结构体用于表示批处理器中的消息
type msg struct {
	key string      // 消息的键
	val interface{} // 消息的值
}

// Batcher 结构体用于定义批处理器
type Batcher struct {
	opts     options                                                 // 批处理器的配置选项
	Do       func(ctx context.Context, val map[string][]interface{}) // 处理函数，用于处理批量的消息数据
	Sharding func(key string) int                                    // 分片函数，根据消息的键进行分片
	chans    []chan *msg                                             // 消息通道数组，用于接收消息
	wait     sync.WaitGroup                                          // 等待组，用于等待所有工作线程完成
}

// New 函数用于创建一个新的批处理器
func New(opts ...Option) *Batcher {
	b := &Batcher{}
	for _, opt := range opts {
		opt.apply(&b.opts) // 应用配置选项
	}
	b.opts.check() // 检查配置选项是否合法

	b.chans = make([]chan *msg, b.opts.worker) // 初始化消息通道数组
	for i := 0; i < b.opts.worker; i++ {
		b.chans[i] = make(chan *msg, b.opts.buffer)
	}
	return b
}

// Start 方法用于启动批处理器
func (b *Batcher) Start() {
	if b.Do == nil {
		log.Fatal("Batcher: Do func is nil") // 处理函数为空则直接退出程序
	}
	if b.Sharding == nil {
		log.Fatal("Batcher: Sharding func is nil") // 分片函数为空则直接退出程序
	}
	b.wait.Add(len(b.chans)) // 等待组数量与通道数相同
	for i, ch := range b.chans {
		go b.merge(i, ch) // 启动工作线程
	}
}

// Add 方法用于向批处理器添加消息
func (b *Batcher) Add(key string, val interface{}) error {
	ch, msg := b.add(key, val) // 添加消息到指定通道中
	select {
	case ch <- msg: // 将消息放入通道
	default:
		return ErrFull // 如果通道已满则返回错误
	}
	return nil
}

// add 方法用于向指定通道添加消息
func (b *Batcher) add(key string, val interface{}) (chan *msg, *msg) {
	sharding := b.Sharding(key) % b.opts.worker // 根据消息键进行分片
	ch := b.chans[sharding]                     // 选择对应的通道
	msg := &msg{key: key, val: val}             // 创建消息
	return ch, msg
}

// merge 方法用于合并消息并进行批量处理
func (b *Batcher) merge(idx int, ch <-chan *msg) {
	defer b.wait.Done() // 减少等待组数量

	var (
		msg        *msg                                          // 消息
		count      int                                           // 计数器
		closed     bool                                          // 是否关闭通道
		lastTicker = true                                        // 上次是否是定时器触发
		interval   = b.opts.interval                             // 批处理间隔
		vals       = make(map[string][]interface{}, b.opts.size) // 消息键值对
	)
	if idx > 0 {
		interval = time.Duration(int64(idx) * (int64(b.opts.interval) / int64(b.opts.worker))) // 计算批处理间隔
	}
	ticker := time.NewTicker(interval) // 创建定时器
	for {
		select {
		case msg = <-ch: // 从通道中接收消息
			if msg == nil { // 如果接收到空消息
				closed = true // 标记通道已关闭
				break         // 结束当前循环
			}
			count++                                        // 计数加一
			vals[msg.key] = append(vals[msg.key], msg.val) // 添加消息到键值对中
			if count >= b.opts.size {                      // 如果达到批处理大小
				break // 结束当前循环
			}
			continue // 继续下一次循环
		case <-ticker.C: // 定时器触发
			if lastTicker { // 如果上次也是定时器触发
				ticker.Stop()                            // 停止定时器
				ticker = time.NewTicker(b.opts.interval) // 重新创建定时器
				lastTicker = false                       // 标记上次不是定时器触发
			}
		}
		if len(vals) > 0 { // 如果有消息需要处理
			ctx := context.Background()                        // 创建上下文
			b.Do(ctx, vals)                                    // 执行处理函数
			vals = make(map[string][]interface{}, b.opts.size) // 重置键值对
			count = 0                                          // 计数器清零
		}
		if closed { // 如果通道已关闭
			ticker.Stop() // 停止定时器
			return        // 结束函数
		}
	}
}

// Close 方法用于关闭批处理器
func (b *Batcher) Close() {
	for _, ch := range b.chans {
		ch <- nil // 向通道中发送空消息
	}
	b.wait.Wait() // 等待所有工作线程完成
}
