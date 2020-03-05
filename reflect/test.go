package reflect

import "fmt"

type Fn struct {}

func (f * Fn) ProcessA(b string) {
	fmt.Println("hahah")
}

func init()  {
	handlers := newSubscriber(new(Fn))
	if err := createSubHandler(handlers); err != nil {
		fmt.Println("loi")
	}
}
