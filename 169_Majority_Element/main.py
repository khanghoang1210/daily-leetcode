
# Time: O(N) / Space: O(N)
def majorityElement(nums) -> int:
    map = {}
    for i in nums:
        map[i] = map.setdefault(i, 0) + 1
    
    max = 0
    res = 0
    for k, v in map.items():
        if max < v:
            max = v
            res = k
    return res

# Test Case
nums = [2,2,1,1,1,2,2]
print(majorityElement(nums))