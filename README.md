## 说明
* 十三水的AI出牌算法，速度非常快，一般在微秒级，附带测试方案
* 目前该算法基于golang实现，后续将实现js和lua版本
* 算法基于常归玩法，即52张不带大小王的情况开发

## 优势
* 通常的十三水AI算法都是穷举法,即13张牌遍历后选出5张，再依次检测是不是特定牌型，速度较慢
* 本方案反其道页行，从任意张牌中，选出特殊的牌型的5张，划到左节点，再将右节点递归分拆出特定的牌型，再复杂的牌计算一般都在毫秒级，部分在微秒级  
* 这个十三水的算法基于二叉树的思想，将13张牌先拆出5张特定组合来，再将剩下的牌再拆出5张特定的牌型出来
* 该算法会选出大量组合方案，然后横向比较，选出最优解
