package com.lxy.leetcode;

import java.util.LinkedList;
import java.util.Queue;

import com.lxy.leetcode.common.TreeNode;

import javafx.util.Pair;

/**
 * @author lixinyan
 * @date 2018/10/14.
 */
//Given a binary tree, find its maximum depth.
//The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//Given binary tree [3,9,20,null,null,15,7],
//给定一个二叉树，找出其最大深度。
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
//      3
//     / \
//    9  20
//       /  \
//      15   7
//  return its depth = 3.
public class MaxDepth104 {

    // 递归
    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int leftHeight = maxDepth(root.left);
        int rightHeight = maxDepth(root.right);
        return Math.max(leftHeight, rightHeight) + 1;
    }

    // 迭代
    @SuppressWarnings("unchecked")
    public int maxDepth2(TreeNode root) {
        Queue<Pair<TreeNode, Integer>> stack = new LinkedList<>();
        if (root != null) {
            stack.add(new Pair(root, 1));
        }
        int depth = 0;
        while (!stack.isEmpty()) {
            Pair<TreeNode, Integer> current = stack.poll();
            root = current.getKey();
            int currentDepth = current.getValue();
            if (root != null) {
                depth = Math.max(depth, currentDepth);
                stack.add(new Pair(root.left, currentDepth + 1));
                stack.add(new Pair(root.right, currentDepth + 1));
            }
        }
        return depth;
    }
}