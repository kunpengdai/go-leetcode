/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 */

// @lc code=start
func numEnclaves(A [][]int) int {

	if len(A)==0{
		return 0
	}

	dp := make([][]int, len(A))
	for i := range dp{
		dp[i] = make([]int, len(dp[0]))
	}


	for i:=0;i<len(A);i++{
		bfs(&A,i,0)
	}
	for i:=len(A)-1;i>=0;i--{
		bfs(&A,i,len(A[0])-1)
	}
	for i:=0;i<len(A[0]);i++{
		bfs(&A,0,i)
	}
	for i:=len(A[0])-1;i>=0;i--{
		bfs(&A,len(A)-1,i)
	}

	res := 0
	for i :=0;i<len(A);i++{
		for j :=0;j<len(A[0])-1;j++{
			if A[i][j]==1{
				res++
			}
		}
	}

	return res
}

func bfs(A *[][]int,i,j int){
	dr := [][]int{
		{0,1},
		{0,-1},
		{1,0},
		{-1,0},
	}
	if (*A)[i][j]==1{
		(*A)[i][j]=0
		for k :=0;k<4;k++{
			ix := i+ dr[k][0]
			jx := j+ dr[k][1]
			if ix>=0 && ix <len(*A)&& jx >=0&&jx < len((*A)[0]){
				bfs(A,ix,jx)
			}
		}
	} 
}
// @lc code=end

