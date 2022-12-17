# wallet
Generate a personalized crypto wallet.

![avatar](https://cdn.sputniknews.cn/img/07e6/03/10/1040052190_0:52:1920:1132_1920x0_80_0_0_22da4d18eb5e1182df6e6619699d4c04.jpg)

```bash
git clone https://github.com/wheelplan/wallet.git
cd wallet/
docker build -t go-wallet .
docker run -d --cpus 1.6 --network none --restart always --name hdwallet go-wallet
```