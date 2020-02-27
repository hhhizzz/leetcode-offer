package _7

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    rootNumber := preorder[0]
    rootIn := 0
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == rootNumber {
            rootIn = i
            break
        }
    }
    result := &TreeNode{}
    result.Val = rootNumber
    result.Left = buildTree(preorder[1:1+rootIn], inorder[0:rootIn])
    result.Right = buildTree(preorder[rootIn+1:], inorder[1+rootIn:])

    return result
}
