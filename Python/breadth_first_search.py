from collections import deque
from typing import Optional

graph = {}
graph["me"] = ["alice", "bob", "calire", "Beea"]
graph["Beea"] = ["mango seller"]
graph["bob"] = ["anuj", "peggy"]
graph["alice"] = ["peggy"]
graph["calire"] = ["thom", "jhonny"]
graph["anuj"] = []
graph["peggy"] = []
graph["thom"] = []
graph["jhonny"] = []


def breadth_first_search(name: str):
    seen = []
    search_queue = deque()
    search_queue += graph[name]
    while search_queue:
        person = search_queue.popleft()
        print(person)
        if person not in seen:
            if person == "mango seller":
                return True
            else:
                search_queue += graph[person]
                seen.append(person)
    return False


print(breadth_first_search("me"))


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        res = []
        while head:
            print(res)
            res.append(head.val)
            head = head.next
            if not head.next:
                return res == res.reverse()


print(Solution.isPalindrome([], [1, 2, 2, 1]))
