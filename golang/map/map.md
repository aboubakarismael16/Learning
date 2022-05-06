## 1. map struct

Golang's map uses hash table as the underling implemention. There can be multiple hash table nodes in a hash table, that is buckets, and each bucket is saved one or a set of key value pairs in hte map.

map struct is set on `runtime/map/map.go/hmap`

```go
type hmap struct{
    count     int// 当前保存的元素个数
    ...
    B         uint8  // 指示bucket数组的大小
    ...
    buckets    unsafe.Pointer// bucket数组指针，数组的大小为2^B
    ...
}
```

![map01](/img/map01.jpg)

In this example , `hmap.B = 2`, while the length of hmap.buckets is 2^B egal to 4 . After th hash operation , the element will fall into a bucket for storage .The search process is similar. 

`bucket` is often translated into buckets. The so-called `hash bucket` is actually bucket.



## 2. bucket struct

bucket struct is ser on `runtime/map.go/bmap`:

```go
type bmap struct{
    tophash [8]uint8  //存储哈希值的高8位
    data    byte[1]   //key value数据:key/key/key/.../value/value/value...
    overflow *bmap    //溢出bucket的地址
}
```

Each bucket can store 8 key-value pairs.

- `tophash` is an array with a length of 8. When the keys with the same hash value (to be precise, the keys with the same low bits of the hash value) are stored in the current bucket, the high bits of the hash value will be stored in the array to facilitate subsequent matching.
- `data` area stores key-value data, and the storage order is key/key/key/...value/value/value. This storage is to save space waste caused by byte alignment.
- `overflow` pointer points to the next bucket, according to which all conflicting keys are connected

**Note : `data` and `overflow` in the above are not defined in the structure but accessed directly through pointer operations.**

![map02](/home/aboubakar/Pictures/map02.jpg)

## 3. Hash conflict

When two or more keys are hashed to the same bucket, we say that these keys conflict. Go uses the chain address method to resolve key conflicts. Since each bucket can store 8 key-value pairs, when the same bucket stores more than 8 key-value pairs, another key-value pair will be created and the buckets will be connected in a similar way to a linked list.

![map03](/home/aboubakar/Pictures/map03.jpg)



The bucket data structure indicates that the pointer to the next bucket is called an overflow bucket, which means the overflow bucket that is overflowing from the current bucket. In fact, hash conflict is not a good thing. It reduces the access efficiency. A good hash algorithm can ensure the randomness of the hash value, but too many conflicts must be controlled, which will be described in detail later.



## 4. Load factor

The load factor is used to measure the conflict of a hash table, the formula is:

```go
Load factor = number of keys / number of buckets
```

For example, for a hash table with 4 buckets and 4 key-value pairs, the load factor of this hash table is 1.

The hash table needs to control the load factor to an appropriate size, and it needs to be rehashed if it exceeds its threshold, that is, the key-value pairs are reorganized :

- The hash factor is too small, indicating low space utilization.
- The hash factor is too large, indicating that the conflict is serious and the access efficiency is low.

Each hash table implementation has a different tolerance for load factors. For example, in Redis implementation, rehash will be stimulate when the load factor is greater than 1, while Go will stimulate rehash when the load factor reaches 6.5, because each bucket of Redis only can store 1 key-value pair, and Go’s bucket may store 8 key-value pairs, so Go can tolerate a higher load factor.

## 5. Progressive expansion

### 5.1. Prerequisites for expansion

In order to ensure access efficiency, when new elements are about to be added to the map, they will check whether they need to be expanded. Expansion is actually a means of changing space for time. There are two conditions for stimulating expansion :

- When the load factor > 6.5, that is the average key-value pair stored in each bucket reaches 6.5.
- When the number of overflows > 2^15, that is when the number of overflows exceeds 32768.

### 5.2. Incremental expansion

When the load factor is too large, a new bucket is created, the length of the new bucket is twice the original, and then the data of the old bucket is moved to the new bucket. Considering that if the map stores hundreds of millions of key-values, a one-time relocation will cause a relatively large delay. Go adopts a gradual relocation strategy, that is a relocation is stimulated every time the map is accessed, and 2 keys are relocated each time value pair.

The following figure shows a map that contains a bucket full of loads (for the convenience of description, the value area of the bucket is omitted in the figure) :



![map04](/home/aboubakar/Pictures/map04.jpg)

The current map stores 7 key-value pairs and only 1 bucket. The load factor here is 7. When data is inserted again, an expansion operation will be stimulated. After the expansion, the newly inserted key will be written into the new bucket.

When the 8 key-value pair is inserted, expansion will be stimulate. The diagram after expansion is as follows :



![map05](/home/aboubakar/Pictures/map05.jpg)

The oldbuckets member in the hmap data structure refers to the original bucket, and the buckets point to the newly applied bucket. The new key-value pair is inserted into the new bucket. Subsequent access operations to the map will stimulate the migration, and the key-value pairs in the oldbuckets will be gradually migrated. When all the key-value pairs in oldbuckets are relocated, delete oldbuckets.

The figure after the relocation is as follows :

![map06](/home/aboubakar/Pictures/map06.jpg)

During the data relocation process, the key-value pairs in the original bucket will exist in front of the new bucket, and the newly inserted key-value pairs will exist behind the new bucket. The actual relocation process is more complicated and will be introduced in detail in the subsequent source code analysis.

