package binary_tree

func GetSuccessorNode(node *NodeWithParent) *NodeWithParent {
	if node == nil {
		return node
	}
	if node.right != nil {
		return getLeftMost(node.right)
	} else { // 無右子樹
		parent := node.parent
		for parent != nil && parent.right == node { // 父節點為 nil，或 node 為父節點的左孩子時停止
			node = parent
			parent = node.parent
		}
		return parent
	}
}

func getLeftMost(node *NodeWithParent) *NodeWithParent {
	if node == nil {
		return node
	}
	for node.left != nil {
		node = node.left
	}
	return node
}
