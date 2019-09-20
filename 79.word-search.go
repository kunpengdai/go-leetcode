/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (31.88%)
 * Likes:    2196
 * Dislikes: 111
 * Total Accepted:    328.8K
 * Total Submissions: 1M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 * 
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 * 
 * Example:
 * 
 * 
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 * 
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 * 
 * 
 */
func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]);j++ {
			if found(&board,word,i,j,0){
				return true
			}
		}
	}
	return false
}

func found(board *[][]byte, word string,i,j,pos int) bool {
	if pos == len(word){
		return true
	}
	if i>=len(*board) || j>=len((*board)[0]) || i<0|| j<0{
		return false
	}
	if word[pos]!=(*board)[i][j]{
		return false
	}
	(*board)[i][j] = 0
	ints := [][]int{[]int{0,1},[]int{0,-1},[]int{-1,0},[]int{1,0}}
	for _,t := range ints{
		(*board)[i][j] = 0
		if found(board,word,i+t[0],j+t[1],pos+1) {
			return true
		}
		(*board)[i][j]=word[pos]
	}
	return false
}

