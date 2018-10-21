package com.lxy.leetcode.common;

/**
 * @author lixinyan
 * @date 2018/10/14.
 */
// CHECKSTYLE:OFF
public class TreeNode {

    // 为方便外部访问，这里直接使用public修饰
    public int val;
    public TreeNode left;
    public TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}
// CHECKSTYLE:ON
