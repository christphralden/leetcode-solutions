int min(int* a, int* b){
  return *a < *b ? *a : *b;
}

int maxArea(int* height, int heightsize) {
  int left = 0;
  int right = heightsize - 1;
  
  int ans = 0;
  int newsum = -999;

  while(left < right){
    newsum = min(&height[left], &height[right]) * (right-left);

    if(newsum >= ans){
      ans = newsum;
    }

    (height[left] < height[right]) ? left++ : right--;
  }

  return ans;
}
