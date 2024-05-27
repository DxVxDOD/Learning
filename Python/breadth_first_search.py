from collections import deque

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


def bfs(head, needle, graph):
    q = deque()
    q += graph[head]
    visited = []
    while q:
        curr = q.popleft()
        if curr not in visited:
            if curr == needle:
                return True
            else:
                q += graph[curr]
                visited.append(curr)
    return False


print(bfs("me", "jhonny", graph))
