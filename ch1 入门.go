// test
package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

var mu sync.Mutex
var count int

func main() {
	// 1.1 Hello, World
	// fmt.Println("Hello, World")

	// 1.2 命令行参数
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// 1.2 命令行参数
	// s, sep := "", ""
	// for idx, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// 	fmt.Println(idx, "->"+arg)
	// }
	// fmt.Println(s)

	// 1.2 命令行参数
	// s := ""		// 短变量声明，最简洁，但只能用在函数内部，而不能用于包变量
	// var s string // 依赖于字符串的默认初始化零值机制，被初始化为""
	// var s = ""	// 用得很少，除非同时声明多个变量
	// var s string = ""	// 用得很少，显式地标明变量的类型

	// 1.2 命令行参数
	// fmt.Println(strings.Join(os.Args[:], " "))

	// 1.2 命令行参数
	// fmt.Println(os.Args[1:])

	// 1.3 查找重复的行
	// counts := make(map[string]int) // map存储了键/值
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	// 每次调用 input.Scan() 则开始读入下一行，并移除行末的换行符，内容为 input.Text()
	// 	// Scan 函数在读到一行时返回 true，不再有输入时返回 false
	// 	if input.Text() == "false" {
	// 		break
	// 	}
	// 	counts[input.Text()]++
	// 	// 等价于下面两条语句，作统计每行出现的次数
	// 	// line := input.Text()
	// 	// counts[line] = counts[line] + 1
	// }
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d times:\t%s\n", n, line)
	// 	}
	// }

	// 1.3 查找重复的行
	// counts := make(map[string]int)
	// files := os.Args[1:]
	// if len(files) == 0 {
	// 	countLines(os.Stdin, counts) // 传递 map 为地址引用，类似于C++里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存
	// } else {
	// 	for _, arg := range files {
	// 		f, err := os.Open(arg)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	// 			continue
	// 		}
	// 		countLines(f, counts)
	// 		f.Close()
	// 	}
	// }
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line) // %d times -> line text
	// 	}
	// }

	// 1.3 查找重复的行
	// counts := make(map[string]int)
	// for _, filename := range os.Args[1:] {
	// 	data, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	// 		continue
	// 	}
	// 	for _, line := range strings.Split(string(data), "\n") {
	// 		counts[line]++
	// 	}
	// }
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }

	// 1.4 GIF动画
	// rand.Seed(time.Now().UnixNano())
	// lissajous(os.Stdout)

	// 1.5 获取URL
	// for _, url := range os.Args[1:] {
	// 	resp, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	b, err := ioutil.ReadAll(resp.Body)
	// 	resp.Body.Close()
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("%s", b)
	// }

	// 1.6 并发获取多个URL
	// start := time.Now()
	// ch := make(chan string)
	// for _, url := range os.Args[1:] {
	// 	go fetch(url, ch) // start a goroutine
	// }
	// for range os.Args[1:] {
	// 	fmt.Println((<-ch))
	// }
	// fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	// 1.7 Web服务
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8016", nil))
}

func countLines(f *os.File, counts map[string]int) { // 函数和包级别的变量可以任意顺序声明，并不影响其被调用
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "false" {
			break
		}
		counts[input.Text()]++
	}
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func Signum(x int) int {
	switch {
	case x > 0:
		return x + 1
	default:
		return 0
	case x < 0:
		return -1
	}
}
