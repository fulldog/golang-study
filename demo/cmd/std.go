package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func a1(ctx context.Context) {
	ctx = context.WithValue(ctx, "a1", "a1")
	a2(ctx)
}
func a2(ctx context.Context) {
	ctx = context.WithValue(ctx, "a2", "a22222")
	ctx = context.WithValue(ctx, "a333", "33333333333")
	a3(ctx)
}
func a3(ctx context.Context) {
	ctx = context.WithValue(ctx, "dddddddddd", "ddddddddddddd")
	a4(ctx)
}
func a4(ctx context.Context) {
	fmt.Println(ctx.Value("a1"))
	fmt.Println(ctx.Value("a2"))
	fmt.Println(ctx.Value("a333"))
	fmt.Println(ctx.Value("dddddddddd"))
}

var rmx sync.RWMutex

func read(string2 int) {
	rmx.RLock()
	fmt.Println(string2)
}
func rw(string2 string) {
	rmx.Lock()
	fmt.Println("rw success")
}

func add(x *int) {
	*x += *x
	x = nil
}

func g1(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	go g2(ctx)
	n := 1
	for {
		select {
		case x := <-ctx.Done():
			fmt.Println(ctx.Err(), "g1 tuichu", x)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("g1", n)
			n++
		}
	}
}

func g2(ctx context.Context) {
	n := 1
	for {
		select {
		case x := <-ctx.Done():
			fmt.Println(ctx.Err(), "g2 tuichu", x)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("g2", n)
			n++
		}
	}
}

func main() {
	var k uint32
	for i := 0; i < 100; i++ {
		go func() {
			if atomic.CompareAndSwapUint32(&k, 0, 1) {
				fmt.Println("ok")
			} else {
				fmt.Println("not")
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println(k)

	fmt.Println(s2b("sdfdsfdsf"))
	//ctx, cancle := context.WithTimeout(context.Background(),5*time.Second)
	//ctx = context.WithValue(ctx,"p","p")
	//go g1(ctx)
	//go g2(ctx)
	//cancle()
	//
	//time.Sleep(5*time.Second)
	// //cancle()
	//time.Sleep(2*time.Second)

}

func sdt() {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)

	ss := strings.Split(msg, " ")
	fmt.Println(len(ss[len(ss)-1]))
}

//100个并发 每次只执行10个
func t() {
	sy := sync.WaitGroup{}
	sy.Add(100)
	ch := make(chan struct{}, 10)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer func() {
				sy.Done()
				<-ch
			}()
			ch <- struct{}{}
			time.Sleep(8 * time.Second)
			fmt.Println(i)
		}(i)
	}
	sy.Wait()
	close(ch)
}

// s2b converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func s2b(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
