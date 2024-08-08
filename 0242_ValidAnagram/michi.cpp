#include <stdio.h>
#include<string.h>

bool isAnagram(char* s, char* t){
    int map[27] = {0};

    for(int i=0; s[i] != '\0'; i++){
        map[s[i]-'a']++;
    }

    for(int i=0; t[i] != '\0'; i++){
        map[t[i]-'a']--;
    }

    for(int i=0; i<27; i++){
        if(map[i]!=0) return false;
    }
    return true;
}

int main()
{
    char s[10] = "rat";
    char t[10] = "car";

    printf("%s\n", isAnagram(s, t) ? "true" : "false");
}