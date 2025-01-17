## 说明
* 十三水的AI出牌算法，速度非常快，一般在微秒级，附带测试方案
* 目前该算法基于golang实现，后续将实现js和lua版本
* 算法基于常归玩法，即52张不带大小王的情况开发

## 优势
* 通常的十三水AI算法都是穷举法，即13张牌遍历出5张，再依次检测是不是特定牌型，速度较慢，因为13张遍历出5张就有（13x12x11x10x9）种组合
* 本算法基于二叉树的思想，将13张牌先拆出5张特定组合来，再将剩下的牌再拆出5张特定的牌型出来
* 具体步骤1：从N张牌中，选出特定的牌型的5张，归到左节点；再将剩余的牌归档到右节点，
* 具体步骤2：将步骤1的右节点递归上述步骤1，这样拆成的3组节点，组成牌型组合
* 具体步骤3：将上述步骤中得到的所有牌型组合进行横向比较，一定是左节点大于右节点的情况，才算正确的牌型组合
* 具体步骤4：将所有正确的牌型组合中，进行纵向比较，选出最优解  
* 本算法速度较快：特殊牌型的计算一般在微秒级，普通牌型因组合较多，计算一般都在10毫秒级，部分较简单的在微秒级  
* 该算法会选出大量组合方案，然后横向比较，选出最优解
* 项目中有测试示例，基本涵盖所有牌型的组合

## 比较
* tradition_func包为传统方法，即最常见的遍历法
* algorithm_func包为新的算法方法，为最新节约时间的方法，
* 可以在main函数中，用两种不同的方法进行速度对比，部份效率有近100倍提升

## 部分测试示例
```
测试组合=青龙测试,牌型={♦A,♦2,♦3,♦4,♦5,♦6,♦7,♦8,♦9,♦T,♦J,♦Q,♦K},开始时间=739042000
2021/06/18 22:31:26 该牌型为特殊牌型：至尊青龙
结束时间=739196000,AI算法耗时【154】微秒

测试组合=一条龙,牌型={♣A,♦2,♦3,♦4,♦5,♦6,♦7,♦8,♦9,♦T,♦J,♦Q,♦K},开始时间=739207000
2021/06/18 22:31:26 该牌型为特殊牌型：一条龙
结束时间=739225000,AI算法耗时【18】微秒

测试组合=十二皇族,牌型={♠A,♦T,♣T,♥T,♦J,♣J,♥J,♦Q,♣Q,♥Q,♦K,♣K,♥K},开始时间=739231000
2021/06/18 22:31:26 该牌型为特殊牌型：十二皇族
结束时间=739247000,AI算法耗时【16】微秒

测试组合=三同花顺,牌型={♦2,♦3,♦4,♥3,♥4,♥5,♥6,♥7,♣T,♣J,♣Q,♣K,♣A},开始时间=739253000
2021/06/18 22:31:26 该牌型为特殊牌型：三同花顺
结束时间=739270000,AI算法耗时【17】微秒

测试组合=三分天下,牌型={♦3,♣3,♥3,♠3,♦5,♣5,♥5,♠5,♦7,♣7,♥7,♠7,♦9},开始时间=739278000
2021/06/18 22:31:26 该牌型为特殊牌型：三分天下
结束时间=739291000,AI算法耗时【13】微秒

测试组合=全大,牌型={♦8,♣9,♥T,♠Q,♦K,♣8,♥9,♠A,♦A,♣T,♥T,♠K,♦Q},开始时间=739297000
2021/06/18 22:31:26 该牌型为特殊牌型：全大
结束时间=739311000,AI算法耗时【14】微秒

测试组合=全小,牌型={♦2,♣3,♥4,♠5,♦6,♣7,♥2,♠3,♦4,♣5,♥6,♠7,♦8},开始时间=739316000
2021/06/18 22:31:26 该牌型为特殊牌型：全小
结束时间=739332000,AI算法耗时【16】微秒

测试组合=凑一色,牌型={♦2,♦3,♦5,♦6,♦7,♦T,♦K,♥2,♥6,♥9,♥A,♥T,♥J},开始时间=739338000
2021/06/18 22:31:26 该牌型为特殊牌型：凑一色
结束时间=739353000,AI算法耗时【15】微秒

测试组合=四套三条,牌型={♦3,♣3,♥3,♦5,♣5,♥5,♦7,♥7,♠7,♦9,♣9,♥9,♠T},开始时间=739359000
2021/06/18 22:31:26 该牌型为特殊牌型：四套三条
结束时间=739372000,AI算法耗时【13】微秒

测试组合=五对三条,牌型={♦3,♣3,♦4,♣4,♦6,♣6,♦7,♣7,♦9,♣9,♦T,♣T,♥T},开始时间=739384000
2021/06/18 22:31:26 该牌型为特殊牌型：五对三条
结束时间=739397000,AI算法耗时【13】微秒

测试组合=六对半,牌型={♦3,♣3,♦4,♣4,♦6,♣6,♦7,♣7,♦9,♣9,♦T,♣T,♥K},开始时间=739404000
2021/06/18 22:31:26 该牌型为特殊牌型：六对半
结束时间=739423000,AI算法耗时【19】微秒

测试组合=三同花,牌型={♦2,♦5,♦6,♥3,♥4,♥K,♥A,♥7,♣T,♣2,♣7,♣K,♣A},开始时间=739433000
2021/06/18 22:31:26 该牌型为特殊牌型：三同花
结束时间=739455000,AI算法耗时【22】微秒

测试组合=三顺子1：最简单的顺子情况,牌型={♠2,♥3,♦4,♥3,♥4,♥5,♣6,♥7,♣T,♣J,♣Q,♦K,♣A},开始时间=739461000
2021/06/18 22:31:26 该牌型为特殊牌型：三顺子
结束时间=739509000,AI算法耗时【48】微秒

测试组合=三顺子2：两叠顺子情况,牌型={♠J,♥Q,♦K,♥3,♥4,♥5,♣6,♥7,♣3,♣4,♥5,♥6,♦7},开始时间=739516000
2021/06/18 22:31:26 该牌型为特殊牌型：三顺子
结束时间=739538000,AI算法耗时【22】微秒

测试组合=三顺子3：A2345的情况,牌型={♠J,♥Q,♦K,♥3,♥4,♥5,♣6,♥7,♣A,♣2,♥3,♥4,♦5},开始时间=739546000
2021/06/18 22:31:26 该牌型为特殊牌型：三顺子
结束时间=739593000,AI算法耗时【47】微秒

测试组合=三顺子4：头部为A23也是三顺子,牌型={♠A,♥2,♦3,♥3,♥4,♥5,♣6,♥7,♣A,♣2,♥3,♥4,♦5},开始时间=739598000
2021/06/18 22:31:26 该牌型为特殊牌型：三顺子
结束时间=739642000,AI算法耗时【44】微秒

测试组合=普通牌型【4对+5单】,牌型={♦A,♣A,♥Q,♠Q,♣3,♦3,♣4,♥4,♠5,♥6,♥J,♠8,♠9},开始时间=741825000
2021/06/18 22:31:26 普通牌型[0] = {左:【两对】= {♦3♣3♣4♥4♠5},中:【对子】= {♥6♠8♠9♦A♣A},右:【对子】= {♥J♥Q♠Q},好牌值=【5】}
2021/06/18 22:31:26 最好牌型 =  {左:【两对】= {♦3♣3♣4♥4♠5},中:【对子】= {♥6♠8♠9♦A♣A},右:【对子】= {♥J♥Q♠Q},好牌值=【5】}
结束时间=741885000,AI算法耗时【60】微秒

```

