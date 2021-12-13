/*
* @Author: Yajun
* @Date:   2021/12/12 17:34
 */

package chap3

import (
	"fmt"
	"testing"
)

func Test3_1_1(t *testing.T) {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isValidPalindrome(str))
}

func Test3_1_2(t *testing.T) {
	str := "ATGTGAGCTGGTGTGTGCFAA"
	pattern := "GTGTGC"
	fmt.Println(strStr(str, pattern))
	fmt.Println(strStrB(str, pattern))
	fmt.Println(strStrC(str, pattern))
	fmt.Println(strStrD(str, pattern))
}
