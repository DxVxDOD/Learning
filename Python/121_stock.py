def best_time(arr):
    i = 0
    m = 0
    for j in range(1, len(arr)):
        m = max(m, arr[i] - arr[j])
        print(m)
        if i > j:
            i += 1
    return m


best_time([7,1,5,3,6,4])
print(best_time([7,1,5,3,6,4]))
