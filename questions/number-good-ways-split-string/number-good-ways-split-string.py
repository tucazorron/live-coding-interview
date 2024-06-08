def solution(s):
    lenstr = len(s)
    if lenstr == 1:
        return 0
    if lenstr == 2:
        return 1
    
    first, last = {}, {}
    
    for index, char in enumerate(s):
        if char not in first:
            first[char] = index
        last[char] = index
        
    indices = list(first.values()) + list(last.values())
    indices.sort()
    
    mid = len(indices) // 2
    
    return indices[mid] - indices[mid-1]