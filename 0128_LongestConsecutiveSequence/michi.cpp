#include <stdio.h>
#include <stdlib.h>

void swap(int* a, int* b) {
    int t = *a;
    *a = *b;
    *b = t;
}
int compare(int* a, int* b){
    return (*a - *b);
}

int partition(int arr[], int low, int high) {
    int pivot = arr[high];
    int i = low - 1;

    for (int j = low; j <= high - 1; j++) {
        if (arr[j] < pivot) {
            i++;
            swap(&arr[i], &arr[j]);
        }
    }
    swap(&arr[i + 1], &arr[high]);
    return i + 1;
}

void quickSort(int arr[], int low, int high) {
    if (low < high) {
        int pi = partition(arr, low, high);

        quickSort(arr, low, pi - 1);
        quickSort(arr, pi + 1, high);
    }
}

void merge(int arr[], int l, int m, int r)
{
    int i, j, k;
    int n1 = m - l + 1;
    int n2 = r - m;

    int L[n1], R[n2];

    for (i = 0; i < n1; i++)
        L[i] = arr[l + i];
    for (j = 0; j < n2; j++)
        R[j] = arr[m + 1 + j];

    i = 0;
    j = 0;
    k = l;
    while (i < n1 && j < n2) {
        if (L[i] <= R[j]) {
            arr[k] = L[i];
            i++;
        }
        else {
            arr[k] = R[j];
            j++;
        }
        k++;
    }

    while (i < n1) {
        arr[k] = L[i];
        i++;
        k++;
    }

    while (j < n2) {
        arr[k] = R[j];
        j++;
        k++;
    }
}


void mergeSort(int arr[], int l, int r)
{
    if (l < r) {
        int m = l + (r - l) / 2;

        mergeSort(arr, l, m);
        mergeSort(arr, m + 1, r);

        merge(arr, l, m, r);
    }
}

int longestConsecutive(int* nums, int numsSize){
    int count = 0;
    int maxCount =0;
    // mergesort(nums, 0, numsSize-1);
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

int main(){
    int nums[12] = {1,2,0,1};
    int output = 3;

    int result = longestConsecutive(nums, 4);
    printf("\n%d\n", result);
}