package concurrency

import (
	"strconv"
	"sync"
	"testing"
)

var elems []string

const (
	numOfWriter = 50000
	numOfReader = 100000
)

func init() {
	for i := 0; i < 10; i++ {
		elems = append(elems, strconv.Itoa(i))
	}
}

type BasicSlice interface {
	Append(elem string)
	Travel(func(elem string))
}

type ThreadSafeWithRWLock struct {
	rwLock sync.RWMutex
	store  []string
}

func (s *ThreadSafeWithRWLock) Append(elem string) {
	s.rwLock.Lock()
	s.store = append(s.store, elem)
	s.rwLock.Unlock()
}

func (s *ThreadSafeWithRWLock) Travel(fn func(elem string)) {
	s.rwLock.Lock()
	for _, e := range s.store {
		fn(e)
	}
	s.rwLock.Unlock()
}

type ThreadSafeWithChannel struct {
	store chan string
}

func NewThreadSafeWithChannel(bufSize int) *ThreadSafeWithChannel {
	ch = make(chan string, bufSize)
	return &ThreadSafeWithChannel{ch}
}

func (s *ThreadSafeWithChannel) Append(elem string) {
	select {
	case s.store <- elem:
	default:

	}
}

func (s *ThreadSafeWithChannel) Travel(fn func(elem string)) {
	tmp, ok := s.store.Load().(*[]string)
	if !ok {
		panic("can not convert to *[]string")
	}
	return *tmp
}

func testBasicOP(bs BasicSlice, t *testing.T) {
	for _, e := range elems {
		bs.Append(e)
	}
	if len(elems) != len(bs.Get()) {
		t.Error("Failed to implement basic functions")
	}
}

func TestWithRWLock(t *testing.T) {
	testBasicOP(&ThreadSafeWithRWLock{}, t)
}

func TestWithAtomic(t *testing.T) {
	testBasicOP(NewThreadSafeWithAtomic(), t)
}

func ConcurrenctTest(bs BasicSlice, b *testing.B) {
	var wg sync.WaitGroup
	wfn := func(bs BasicSlice) {
		for _, e := range elems {
			bs.Append(e)
		}
		wg.Done()
	}
	rfn := func(bs BasicSlice) {
		_ = bs.Get()
		wg.Done()
	}
	for i := 0; i < numOfWriter; i++ {
		go wfn(bs)
		wg.Add(1)
	}
	for j := 0; j <= numOfReader; j++ {
		go rfn(bs)
		wg.Add(1)
	}
	wg.Wait()
	if len(bs.Get()) != numOfWriter*len(elems) {
		b.Error("Missed some writings")
	}
}

// func BenchmarkWithRWLock(b *testing.B) {
// 	b.ResetTimer()
// 	ConcurrenctTest(&ThreadSafeWithRWLock{}, b)
// }

func BenchmarkWithAtomic(b *testing.B) {
	b.ResetTimer()
	ConcurrenctTest(NewThreadSafeWithAtomic(), b)
}
