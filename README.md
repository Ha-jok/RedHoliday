## 接口文档

## RedHoliday


### front-page


#### `/redholiday/front-page` `GET`

- 网站首页


|请求参数|说明|
|---|---|
|token|用户token|


| 返回参数 | 说明 |
|--- | ---|
|status|状态码|
|message|提示信息|
|uid|用户id|


|status|message|说明|
|---|---|---|
|`true`|`"欢迎回来，uid"`|用户已登录|
|`false`|`"欢迎访问，游客"`|用户未登录|


### user


#### `/redholiday/user/regist` `POST`

- `application/x-www-form-urlencoded`
- 注册新用户；

|请求参数|类型|说明|
|---|---|---|
|Username|必选|用户名|
|Password|必选|用户密码|
|Email|可选|绑定邮箱|
|Phone|可选|绑定手机号|


|返回参数|说明|
|---|---|
|status|状态码|
|message|提示信息|


|status|message|说明|
|---|---|---|
|`false`|`"用户名格式错误"`|`username`大于6位或为空|
|`false`|`"密码格式错误"`|`password`小于六位或大于十二位或为空|
|`false`|`"手机号格式错误"`|`phone`不是十一位|
|`false`|`"邮箱格式错误"`|邮箱不满足格式|
|`true`|`"注册成功",username`|参数合法|


#### `/redholiday/user/login/pw` `POST`

- `application/x-www-form-urlencoded`
- 密码登录

| 请求参数  | 类型 | 说明       |
| --------- | ---- | ---------- |
| Username | 必选 | 用户名 |
| Password | 必选 | 密码|


| 返回参数 | 说明 |
|---|---|
|status|状态码|
|message| 提示信息|
|token|用户token|


|status|message|说明|
|---|---|---|
|`false`|`"用户名不能为空"`|`Username`为空|
|`false`|`"密码不能为空"`|`password`为空|
| `false` | `"用户名或密码错误"`|`Username`和`password`不匹配或`Username`不存在|
| `true`|`"欢迎回来"，username`|`Username`和`password`匹配|





#### `/redholiday/user/login/email/verify` ``POST`

- `application/x-www-form-urlencoded`
- 获取邮箱验证码

|请求参数|类型|说明|
|---|---|---|
|email|必选|用户邮箱|


|返回参数|说明|
|---|---|
|status|状态码|
|message|提示信息|

|status|message|说明|
|---|---|---|
|`false`|`"邮箱格式错误"`|`email`不合法|
| `false` | `"用户不存在“`|`email`无对应的用户|
|`true`|` `|`email`合法|

#### `/redholiday/user/login/email` ``POST`

- `application/x-www-form-urlencoded`
- 邮箱验证码登录，先调用`/redholiday/user/login/email/verify`接口获取验证码，然后通过验证码验证登录

|请求参数|类型|说明|
|---|---|---|
|verify_code|必选|邮箱验证码|


|返回参数|说明|
|---|---|
|status|状态码|
|message|提示信息|
|token|用户token|
|username|用户名|

|status|message|说明|
|---|---|---|
|`false`|`"验证码不能为空"`|`verify_code`为空|
| `false`| `"验证码错误“`|`verify_code`和用户邮箱接收到的验证码不一样|
|`true`|`"欢迎回来，"username`|`verify_code`和用户邮箱接收到的验证码一样|







#### `/redholiday/user/personal-information/:username` `GET`

- 获取Username为`:username`的用户的个人信息

|请求参数|类型|说明|
|---|---|---|
|username|必选|用户名|


|返回参数|说明|
|---|---|
|status|状态码|
|data|返回信息|
|message|提示信息|

|status|data|message|说明|
|---|---|---|---|
|`false`| ` ` | `"username无效"`|`username`为空或无效|
|`ture`|参考下列代码|` `|用户个人信息|
```go
{
	UID: Number,
	Username: String,
	Balance: Number, //余额
	Friends: Array, //[]Friend; 好友切片
	Follow-business: Array, //[]Follow-business;关注商家切片
}
```


#### `/redholiday/user/shopping-cart` `GET`

- 查看用户购物车

|请求参数|类型|说明|
|---|---|---|
|token|用户token|


|返回参数|说明|
|---|---|
|status|状态码|
|data|返回信息|
|message|提示信息|


|status|data|message|说明|
|---|---|---|---|
|`false`|` `|`"token无效"`|用户`token`不存在或已失效|
|`ture`|参考下列代码|` `|用户`token`有效|
```go
{
	Shopping-cart: Array //[]Shopping-cart;购物车切片
}
```


#### `/redholiday/user/shopping-cart` `POST`

- 修改用户购物车

|请求参数|类型|说明|
|---|---|---|
|token|必选|用户token|
|settlement|可选|结算商品uid|
|delete|可选|删除商品uid|


|返回参数|说明|
|---|---|
|status|状态码|
|data|返回信息|
|message|提示信息|


|status|data|message|说明|
|---|---|---|---|
|`false`|`参考下列代码 `|`"token无效"`|用户`token`不存在或已失效|
|`ture`|` 参考下列代码`|"结算成功，请支付" |`add`不为空|
|`ture`|` 参考下列代码`|删除成功|`delete`不为空|
```go
{
	Shopping-cart: Array //[]Shopping-cart;购物车切片
}
```


#### `/redholiday/user/order` `GET`

- 查看用户订单

|请求参数|类型|说明|
|---|---|---|
|token|必选|用户token|


|返回参数|说明|
|---|---|
|status|状态码|
|data|返回信息|
|message|提示信息|


|status|data|message|说明|
|---|---|---|---|
|`false`|` `|`"token无效"`|用户`token`不存在或已失效|
|`ture`|参考下列代码|` `|用户`token`有效|
```go
	Order-paid: Array, //[]Order-paid;已支付订单切片
	Order-unpaid: Array, //[]Order-unpaid;未支付切片
	Pending-payment: Array, //[]Pending-payment;待支付订单
	Order-received: Array,  //[]Order-received;已收货订单
```


#### `/redholiday/user/order` `POST`

- 修改用户订单，确定收货或取消订单

|请求参数|类型|说明|
|---|---|---|
|token|用户token|
|cancel|可选|用户取消订单的商品uid|
|receit|可选|用户确定收货订单的商品uid|


|返回参数|说明|
|---|---|
|status|状态码|
|data|返回信息|
|message|提示信息|

|status|data|message|说明|
|---|---|---|---|
|`false`|`参考下列代码 `|`"token无效"`|用户`token`不存在或已失效|
|`ture`|` 参考下列代码`|"确认收货" |`receit`不为空|
|`ture`|` 参考下列代码`|"取消成功"|`delete`不为空|
```go
	Order-paid: Array, //[]Order-paid;已支付订单切片
	Order-unpaid: Array, //[]Order-unpaid;未支付切片
	Order-received: Array,  //[]Order-received;已收货订单
```


####  `/redholiday/user/avatar`  `PUT`

- `multipart/form-data`
- 修改/添加头像


|请求参数|类型|说明|
|---|---|---|
|avatar|必选|头像（二进制文件）|
|token|必选|用户token|


|返回参数|说明|
|---|---|
|status|状态码|
|message|提示信息|


|status|message|说明|
|---|---|---|
|`false`|`"请登录"`|`token`为空或失效|
|`false`|`"上传失败"`|图像上传失败|
|`true`|`"上传成功"`|图像成功上传|


### commodity

#### `/redholiday/commodity/:uid`   `GET`

- UID为`:uid`的商品的详情页

|请求参数|类型|说明|
|---|---|---|
|uid|必选|商品uid|


|返回参数|说明|
|---|---|
|status|状态码|
|data|返回信息|
|message|提示信息|


|status|data|message|说明|
|---|---|---|---|
|`false`|` `|`商品不存在或已下架`|`uid`不存在|
|`true`|参考下列代码|` `|商品信息|
```go
{
	UID: Number,
	Commidity-Name: String,
	Volume: Number,//商品成交量
	Evaluations: Array, //[]Evaluations;评论参数
	Detailed-Introduction: String, //商品详细介绍
}
```



#### `/redholiday/commidity/:uid` `Post`

- 对`uid`为`:uid`的商品进行评论及添加购物车

|请求参数|类型|说明|
|---|---|---|
|uid|必选|商品uid|
|comment|可选|评论|
|token|必选|用户token|
|add|可选|添加购物车|


|返回参数|说明|
|---|---|
|status|状态码|
|message|提示信息|


|status|message|说明|
|---|---|---|
|`false`|`"token无效"`|`token`无效或为空|
|`true`|`"评论成功"`|`comment`不为空|
|`true`|`"添加成功"`|`add`为`add`且`token`有效|



















