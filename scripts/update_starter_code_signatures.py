#!/usr/bin/env python3
"""
update_starter_code_signatures.py - Upgrades starter_code and test cases
across all algorithmic problems in data/problems/dataset_all.json from
generic `*args, **kwargs` to clean, semantic parameter signatures.
"""

import json
import os
import re
import sys

EXACT_SIGNATURES = {
    # Canonical LeetCode exact signatures
    "twoSum": ["nums", "target"],
    "twoSumIiInputArrayIsSorted": ["numbers", "target"],
    "twoSumIiiDataStructureDesign": ["nums"],
    "twoSumIvInputIsABst": ["root", "k"],
    "twoSumLessThanK": ["nums", "k"],
    "twoSumBsts": ["root1", "root2", "target"],
    "prob3Sum": ["nums"],
    "threeSum": ["nums"],
    "prob3SumClosest": ["nums", "target"],
    "threeSumClosest": ["nums", "target"],
    "prob3SumSmaller": ["nums", "target"],
    "prob4Sum": ["nums", "target"],
    "fourSum": ["nums", "target"],
    "prob4SumIi": ["nums1", "nums2", "nums3", "nums4"],
    "addTwoNumbers": ["l1", "l2"],
    "addTwoNumbersIi": ["l1", "l2"],
    "medianOfTwoSortedArrays": ["nums1", "nums2"],
    "longestSubstringWithoutRepeatingCharacters": ["s"],
    "longestPalindromicSubstring": ["s"],
    "zigzagConversion": ["s", "numRows"],
    "reverseInteger": ["x"],
    "stringToIntegerAtoi": ["s"],
    "palindromeNumber": ["x"],
    "regularExpressionMatching": ["s", "p"],
    "containerWithMostWater": ["height"],
    "integerToRoman": ["num"],
    "romanToInteger": ["s"],
    "longestCommonPrefix": ["strs"],
    "letterCombinationsOfAPhoneNumber": ["digits"],
    "removeNthNodeFromEndOfList": ["head", "n"],
    "validParentheses": ["s"],
    "mergeTwoSortedLists": ["list1", "list2"],
    "generateParentheses": ["n"],
    "mergeKSortedLists": ["lists"],
    "swapNodesInPairs": ["head"],
    "reverseNodesInKGroup": ["head", "k"],
    "removeDuplicatesFromSortedArray": ["nums"],
    "removeElement": ["nums", "val"],
    "findTheIndexOfTheFirstOccurrenceInAString": ["haystack", "needle"],
    "divideTwoIntegers": ["dividend", "divisor"],
    "substringWithConcatenationOfAllWords": ["s", "words"],
    "nextPermutation": ["nums"],
    "longestValidParentheses": ["s"],
    "searchInRotatedSortedArray": ["nums", "target"],
    "findFirstAndLastPositionOfElementInSortedArray": ["nums", "target"],
    "searchInsertPosition": ["nums", "target"],
    "validSudoku": ["board"],
    "sudokuSolver": ["board"],
    "countAndSay": ["n"],
    "combinationSum": ["candidates", "target"],
    "combinationSumIi": ["candidates", "target"],
    "combinationSumIii": ["k", "n"],
    "combinationSumIv": ["nums", "target"],
    "firstMissingPositive": ["nums"],
    "trappingRainWater": ["height"],
    "multiplyStrings": ["num1", "num2"],
    "wildcardMatching": ["s", "p"],
    "jumpGame": ["nums"],
    "jumpGameIi": ["nums"],
    "permutations": ["nums"],
    "permutationsIi": ["nums"],
    "rotateImage": ["matrix"],
    "groupAnagrams": ["strs"],
    "myPow": ["x", "n"],
    "nQueens": ["n"],
    "nQueensIi": ["n"],
    "maximumSubarray": ["nums"],
    "spiralMatrix": ["matrix"],
    "canJump": ["nums"],
    "mergeIntervals": ["intervals"],
    "insertInterval": ["intervals", "newInterval"],
    "lengthOfLastWord": ["s"],
    "spiralMatrixIi": ["n"],
    "rotateList": ["head", "k"],
    "uniquePaths": ["m", "n"],
    "uniquePathsWithObstacles": ["obstacleGrid"],
    "minimumPathSum": ["grid"],
    "validNumber": ["s"],
    "plusOne": ["digits"],
    "addBinary": ["a", "b"],
    "textJustification": ["words", "maxWidth"],
    "mySqrt": ["x"],
    "climbingStairs": ["n"],
    "simplifyPath": ["path"],
    "editDistance": ["word1", "word2"],
    "setMatrixZeroes": ["matrix"],
    "searchA2dMatrix": ["matrix", "target"],
    "searchA2dMatrixIi": ["matrix", "target"],
    "sortColors": ["nums"],
    "minimumWindowSubstring": ["s", "t"],
    "subsets": ["nums"],
    "subsetsIi": ["nums"],
    "wordSearch": ["board", "word"],
    "wordSearchIi": ["board", "words"],
    "removeDuplicatesFromSortedArrayIi": ["nums"],
    "searchInRotatedSortedArrayIi": ["nums", "target"],
    "removeDuplicatesFromSortedList": ["head"],
    "removeDuplicatesFromSortedListIi": ["head"],
    "largestRectangleInHistogram": ["heights"],
    "maximalRectangle": ["matrix"],
    "partitionList": ["head", "x"],
    "scrambleString": ["s1", "s2"],
    "mergeSortedArray": ["nums1", "m", "nums2", "n"],
    "grayCode": ["n"],
    "decodeWays": ["s"],
    "decodeWaysIi": ["s"],
    "reverseLinkedListIi": ["head", "left", "right"],
    "restoreIpAddresses": ["s"],
    "binaryTreeInorderTraversal": ["root"],
    "uniqueBinarySearchTrees": ["n"],
    "uniqueBinarySearchTreesIi": ["n"],
    "validateBinarySearchTree": ["root"],
    "recoverBinarySearchTree": ["root"],
    "sameTree": ["p", "q"],
    "symmetricTree": ["root"],
    "binaryTreeLevelOrderTraversal": ["root"],
    "binaryTreeZigzagLevelOrderTraversal": ["root"],
    "maximumDepthOfBinaryTree": ["root"],
    "constructBinaryTreeFromPreorderAndInorderTraversal": ["preorder", "inorder"],
    "constructBinaryTreeFromInorderAndPostorderTraversal": ["inorder", "postorder"],
    "binaryTreeLevelOrderTraversalIi": ["root"],
    "convertSortedArrayToBinarySearchTree": ["nums"],
    "convertSortedListToBinarySearchTree": ["head"],
    "balancedBinaryTree": ["root"],
    "minimumDepthOfBinaryTree": ["root"],
    "pathSum": ["root", "targetSum"],
    "pathSumIi": ["root", "targetSum"],
    "pathSumIii": ["root", "targetSum"],
    "flattenBinaryTreeToLinkedList": ["root"],
    "distinctSubsequences": ["s", "t"],
    "populatingNextRightPointersInEachNode": ["root"],
    "populatingNextRightPointersInEachNodeIi": ["root"],
    "pascalsTriangle": ["numRows"],
    "pascalsTriangleIi": ["rowIndex"],
    "triangle": ["triangle"],
    "bestTimeToBuyAndSellStock": ["prices"],
    "bestTimeToBuyAndSellStockIi": ["prices"],
    "bestTimeToBuyAndSellStockIii": ["prices"],
    "bestTimeToBuyAndSellStockIv": ["k", "prices"],
    "bestTimeToBuyAndSellStockWithCooldown": ["prices"],
    "bestTimeToBuyAndSellStockWithTransactionFee": ["prices", "fee"],
    "binaryTreeMaximumPathSum": ["root"],
    "validPalindrome": ["s"],
    "validPalindromeIi": ["s"],
    "wordLadder": ["beginWord", "endWord", "wordList"],
    "wordLadderIi": ["beginWord", "endWord", "wordList"],
    "longestConsecutiveSequence": ["nums"],
    "sumRootToLeafNumbers": ["root"],
    "surroundedRegions": ["board"],
    "palindromePartitioning": ["s"],
    "palindromePartitioningIi": ["s"],
    "cloneGraph": ["node"],
    "gasStation": ["gas", "cost"],
    "candy": ["ratings"],
    "singleNumber": ["nums"],
    "singleNumberIi": ["nums"],
    "singleNumberIii": ["nums"],
    "copyListWithRandomPointer": ["head"],
    "wordBreak": ["s", "wordDict"],
    "wordBreakIi": ["s", "wordDict"],
    "linkedListCycle": ["head"],
    "linkedListCycleIi": ["head"],
    "reorderList": ["head"],
    "binaryTreePreorderTraversal": ["root"],
    "binaryTreePostorderTraversal": ["root"],
    "insertionSortList": ["head"],
    "sortList": ["head"],
    "maxPointsOnALine": ["points"],
    "evaluateReversePolishNotation": ["tokens"],
    "reverseWordsInAString": ["s"],
    "maximumProductSubarray": ["nums"],
    "findMinimumInRotatedSortedArray": ["nums"],
    "findMinimumInRotatedSortedArrayIi": ["nums"],
    "findPeakElement": ["nums"],
    "fractionToRecurringDecimal": ["numerator", "denominator"],
    "majorityElement": ["nums"],
    "majorityElementIi": ["nums"],
    "dungeonGame": ["dungeon"],
    "houseRobber": ["nums"],
    "houseRobberIi": ["nums"],
    "houseRobberIii": ["root"],
    "numberOfIslands": ["grid"],
    "numberOfIslandsIi": ["m", "n", "positions"],
    "numIslands": ["grid"],
    "bitwiseAndOfNumbersRange": ["left", "right"],
    "happyNumber": ["n"],
    "removeLinkedListElements": ["head", "val"],
    "isomorphicStrings": ["s", "t"],
    "reverseLinkedList": ["head"],
    "courseSchedule": ["numCourses", "prerequisites"],
    "courseScheduleIi": ["numCourses", "prerequisites"],
    "courseScheduleIii": ["courses"],
    "courseScheduleIv": ["numCourses", "prerequisites", "queries"],
    "canFinish": ["numCourses", "prerequisites"],
    "findOrder": ["numCourses", "prerequisites"],
    "minimumSizeSubarraySum": ["target", "nums"],
    "shortestPalindrome": ["s"],
    "kthLargestElementInAnArray": ["nums", "k"],
    "findKthLargest": ["nums", "k"],
    "containsDuplicate": ["nums"],
    "containsDuplicateIi": ["nums", "k"],
    "containsDuplicateIii": ["nums", "indexDiff", "valueDiff"],
    "invertBinaryTree": ["root"],
    "basicCalculator": ["s"],
    "basicCalculatorIi": ["s"],
    "basicCalculatorIii": ["s"],
    "palindromeLinkedList": ["head"],
    "lowestCommonAncestorOfABinarySearchTree": ["root", "p", "q"],
    "lowestCommonAncestorOfABinaryTree": ["root", "p", "q"],
    "lowestCommonAncestor": ["root", "p", "q"],
    "deleteNodeInALinkedList": ["node"],
    "productOfArrayExceptSelf": ["nums"],
    "slidingWindowMaximum": ["nums", "k"],
    "differentWaysToAddParentheses": ["expression"],
    "validAnagram": ["s", "t"],
    "binaryTreePaths": ["root"],
    "addDigits": ["num"],
    "uglyNumber": ["n"],
    "uglyNumberIi": ["n"],
    "superUglyNumber": ["n", "primes"],
    "missingNumber": ["nums"],
    "alienOrder": ["words"],
    "findTheDuplicateNumber": ["nums"],
    "gameOfLife": ["board"],
    "wordPattern": ["pattern", "s"],
    "longestIncreasingSubsequence": ["nums"],
    "lengthOfLIS": ["nums"],
    "minimumHeightTrees": ["n", "edges"],
    "burstBalloons": ["nums"],
    "superPow": ["a", "b"],
    "countOfSmallerNumbersAfterSelf": ["nums"],
    "removeInvalidParentheses": ["s"],
    "coinChange": ["coins", "amount"],
    "coinChangeIi": ["amount", "coins"],
    "numberOfConnectedComponentsInAnUndirectedGraph": ["n", "edges"],
    "countComponents": ["n", "edges"],
    "wiggleSort": ["nums"],
    "wiggleSortIi": ["nums"],
    "topKFrequentElements": ["nums", "k"],
    "topKFrequent": ["nums", "k"],
    "intersectionOfTwoArrays": ["nums1", "nums2"],
    "intersectionOfTwoArraysIi": ["nums1", "nums2"],
    "russianDollEnvelopes": ["envelopes"],
    "pacificAtlanticWaterFlow": ["heights"],
    "pacificAtlantic": ["heights"],
    "accountsMerge": ["accounts"],
    "networkDelayTime": ["times", "n", "k"],
    "cheapestFlightsWithinKStops": ["n", "flights", "src", "dst", "k"],
    "dailyTemperatures": ["temperatures"],
    "swimInWater": ["grid"],
    "swimInRisingWater": ["grid"],
    "busRoutes": ["routes", "source", "target"],
    "shortestPathInBinaryMatrix": ["grid"],
    "shortestPathBinaryMatrix": ["grid"],
    "rottingOranges": ["grid"],
    "orangesRotting": ["grid"],
    "asFarFromLandAsPossible": ["grid"],
    "shortestPathVisitingAllNodes": ["graph"],
    "criticalConnectionsInANetwork": ["n", "connections"],
    "criticalConnections": ["n", "connections"],
}

def infer_parameters(problem):
    entry = problem.get("entrypoint", "solve")
    title = problem.get("title", "")
    t_low = title.lower()
    e_low = entry.lower()
    tags = set(problem.get("topic_tags") or [])
    concepts = set(problem.get("primary_concepts") or [])

    if entry in EXACT_SIGNATURES:
        return EXACT_SIGNATURES[entry]

    # Specific problem name matching
    if "accounts merge" in t_low or "accountsmerge" in e_low:
        return ["accounts"]
    if "two sum" in t_low or "twosum" in e_low:
        return ["nums", "target"]
    if "3sum" in t_low or "threesum" in e_low:
        return ["nums"]
    if "4sum" in t_low or "foursum" in e_low:
        return ["nums", "target"]

    # Tree structures
    if any(tag in tags for tag in ("tree", "binary-tree", "binary-search-tree")) or "trees" in concepts or "tree" in t_low:
        if "same" in t_low or "issame" in e_low:
            return ["p", "q"]
        if "lowest common ancestor" in t_low or "lca" in e_low:
            return ["root", "p", "q"]
        if "construct" in t_low or "build" in e_low:
            return ["preorder", "inorder"]
        return ["root"]

    # Linked Lists
    if "linked-list" in tags or "linked_list" in concepts or "list node" in t_low or "linked list" in t_low:
        if "merge" in t_low and "two" in t_low:
            return ["list1", "list2"]
        if "k-group" in t_low or "kgroup" in e_low:
            return ["head", "k"]
        return ["head"]

    # Matrices and Grids
    if "matrix" in tags or "grid" in e_low or "matrix" in e_low or "board" in e_low or "grid" in t_low:
        if "target" in e_low or "search" in e_low:
            return ["matrix", "target"]
        if "board" in e_low or "sudoku" in t_low or "queens" in t_low:
            return ["board"]
        return ["grid"]

    # Graphs
    if "graph" in tags or "graphs" in concepts or "topological" in t_low or "scc" in t_low or "graph" in t_low:
        if "course" in t_low or "prereq" in t_low:
            return ["numCourses", "prerequisites"]
        if "flow" in t_low or "dinic" in t_low or "edmonds" in t_low:
            return ["n", "edges", "source", "sink"]
        if "shortest" in t_low or "dijkstra" in t_low or "bellman" in t_low:
            return ["n", "edges", "src", "dst"]
        if "bipartite" in t_low or "matching" in t_low or "mst" in t_low or "kruskal" in t_low or "prim" in t_low:
            return ["n", "edges"]
        if "bridge" in t_low or "articulation" in t_low or "tarjan" in t_low or "kosaraju" in t_low:
            return ["n", "edges"]
        if "clone" in t_low:
            return ["node"]
        return ["n", "edges"]

    # Strings
    if "string" in tags or "strings" in concepts or "string" in t_low or "substring" in t_low or "palindrome" in t_low:
        if "match" in e_low or "isomorphic" in e_low or "edit" in t_low:
            return ["s1", "s2"]
        if "anagram" in t_low and "group" in t_low:
            return ["strs"]
        if "words" in e_low or "dictionary" in t_low:
            return ["words"]
        return ["s"]

    # Target / K parameters
    if "target" in e_low or "target" in t_low:
        return ["nums", "target"]
    if re.search(r"\b[Kk]\b", title) or re.search(r"[Kk]th", entry) or "withk" in e_low:
        return ["nums", "k"]

    # Intervals
    if "interval" in tags or "intervals" in t_low:
        return ["intervals"]

    # Arrays and Sorting
    if "array" in tags or "arrays" in concepts or "sorting" in tags or "hashmap" in concepts or "array" in t_low:
        if "two" in t_low and ("arrays" in t_low or "sorted" in t_low):
            return ["nums1", "nums2"]
        return ["nums"]

    # Math and Numbers
    if "math" in tags or "geometry" in tags or "number" in t_low:
        if "gcd" in e_low or "lcm" in e_low:
            return ["a", "b"]
        if "points" in t_low or "points" in e_low:
            return ["points"]
        return ["n"]

    if "dynamic_programming" in concepts or "dynamic-programming" in tags:
        return ["nums"]

    return ["nums"]

def make_test_input(params, is_edge=False):
    d = {}
    for p in params:
        if p in ("s", "word", "pattern", "str", "haystack", "needle", "s1", "s2"):
            d[p] = "" if is_edge else "abc"
        elif p in ("n", "num", "x", "numRows", "k", "val", "target", "m", "capacity", "amount", "targetSum"):
            d[p] = 0 if is_edge else 3
        elif p in ("root", "head", "node", "l1", "l2", "root1", "root2", "p", "q"):
            d[p] = [] if is_edge else [1, 2, 3]
        elif p in ("grid", "matrix", "board", "edges", "connections", "prerequisites", "flights", "intervals"):
            d[p] = [] if is_edge else [[1, 2], [3, 4]]
        elif p in ("accounts",):
            d[p] = [] if is_edge else [["John", "johnsmith@mail.com", "john00@mail.com"]]
        elif p in ("strs", "words", "tokens", "lists", "digits"):
            d[p] = [] if is_edge else ["a", "b"]
        else:
            d[p] = [] if is_edge else [1, 2, 3]
    return json.dumps(d)

def main():
    repo_root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    dataset_path = os.path.join(repo_root, "data", "problems", "dataset_all.json")

    with open(dataset_path, "r", encoding="utf-8") as f:
        problems = json.load(f)

    updated_count = 0
    for p in problems:
        starter = p.get("starter_code", "")
        if "*args" not in starter and "**kwargs" not in starter:
            continue

        entrypoint = p.get("entrypoint", "solve")
        title = p.get("title", entrypoint)
        params = infer_parameters(p)
        param_str = ", ".join(params)

        new_starter = f"def {entrypoint}({param_str}):\n    # Candidate solution for {title}\n    pass\n"
        p["starter_code"] = new_starter

        # Update test cases so input dictionary matches param signatures
        tests = p.get("tests") or []
        if len(tests) >= 1:
            tests[0]["input"] = make_test_input(params, is_edge=False)
            tests[0]["name"] = f"Standard input ({param_str})"
        if len(tests) >= 2:
            tests[1]["input"] = make_test_input(params, is_edge=True)
            tests[1]["name"] = f"Boundary edge input ({param_str})"
        p["tests"] = tests

        updated_count += 1

    with open(dataset_path, "w", encoding="utf-8") as f:
        json.dump(problems, f, indent=2)
        f.write("\n")

    print(f"Successfully updated {updated_count} problems in {dataset_path} with explicit parameter signatures.")

if __name__ == "__main__":
    main()
