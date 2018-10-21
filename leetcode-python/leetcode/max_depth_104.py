# coding: utf-8

# Given a binary tree, find its maximum depth.
# The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
# Given binary tree [3,9,20,null,null,15,7],
# 给定一个二叉树，找出其最大深度。
# 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
#      3
#     / \
#    9  20
#       /  \
#      15   7
#  return its depth = 3.
class Solution:

    # 递归
    def max_depth(self, root):
        if root is None:
            return 0
        else:
            left_height = self.max_depth(root.left)
            right_height = self.max_depth(root.right)
            return max(left_height, right_height) + 1

    # 迭代
    def max_depth_2(self, root):
        stack = []
        if root is not None:
            stack.append((root, 1))

        depth = 0
        while stack != []:
            root, current_depth = stack.pop()
            if root is not None:
                depth = max(current_depth, depth)
                stack.append((root.left, current_depth + 1))
                stack.append((root.right, current_depth + 1))

        return depth



