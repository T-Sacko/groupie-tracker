package main

// import (
// 	"fmt"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	Chunk([]int{22, 33, 44, 5, 6}, 2)
// 	z01.PrintRune('4' + 48)
// }
// func Chunk(slice []int, size int) {
// 	var temp [][]int

// 	for i := 0; i < len(slice); i += size {

// 		end := i + size
// 		if end > len(slice) {
// 			end = len(slice)
// 		}
// 		temp = append(temp, slice[i:end])
// 	}
// 	fmt.Println(temp)
// }

// // func chunk(slice []int, size int) {
// // 	if size == 0 {
// // 		fmt.Println()
// // 		return
// // 	}
// // 	var chunks [][]int
// // 	for i := 0; i < len(slice); i += size {
// // 		end := i + size
// // 		if end > len(slice) {
// // 			end = len(slice)
// // 		}
// // 		chunks = append(chunks, slice[i:end])
// // 	}
// // 	fmt.Println(chunks)
// // }
