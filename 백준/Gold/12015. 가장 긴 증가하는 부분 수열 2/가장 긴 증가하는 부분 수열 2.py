n = int(input())
arr = list(map(int, input().split()))

LIS = [arr[0]]

def binarySearch(e):
    l, r = 0, len(LIS)-1
    while l <= r:
        mid = (l+r)//2
        if LIS[mid] == e:
            return mid
        elif LIS[mid] < e:
            l = mid + 1
        else:
            r = mid - 1
    return l

for i in arr:
    if LIS[-1] < i:
        LIS.append(i)
    else:
        idx = binarySearch(i)
        LIS[idx] = i

print(len(LIS))