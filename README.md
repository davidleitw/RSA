# RSA 心得

紀錄一些學習密碼學的心得，並用 openssl 搭配 gin 寫一點簡單的範例。


## 非對稱加密(公開金鑰加密)

RSA加密演算法是一種非對稱加密演算法，非對稱加密的特色在於它需要兩個金鑰，一個是公開密鑰，另一個是私有密鑰；公鑰用作加密，私鑰則用作解密。使用公鑰把明文加密後所得的密文，只能用相對應的私鑰才能解密並得到原本的明文，最初用來加密的公鑰不能用作解密。

![](https://i.imgur.com/EjFAlMK.png)


由於加密和解密需要兩個不同的金鑰，故被稱為**非對稱加密**。
> 公鑰加密，私鑰解密

在$1976$年以前，還沒有發展出非對稱加密的技術，主要加密方式都是採取**對稱式加密**

簡單來說就是寄信者跟收信者都使用相同的金鑰或者是使用兩個可以簡單地相互推算的金鑰來做加密/解密的動作。

假設今天我們的演算法是用同一個金鑰，這種對稱式加密最大的缺點在於，我們要怎麼把金鑰(對訊息處理的規則)傳給對方，在網路的世界裡面隨時都有被監聽的風險，要怎麼進行安全的傳輸金鑰變成一個很大的問題，所以為了解決這個問題，科學家提出了非對稱式加密的概念。

## RSA

> RSA加密演算法是一種非對稱加密演算法，在公開金鑰加密和電子商業中被廣泛使用。RSA是由羅納德·李維斯特（Ron Rivest）、阿迪·薩莫爾（Adi Shamir）和倫納德·阿德曼（Leonard Adleman）在1977年一起提出的。當時他們三人都在麻省理工學院工作。RSA 就是他們三人姓氏開頭字母拼在一起組成的。

![](https://i.imgur.com/lE7U7Ky.png)

[圖片來源](https://www.techapple.com/archives/25855)

### 金鑰計算方式
- 選出兩個較大的質數 ![](https://render.githubusercontent.com/render/math?math=p), ![](https://render.githubusercontent.com/render/math?math=q)
- 計算兩個質數的乘積 ![](https://render.githubusercontent.com/render/math?math=n\=p*q)
- 計算出小於 n 且與 n 互質的整數個數
   
  ![](https://render.githubusercontent.com/render/math?math=\varphi(n)=(p-1)*(q-1))
- 選擇一個整數 **e**(拿來當作公鑰)
    - 選擇條件
        - ![](https://latex2image-output.s3.amazonaws.com/img-heUvAk9X.svg)
        - ![](https://latex2image-output.s3.amazonaws.com/img-VRHdeXUh.svg)互質

- 計算![](https://latex2image-output.s3.amazonaws.com/img-D6h1FGmQ.svg)相對於![](https://latex2image-output.s3.amazonaws.com/img-YS3FV8Jy.svg)的模反元素![](https://latex2image-output.s3.amazonaws.com/img-5VfEC4JX.svg)拿來當作私鑰
  
  ![](https://latex2image-output.s3.amazonaws.com/img-rNB5W1k7.svg)

  所以可以得出

  ![](https://latex2image-output.s3.amazonaws.com/img-7sZ11Wd4.svg)

  移項得到

  ![](https://latex2image-output.s3.amazonaws.com/img-S7BEqV2x.svg)

  // 待補，接著需要使用擴展歐幾里得算法


![](https://i.imgur.com/im4zugs.png)
[來源](https://ithelp.ithome.com.tw/articles/10250721)

經過上述求金鑰的過程，可以得到
- 公鑰 ![](https://latex2image-output.s3.amazonaws.com/img-MWWWYstf.svg)
- 私鑰 ![](https://latex2image-output.s3.amazonaws.com/img-qJFBdKjw.svg)

### 議題: 如何選擇質數

先討論只選擇一個質數的情況，假設我們今天要一個 ![](https://latex2image-output.s3.amazonaws.com/img-dC9QCq81.svg) 位元的質數，我們可以隨機挑一個 ![](https://latex2image-output.s3.amazonaws.com/img-dC9QCq81.svg) 位元的奇數，然後使用質數判斷法來確認隨機選取的數是不是質數。如果不是質數則在重新選取一次。
我一開始看到這種作法會認為隨機挑選應該是很沒有效率的作法，後來查了一些資料，質數的佔比其實比想像中的還要多，詳細數據可以參考[質數計算函數](https://zh.wikipedia.org/wiki/%E7%B4%A0%E6%95%B0%E8%AE%A1%E6%95%B0%E5%87%BD%E6%95%B0)。

如果使用隨機挑選的作法，效能瓶頸就最有可能出現在判斷一個數是不是質數這個動作，參考了一篇文章[\[8\]](https://www.zhihu.com/question/54779059)裡面用 Java 的 [Bouncy Castle lib(一個密碼學相關的函式庫)](https://github.com/bcgit/bc-java) 作為舉例。

```java
public BigInteger(int bitLength, int certainty, Random rnd) {
    BigInteger prime;

    if (bitLength < 2)
        throw new ArithmeticException("bitLength < 2");
    prime = (bitLength < SMALL_PRIME_THRESHOLD
                            ? smallPrime(bitLength, certainty, rnd)
                            : largePrime(bitLength, certainty, rnd));
    signum = 1;
    mag = prime.mag;
}

// Minimum size in bits that the requested prime number has
// before we use the large prime number generating algorithms.
// The cutoff of 95 was chosen empirically for best performance.
private static final int SMALL_PRIME_THRESHOLD = 95;
```

一般為了追求效率，會根據給定的 bit 數來決定使用什麼方法來確定隨機產生的數是不是質數。

```java
// smallPrime function
// Do expensive test if we survive pre-test (or it's inapplicable)
if (p.primeToCertainty(certainty, rnd))
    return p;
```

```java
/**
 * Find a random number of the specified bitLength that is probably prime.
 * This method is more appropriate for larger bitlengths since it uses
 * a sieve to eliminate most composites before using a more expensive
 * test.
 */
    private static BigInteger largePrime(int bitLength, int certainty, Random rnd) {
    BigInteger p;
    p = new BigInteger(bitLength, rnd).setBit(bitLength-1);
    p.mag[p.mag.length-1] &= 0xfffffffe;

    // Use a sieve length likely to contain the next prime number
    int searchLen = getPrimeSearchLen(bitLength);
    BitSieve searchSieve = new BitSieve(p, searchLen);
    BigInteger candidate = searchSieve.retrieve(p, certainty, rnd);

    while ((candidate == null) || (candidate.bitLength() != bitLength)) {
        p = p.add(BigInteger.valueOf(2*searchLen));
        if (p.bitLength() != bitLength)
            p = new BigInteger(bitLength, rnd).setBit(bitLength-1);
        p.mag[p.mag.length-1] &= 0xfffffffe;
        searchSieve = new BitSieve(p, searchLen);
        candidate = searchSieve.retrieve(p, certainty, rnd);
    }
    return candidate;
}
```
有興趣可以參考看看自己常用語言的函式庫，通常作法都大同小異，會根據語言特性而做一些不同的優化。
通常底層還是會使用[Miller–Rabin primality test](https://zh.wikipedia.org/wiki/%E7%B1%B3%E5%8B%92-%E6%8B%89%E5%AE%BE%E6%A3%80%E9%AA%8C)來作為驗證質數的方法，而不是採用類似[AKS](https://en.wikipedia.org/wiki/AKS_primality_test)這種確定性演算法，主要原因還是在於速度上有很大的差異，而且 Miller-Rabin 如果經過多次的驗證，其可靠性已經足夠了。 還有看到一些說法好像是因為硬體的計算因素[\[9\]](https://crypto.stackexchange.com/questions/71/how-can-i-generate-large-prime-numbers-for-rsa)導致，有興趣可以看看。



### 選擇滿足 RSA 安全性的質數
上面只討論了如何隨機產生質數，但是 RSA 的演算法中包含著兩個質數 ![](https://latex2image-output.s3.amazonaws.com/img-Gq8PVexY.svg)，所以在選擇質數上會有一些額外的限制來確保其安全性。

- RSA中質數![](https://latex2image-output.s3.amazonaws.com/img-Gq8PVexY.svg)不能距離太接近
如果![](https://latex2image-output.s3.amazonaws.com/img-Gq8PVexY.svg)距離太近，會有快速算法將 N 分解，一般來說如果 N 的位數為 n，那麼![](https://latex2image-output.s3.amazonaws.com/img-DEaUM6vn.svg)要滿足
    - ![](./image/pq12.svg)
    - ![](./image/pq11.svg)
    
    
## Reference
1. [How to better generate large primes: sieving and then random picking or random picking and then checking?](https://crypto.stackexchange.com/questions/1812/how-to-better-generate-large-primes-sieving-and-then-random-picking-or-random-p)
2. [擴展歐幾里得算法](https://zh.wikipedia.org/wiki/%E6%89%A9%E5%B1%95%E6%AC%A7%E5%87%A0%E9%87%8C%E5%BE%97%E7%AE%97%E6%B3%95)
3. [公開金鑰加密 wiki](https://zh.wikipedia.org/wiki/%E5%85%AC%E5%BC%80%E5%AF%86%E9%92%A5%E5%8A%A0%E5%AF%86)
4. [RSA 的原理與實現](https://cjting.me/2020/03/13/rsa/)
5. [看完眼眶濕濕的App開發者慘烈對抗險惡資安環境血與淚的控訴！](https://ithelp.ithome.com.tw/users/20117445/ironman/3778?page=2)
6. [對稱密鑰加密](https://zh.wikipedia.org/wiki/%E5%B0%8D%E7%A8%B1%E5%AF%86%E9%91%B0%E5%8A%A0%E5%AF%86)
7. [Can the encryption exponent e be greater than ϕ(N)?](https://crypto.stackexchange.com/questions/5729/can-the-encryption-exponent-e-be-greater-than-%CF%95n)
8. [RSA 生成公私钥时质数是怎么选的？
](https://www.zhihu.com/question/54779059)
9. [How can I generate large prime numbers for RSA?](https://crypto.stackexchange.com/questions/71/how-can-i-generate-large-prime-numbers-for-rsa)
10. [Coppersmith D. Finding a small root of a bivariate integer equation; factoring with high bits known. EUROCRYPT 1996. pp. 178-189, ACM, 1996.](https://link.springer.com/content/pdf/10.1007%2F3-540-68339-9_16.pdf)


