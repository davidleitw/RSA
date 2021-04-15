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

先討論只選擇一個質數的情況，假設我們今天要一個 $x$ 位元的質數，我們可以隨機挑一個 $x$ 位元的奇數，然後使用質數判斷法來確認隨機選取的數是不是質數。如果不是質數則在重新選取一次。

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

## Reference
- [How to better generate large primes: sieving and then random picking or random picking and then checking?](https://crypto.stackexchange.com/questions/1812/how-to-better-generate-large-primes-sieving-and-then-random-picking-or-random-p)
- [擴展歐幾里得算法](https://zh.wikipedia.org/wiki/%E6%89%A9%E5%B1%95%E6%AC%A7%E5%87%A0%E9%87%8C%E5%BE%97%E7%AE%97%E6%B3%95)
- [公開金鑰加密 wiki](https://zh.wikipedia.org/wiki/%E5%85%AC%E5%BC%80%E5%AF%86%E9%92%A5%E5%8A%A0%E5%AF%86)
- [RSA 的原理與實現](https://cjting.me/2020/03/13/rsa/)
- [看完眼眶濕濕的App開發者慘烈對抗險惡資安環境血與淚的控訴！](https://ithelp.ithome.com.tw/users/20117445/ironman/3778?page=2)
- [對稱密鑰加密](https://zh.wikipedia.org/wiki/%E5%B0%8D%E7%A8%B1%E5%AF%86%E9%91%B0%E5%8A%A0%E5%AF%86)
- [Can the encryption exponent e be greater than ϕ(N)?](https://crypto.stackexchange.com/questions/5729/can-the-encryption-exponent-e-be-greater-than-%CF%95n)
