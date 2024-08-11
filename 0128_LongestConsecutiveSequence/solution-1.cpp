// Time: 61ms (97.14%) | Memory: 16.04 (38.02%)
int compare(int* a, int* b){
    return (*a - *b);
}

int longestConsecutive(int* nums, int numsSize){
    int count = 0;
    int maxCount =0;
    qsort(nums, numsSize, sizeof(int), compare);

    for(int i = 0; i<numsSize-1; i++){
        int delta = nums[i+1] - nums[i];
        if(delta == 1)count++;
        else if(delta == 0) continue;
        else {
            maxCount = maxCount > count ? maxCount : count;
            count = 0;
        }
    }
    maxCount = maxCount > count ? maxCount : count;
    return numsSize != 0 ? maxCount + 1 : maxCount;
}