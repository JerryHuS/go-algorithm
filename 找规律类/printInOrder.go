/**
 * @Author: alessonhu
 * @Description: 虾皮一面算法
 * @File:  printInOrder
 * @Version: 1.0.0
 * @Date: 2022/5/9 10:33
 */
package 找规律类

import "fmt"

func printDog(dogChan, fishChan chan struct{}) {
	<-dogChan
	fmt.Println("dog")
	fishChan <- struct{}{}
}

func printFish(fishChan, catChan chan struct{}) {
	<-fishChan
	fmt.Println("fish")
	catChan <- struct{}{}
}

func printCat(catChan, dogChan chan struct{}) {
	<-catChan
	fmt.Println("cat")
	dogChan <- struct{}{}
}

func PrintGo() {
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	catChan := make(chan struct{}, 1)

	dogChan <- struct{}{}
	for i := 1; i <= 100; i++ {
		go printDog(dogChan, fishChan)
		go printFish(fishChan, catChan)
		go printCat(catChan, dogChan)
	}
	for {
	}
}
