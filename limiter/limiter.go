package limiter

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶结构
type TokenBucket struct {
	limitRate int           // 限制频率，即每分钟加入多少个令牌
	tokenChan chan struct{} // 令牌通道，可以理解为桶
	cap       int           // 令牌桶的容量
	muLock    sync.Mutex    // 令牌桶锁，保证线程安全
	stop      bool          // 停止标记，结束令牌桶
}

func (b *TokenBucket) produce2() {
	for {
		b.muLock.Lock()
		if b.stop {
			close(b.tokenChan)
			b.muLock.Unlock()
			return
		}
		b.tokenChan <- struct{}{}
		b.muLock.Unlock()
		time.Sleep(time.Minute / time.Duration(b.limitRate))
	}

}

// NewTokenBucket 创建令牌桶
func NewTokenBucket(limitRate, cap int) *TokenBucket {
	if cap < 1 {
		panic("token bucket cap must be large 1")
	}
	return &TokenBucket{
		tokenChan: make(chan struct{}, cap),
		limitRate: limitRate,
		cap:       cap,
	}
}

// Start 开启令牌桶
func (b *TokenBucket) Start() {
	go b.produce()
}

// 生产令牌
func (b *TokenBucket) produce() {
	for {
		b.muLock.Lock()
		if b.stop {
			close(b.tokenChan)
			b.muLock.Unlock()
			return
		}
		b.tokenChan <- struct{}{}
		d := time.Minute / time.Duration(b.limitRate)
		b.muLock.Unlock()
		time.Sleep(d)
	}
}

// Consume 消费令牌
func (b *TokenBucket) Consume() {
	<-b.tokenChan
}

// Stop 停止令牌桶
func (b *TokenBucket) Stop() {
	b.muLock.Lock()
	defer b.muLock.Unlock()
	b.stop = true
}
