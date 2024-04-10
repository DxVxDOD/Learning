from typing import List
from collections import deque


def test(students: List[int], food: List[int]):
    deq = deque(food)
    students.sort()
    food.sort()
    count = 0
    print(students)
    print(food)
    for i in range(len(students)):
        if students[i] != food[i]:
            count += 1
    return count


print(test([1,1,1,0,0,1], [1,0,0,0,1,1]))
