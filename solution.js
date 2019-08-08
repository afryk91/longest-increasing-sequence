const input = [1, 9, 5, 13, 3, 11, 7, 15, 2, 10, 6, 14, 4, 12, 8, 16];

let subsequenceEndingsIdx = [];
let predecesorsIdx = [];

function lengthOfLongestSubseqEndingWithItem(list, len, item) {
    let left = 1;
    let right = len;
    while (left <= right) {
        let mid = Math.floor((left + right) / 2);
        if (list[subsequenceEndingsIdx[mid]] < item) {   
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return left;
}

function findLIS(list) {
    const longestLength = list.reduce((lengthOfLongest, item, i) => {        
        let newLength = lengthOfLongestSubseqEndingWithItem(list, lengthOfLongest, item);
        
        predecesorsIdx[i] = subsequenceEndingsIdx[newLength - 1];
        subsequenceEndingsIdx[newLength] = i;

        return newLength > lengthOfLongest ? newLength : lengthOfLongest;
    }, 0);

    let subseq = [];
    let k = subsequenceEndingsIdx[longestLength];

    for (var j = longestLength - 1; j >= 0; j--) {
        subseq[j] = list[k];
        k = predecesorsIdx[k];
    }

    return subseq;
}

console.log(findLIS(input));
