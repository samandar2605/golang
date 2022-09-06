// package main
// import "fmt"
// func main(){
// 	ls:=[]int{}
// 	fmt.Println(buildArray(ls))
// }

// func buildArray(nums []int) []int {
//     ls:=[]int{}
// 	for i:=range nums{
// 		ls=append(ls, nums[nums[i]])
// 	}
// 	return ls
// }


// package main
// import "fmt"
// func main(){
// 	ls:=[]int{1,2,1}
// 	fmt.Println(getConcatenation(ls))
// }

// func getConcatenation(nums []int) []int {
//     ls:=[]int{}
// 	for i:=range nums{
// 		ls = append(ls, nums[i])
// 	}
// 	for i:=range nums{
// 		ls = append(ls, nums[i])
// 	}
// 	return ls
// }


// package main
// import "fmt"
// func main(){
// 	ls:=[]string{"X++","++X","--X","X--"}
// 	fmt.Println(finalValueAfterOperations(ls))
// }
// func finalValueAfterOperations(operations []string) int {
// 	n:=0
//     for i:=0; i<len(operations); i++{
// 		switch operations[i]{
// 		case "X++":
// 			n++
// 		case "++X":
// 			n++
// 		case "--X":
// 			n--
// 		case "X--":
// 			n--
// 		}
// 	}
// 	return n
// }



// package main
// import "fmt"
// func main(){
// 	s:="1.1.1.1"
// 	fmt.Println(defangIPaddr(s))
// }

// func defangIPaddr(address string) string {
//     h:=""
// 	for i:=range address{
// 		if address[i]==46{
// 			h+="[.]"
// 		}else{
// 			h+=string(address[i])
// 		}
// 	}
// 	return h
// }


// package main
// import "fmt"
// func main(){
// 	ls:=[][]int{{1,5},{7,3},{3,5}}
// 	fmt.Println(maximumWealth(ls))
// }

// func maximumWealth(accounts [][]int) int {
//     h:=sum(accounts[0])
// 	for i:=1; i<len(accounts); i++{
// 		m:=sum(accounts[i])
// 		if h<m{
// 			h=m
// 		}
// 	}
// 	return h
// }

// func sum(ls []int)int{
// 	h:=0
// 	for i:=range ls{
// 		h+=ls[i]
// 	}
// 	return h
// }

// package main
// import "fmt"
// func main(){
// 	ls:=[]int{1,2,3,4,5,6}
// 	n:=3
// 	fmt.Println(shuffle(ls,n))
// }


// func shuffle(nums []int, n int) []int {
//     ls:=[]int{}
// 	for i:=0; i<n; i++{
// 		ls = append(ls, nums[i])
// 		ls = append(ls, nums[i+n])
// 	}
// 	return ls
// }


// package main
// import "fmt"
// func main(){
// 	ls:=[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
// 	fmt.Println(mostWordsFound(ls))
// }
// func mostWordsFound(sentences []string) int {
//     h:=count(sentences[0])
// 	for i:=1; i<len(sentences);i++{
// 		m:=count(sentences[i])
// 		if h<m{
// 			h=m
// 		}
// 	}
// 	return h
// }

// func count(s string)int{
// 	h:=0
// 	for i:=range s{
// 		if s[i]==32{
// 			h++
// 		}
// 	}
// 	return h+1
// }


// package main
// import "fmt"
// func main(){
// 	s:="cat and  dog"
// 	fmt.Println(countValidWords(s))
// }

// func countValidWords(sentence string) int {
    
// }

// func isWord(s string)bool{
// 	h:=0
// 	if 123<s[0] || s[0]<97{
// 		return false
// 	}


// 	return false
// }

// func split(s string)[]string{
// 	h:=""
// 	ls:=[]string{}
// 	for i:=range s{
// 		if s[i]!=32{
// 			h+=string(s[i])
// 		}else{
// 			ls=append(ls, h)
// 			h=""
// 		}
// 	}
// 	if len(h)>0{
// 		ls = append(ls, h)
// 	}
// 	return ls
// }






// package main
// import "fmt"
// func main(){
// 	ls:=[]int{1,2,3,1,1,3}
// 	fmt.Println(numIdenticalPairs(ls))
// }

// func numIdenticalPairs(nums []int) int {
// 	h:=0
// 	for i:=0; i<len(nums)-1; i++{
// 		for j:=i+1; j<len(nums); j++{
// 			if nums[i]==nums[j] && i<j{
// 				h++
// 			}
// 		}
// 	}
// 	return h
// }



package main
import "fmt"
func main(){
	ls:=[]int{2,3,5,1,3}
	n:=3
	fmt.Println(kidsWithCandies(ls,n))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
    m:=max(candies)
	ls:=[]bool{}
	for i:=range candies{
		ls=append(ls, m<=candies[i]+extraCandies)	
	}
	return ls
}

func max(ls []int)int{
	h:=ls[0]
	for i:=1; i<len(ls); i++{
		if h<ls[i]{
			h=ls[i]
		}
	}
	return h
}























