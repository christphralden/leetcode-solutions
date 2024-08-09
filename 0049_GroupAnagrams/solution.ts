// Time: 107ms (86.14%) | Memory: 64.78 (19.82%)
function groupAnagrams(strs: string[]): string[][] {
    let map = new Map<string, string[]>()
    let res:string[][] = []
    let i = 0, len = strs.length
    while(i < len){
        const sorted = strs[i].split("").sort().join("")
        const exist = map.get(sorted) ?? []
        exist.push(strs[i])
        map.set(sorted, exist)
        i++
    }

    for(let entry of map.entries()){
        res.push(entry[1])
    }

    return res
};