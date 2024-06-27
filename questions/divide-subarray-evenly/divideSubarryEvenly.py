def solution(nums):
    print("\ninput array:", nums)

    print("\ndoes the array have less than 3 elements?")
    if len(nums) < 3:
        print("=> yes")
        return False
    print("=> no")

    numsToFind = set()
    print("\nvalues i want to find:", numsToFind)

    pastNums = set()
    print("values i already found:", pastNums)

    candidate = sum(nums)
    print("first candidate (total sum):", candidate)

    print("\nfor each element in the array")
    for num in nums:
        print("\nelement:", num)

        print("am i looking for this element?")
        if num in numsToFind:
            print("=> yes")
            print("element to remove:", num)
            return True
        print("=> no")

        candidate -= 2*num
        print("new candidate:", candidate)

        print("have i already found the negative candidate in my array?")
        if -candidate in pastNums:
            print("=> yes")
            print("element to remove:", -candidate)
            return True
        print("=> no")

        numsToFind.add(candidate)
        print("values i want to find:", numsToFind)

        pastNums.add(num)
        print("values i already found:", pastNums)

    print("cant remove an element from the array and evenly divide the new array into two subarrays with equal sum")
    return False

# print(solution([1, 2, 3, 3, 3, 20]))
# print(solution([8, 1, 2, 3]))
# print(solution([]))
# print(solution([1]))
# print(solution([1, 1]))
print(solution([9, -8, 3, 6, -1, -10]))
# print(solution([4, 5, 7, 9]))
# print(solution([0, 1, 2, 3, 2, 1, 0]))
# print(solution([1, 2, 4, 8]))
