// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Task1 {
    mapping(uint => string) private intToRomanMap;
    mapping(bytes1 => uint) private romanToIntMap;
    uint[] private digitalNum;
    constructor() {
        intToRomanMap[1] = 'I';
        intToRomanMap[5] = 'V';
        intToRomanMap[10] = 'X';
        intToRomanMap[50] = 'L';
        intToRomanMap[100] = 'C';
        intToRomanMap[500] = 'D';
        intToRomanMap[1000] = 'M';

        
        romanToIntMap['I'] = 1;
        romanToIntMap['V'] = 5;
        romanToIntMap['X'] = 10;
        romanToIntMap['L'] = 50;
        romanToIntMap['C'] = 100;
        romanToIntMap['D'] = 500;
        romanToIntMap['M'] = 1000;
    }

    // 2. 反转一个字符串。输入 "abcde"，输出 "edcba"
    function reverseString(string memory input) public pure returns (string memory) {
        bytes memory bytesInput = bytes(input);
        bytes memory reversedBytes = new bytes(bytesInput.length);
        for (uint i = 0; i < bytesInput.length; i++) {
            reversedBytes[i] = bytesInput[bytesInput.length - 1 - i];
        }
        return string(reversedBytes);
    }

    // 3. 用 solidity 实现整数转罗马数字
    function intToRoman(uint input) public returns (string memory romanStr) {
        romanStr = "";
        // 首先计算位数
        uint temp = input;
        // if (digitalNum[0] > 0) {
        //     delete digitalNum;
        // }
        while (temp > 0) {
            digitalNum.push(temp % 10);
            temp /= 10;
        }
        // 然后处理每一位
        for (uint i = 0; i < digitalNum.length; i++) {
            uint oneNum = digitalNum[i] % 5;
            if (digitalNum[i] < 5) {
                if (digitalNum[i] == 4) {
                    romanStr = string.concat(intToRomanMap[10 ** i], intToRomanMap[5* 10 ** i], romanStr);
                } else {
                    while(oneNum > 0) {
                        romanStr = string.concat(intToRomanMap[10 ** i], romanStr);
                        oneNum--;
                    }
                }
            } else {
                if (digitalNum[i] == 9) {
                    romanStr = string.concat(intToRomanMap[10 ** i], intToRomanMap[10 ** (i + 1)], romanStr);
                } else {
                    while(oneNum > 0) {
                        romanStr = string.concat(intToRomanMap[10 ** i], romanStr);
                        oneNum--;
                    }
                    romanStr = string.concat(intToRomanMap[5 * 10 ** i], romanStr);
                }
            }
        }
        delete digitalNum;
        return romanStr;
    }

    // 4. 用 solidity 实现罗马数字转整数
    function romanToInt(string memory romanNumeral) public view returns (uint) {
        uint ret = 0;
        bytes memory bytesRomanNumeral = bytes(romanNumeral);
        for (uint i = 0; i < bytesRomanNumeral.length; i++) {
            if (i < bytesRomanNumeral.length - 1 
                && (romanToIntMap[bytesRomanNumeral[i]] < romanToIntMap[bytesRomanNumeral[i + 1]])){
                ret += romanToIntMap[bytesRomanNumeral[i + 1]] - romanToIntMap[bytesRomanNumeral[i]];
                i++;
            }
            else {
                ret += romanToIntMap[bytesRomanNumeral[i]];
            }
        }
        return ret;
    }

    // 5. 合并两个有序数组 (Merge Sorted Array)
    function mergeArr(int[] memory nums1, uint m, int[] memory nums2, uint n) public pure returns(int[] memory) {
        int[] memory ret = new int[](m + n);
        uint i = 0;
        uint j = 0;
        uint k = 0;
        while (i < m && j < n) {
            if (nums1[i] <= nums2[j]) {
                ret[k] = nums1[i];
                i++;
            } else {
                ret[k] = nums2[j];
                j++;
            }
            k++;
        }
        while (i < m) {
            ret[k] = nums1[i];
            k++;
            i++;
        }
        while (j < n) {
            ret[k] = nums2[j];
            k++;
            j++;
        }
        return ret;
    }
    
    // 6. 二分查找 (Binary Search)，在一个有序数组中查找目标值
    function binarySearch(int[] memory nums, int target) public pure returns(int) {
        uint left = 0;
        uint right = nums.length - 1;
        while (left <= right) {
            uint mid = (left + right) / 2;
            if (nums[mid] == target) {
                return int(mid);
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return -1;
    }
}