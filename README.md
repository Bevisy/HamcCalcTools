# HamcCalcTools
Hmac 计算工具

# 功能
1.根据传入参数，返回对应的哈希值和Base64编译后的值

# 接口设计
```shell script
POST /hmac?message=cloudos&key=cloudos&algorithm=hmac-sha256
Return:
{
  "message":   message,
  "key":       key,
  "algorithm": algorithm,
  "HMAC-SHA1": {hash value},
}
```