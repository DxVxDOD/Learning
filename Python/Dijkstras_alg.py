processed = []
g = {}
g["start"] = {}
g["start"]["a"] = 6
g["start"]["b"] = 2
g["a"] = {}
g["b"] = {}
g["a"]["fin"] = 1
g["b"]["fin"] = 5
g["b"]["a"] = 3
g["fin"] = {}
inf = float("inf")
costs = {}
costs["a"] = 6
costs["b"] = 2
costs["fin"] = inf
p = {}
p["a"] = "start"
p["b"] = "start"
p["fin"] = None


def flc(costs):
    lowest_cost = float("inf")
    lowest_cost_node = None
    for node in costs:
        cost = costs[node]
        if cost < lowest_cost and node not in processed:
            lowest_cost = cost
            lowest_cost_node = node
    return lowest_cost_node


node = flc(costs)
while node is not None:
    cost = costs[node]
    neighbors = g[node]
    for n in neighbors.keys():
        new_cost = cost + neighbors[n]
        if costs[n] > new_cost:
            costs[n] = new_cost
            p[n] = node
        processed.append(node)
        node = flc(costs)
