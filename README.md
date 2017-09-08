# JSE Coin CPU Miner

Mines JSE coins on your CPU.


1. login ```
POST - https://jsecoin.com/server/login/
Host:jsecoin.com
Origin:https://jsecoin.com
Referer:https://jsecoin.com/platform/?
X-Requested-With:XMLHttpRequest
json body
{ email: 'user@email', password: 'dragons' }
```
1. Get a block
```
POST - https://jsecoin.com/server/request/
o=1
```
2. generate a nonce and hash the json of the block
3. if the hash starts with 4 consecutive 0s then submit the block w/ nonce + hash
```
POST - https://jsecoin.com/server/submit/
submission.block = currentBlock.block; // block id
submission.hash = hash; // hash string
submission.nonce = nonce.toString(); // nonce value
submission.uid = user.uid; // your user id
submission.siteid = 'Platform Mining';


form field data:
o=%7B%22block%22%3A20186%2C%22hash%22%3A%220000dd6d23eaaa6145974ef4572b3d2725b7da25903fb4b2407ee5605c3026c2%22%2C%22nonce%22%3A%2216919909%22%2C%22uid%22%3A763%2C%22siteid%22%3A%22Platform%20Mining%22%7D
```
4. repeat

hash `0000760d9ce054d26d5b7565c8f872dc6402189bc5ac86d99195227013f2b054`
