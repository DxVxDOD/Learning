from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        res = []
        q = [root]

        while len(q):
            curr = q.pop(0)

            if not curr:
                continue

            if curr.val:
                diff = k - curr.val

                if diff in res:
                    return True
                res.append(diff)

            q.append(curr.left)
            q.append(curr.right)

        return False


print(Solution.findTarget(0,[5, 3, 6, 2, 4, 0, 7], 9))
