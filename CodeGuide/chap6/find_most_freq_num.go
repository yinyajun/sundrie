package main

// 只用2GB内存，在20亿个整数中找到出现次数最多的数

// 有一个大文件，里面包含了20亿个全是32bit的整数，寻找其中出现次数最多的数。

// 使用hash table做词频统计。key和value使用int32（4B）即可。
// 一条记录需要8B，那么2亿条记录需要1.6GB内存

// 如果key特别多，很容易就会发生超出内存限制。

// 解决方法：
// 将大文件用hash function分成16个小文件
// 对每一个小文件用hash table统计词频
// 选出16个小文件各自的第一名中谁出现的次数最多即可

// 把一个大的集合通过哈希函数分配到多台机器中，或者分配到多个文件里，这种技巧是处理大数据的最常用技巧之一。

// ---------------------------------------------------
// 40亿个非负整数中找到没出现的数

// uint32的范围是0~42亿多，现在有一个包含40亿个无符号整数的文件，在这个范围内必然有没出现的数，最多使用1GB内存
// hash table: 40亿 * 4B = 16GB(不符合)
// bitmap： 2^32bit / 8 = 500M

// 进阶： 内存只给10M
// 将0~2^32-1 划分为 64个区间，每个区间2^26大小
// 如果某个区间的没有2^26个数，那么未出现的数就在这个区间中

// 先使用64大小的数组表征区间，并记录区间的数目
// 遍历每一个数num，num/(2^26)为区间，对应区间数目+1

//找到对应区间后，再次遍历这个区间，使用bitmap，大概 2^23 = 8M bit

// 1. 根据10M限制，确定统计区间大小（第二次遍历时arr大小）
// 2. 利用区间统计的方式，找到那个计数不足的区间
// 3. 对这个区间上的数做bitmap映射，再遍历一次bitmap

// ---------------------------------------------------

//100亿个url中重复url， 每个url占64B，找出所有的重复url
// 百亿数据量，求每天最热的top100

// 先拆到100台机器上，通过hash函数拆，保证一条url只会分配到一台机器上
// 如果分配到每台机器上数目依然很大，用hash 函数拆成多个文件

// 处理每个小文件的时候，通过100大小的小顶堆来求小文件的top100
// 将每个小文件的top100 进行外排序或者继续使用小顶堆

// 机器间也是如此

// ---------------------------------------------------

//40亿个uint32非负整数中，最多使用1GB内存，找出所有出现了两次的数。

//仍然采用bitmap的方式，由于需要找词频为2的数
//
//所以bitmap设计的时候，每个位置分配 2个bit，可记录00，01，02，03
// 遍历bitmap，找到所有02的数

// 补充：最多使用10M，找到40亿个整数的中位数？

// 长度为2M的数组占据空间： 2M*4Byte = 8M Byte
// 将所有uint32，按照2M个数划分区间，一共可以划分2148个区间。
// 用一个长度为2148的数组记录每个区间的hist

// 这样就可以找到中位数所在的区间

// 找到这个区间以后，在申请2M的数组，用来记录该区间的hist
// 当然，再次遍历所有num，不过只记录该区间的数，其他数省略

// ---------------------------------------------------
// ---------------------------------------------------
