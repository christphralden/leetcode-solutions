class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        std::unordered_map<int, int> map;
        map.reserve(nums.size());

        for (int i = 0; i < nums.size(); ++i) {
            int diff = target - nums[i];
            auto it = map.find(diff);
            if (it != map.end()) {
                return {it->second, i};
            }
            map[nums[i]] = i;
        }
        return {}; 
    }
};
