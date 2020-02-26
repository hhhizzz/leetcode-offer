package _4

func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    i, j := len(matrix)-1, 0
    for i >= 0 && j < len(matrix[i]) {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] < target {
            j++
        } else {
            i--
        }
    }
    return false
}
