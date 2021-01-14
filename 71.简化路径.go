/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	pathes := strings.Split(path, "/")
	validPath := make([]string, 0, len(pathes))
	for _, p := range pathes {
		switch {
		case p == "" || p == ".":
		case p == "..":
			if len(validPath) > 0 {
				validPath = validPath[:len(validPath)-1]
			}
		default:
			validPath = append(validPath, p)
		}
	}

	return "/" + strings.Join(validPath, "/")
}

// @lc code=end

