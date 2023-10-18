class Solution(object):
    def twoSum(self, nums, target):
        for i in range (len(nums) - 1):
            for j in range (i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    new_list = [i, j]
                    return new_list