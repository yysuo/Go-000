package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

type indexHandler struct {
	content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ih.content)
}

type Error struct {
	Sig os.Signal
}
func NewError(sig os.Signal) *Error {
	return &Error{Sig:sig}
}
func (err *Error) Error() string {
	return err.Sig.String()
}

func main() {
	var g errgroup.Group
	// 创建一个os.Signal channel
	sigs := make(chan os.Signal, 1)
	//注册要接收的信号，syscall.SIGINT:接收ctrl+c ,syscall.SIGTERM:程序退出
	//信号没有信号参数表示接收所有的信号
	signal.Notify(sigs)




	g.Go(func() error{
			http.Handle("/", &indexHandler{content: "hello world!"})
			err := http.ListenAndServe(":8001", nil)
			if err !=nil {
				return err
			}
			sig := <-sigs
			return NewError(sig)
	})


	g.Go(func() error{
		http.Handle("/2", &indexHandler{content: "hello world!"})
		err := http.ListenAndServe(":8002", nil)
		fmt.Println("dadada")
		if err !=nil {
			return err
		}

		sig := <- sigs
		return NewError(sig)

	})
	
	err := g.Wait()
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
}