# 知道

---
CSA golang后端2023年初期考核


## API参考

---
#### 注册
```
  POST /user/register
```

| 参数         | 位置   | 类型     | 描述               |
|:-----------|:-----|:------- |:-----------------|
| `username` | body | `string` | **必选**, 用户名，不能重复 |
| `password` | body | `string` | **必选**, 用户密码     |

返回参数：

| 参数         | 类型     | 描述   |
|:-----------| :------- |:-----|
| `status` | `int` | 状态码  |
| `info` | `string` | 返回信息 |

成功返回示例：
```json
  {
    "status": 200,
    "info": "success"
  }
```
---
#### 登录

```
  GET /user/login
```

| 参数         | 位置   | 类型     | 描述           |
|:-----------|:-----|:------- |:-------------|
| `username` | body | `string` | **必选**, 用户名  |
| `password` | body | `string` | **必选**, 用户密码 |

返回参数：

| 参数              | 类型       | 描述             |
|:----------------|:---------|:---------------|
| `status`        | `int`    | 状态码            |
| `token`         | `string` | token，有效期12h   |
| `refresh_token` | `string` | 刷新token，有效期24h |

成功返回示例：
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDE0OjU0OjE4LjEwNDE2NzYyMiArMDgwMCBDU1QgbT0rOTAzOTYuOTA3MDQ2OTc2IiwiaWQiOiIyIn0.UCzKCkrhnVOCY3eunSJFIHdjio3ZoB1sCkZLb8t3kbM",
  "status": 200,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDAyOjU0OjE4LjEwNDEzMTg2MyArMDgwMCBDU1QgbT0rNDcxOTYuOTA3MDExMjA3IiwiaWQiOiIyIn0.YPhnKSoEi33lezbaIZyBZjks44LDC9abOqcelDp_QHE"
}
```
---
#### 刷新token

```
  GET /user/login/refresh
```

| 参数              | 位置    | 类型     | 描述              |
|:----------------|:------|:------- |:----------------|
| `refresh_token` | Query | `string` | **必选**, 刷新token |

返回参数：

| 参数              | 类型       | 描述      |
|:----------------|:---------|:--------|
| `status`        | `int`    | 状态码     |
| `token`         | `string` | token   |
| `refresh_token` | `string` | 刷新token |

成功返回示例：
```json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDE1OjA5OjEzLjcxMDU2ODgyMiArMDgwMCBDU1QgbT0rOTEyOTIuNTEzNDQ4MTc1IiwiaWQiOiIyIn0.qEfcN8RXIPik_hS-AR1mr1N-zJywysmXRsQnXCU2BMU",
    "status": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDAzOjA5OjEzLjcxMDU0NDA3MSArMDgwMCBDU1QgbT0rNDgwOTIuNTEzNDIzNDE1IiwiaWQiOiIyIn0.RXTCSbIJ1vGlhG7RTGrscf7TKSfilObsDqaS75TTV_U"
}
```
---
#### 创建提问
```
  POST /question/create
```
| 参数              | 位置     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |
| `title`         | body   | `string` | **必选**, 问题标题  |
| `introduce`     | body   | `string` | 问题详细描述        |
| `topic`         | body   | `string` | 问题分类          |

返回参数：

| 参数            | 类型    | 描述   |
|:--------------|:------|:-----|
| `status`      | `int` | 状态码  |
| `question_id` | `int` | 问题id |

成功返回示例
```json
{
    "question_id": 2,
    "status": 200
}
```
---
#### 创建回答
```
  POST /question/answer
```
| 参数              | 位置     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |
| `question_id`   | query  | `int`    | **必选**, 问题id  |
| `text`          | body   | `string` | **必选**, 回答文本  |

返回参数：

| 参数          | 类型       | 描述   |
|:------------|:---------|:-----|
| `status`    | `int`    | 状态码  |
| `answer_id` | `string` | 回答id |

成功返回示例
```json
{
  "answer_id": 2,
  "status": 200
}
```
---
#### 获取用户创建的问题与回答
```
  GET /question/my
```
| 参数              | 位置     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |

返回参数：

| 参数              | 类型             | 描述   |
|:----------------|:---------------|:-----|
| `status`        | `int`          | 状态码  |
| `question_list` | `array struct` | 回答id |
| `answer_list`   | `array struct` | 回答id |

成功返回示例
```json
{
  "answer_list": [
    {
      "id": 2,
      "question_id": 2,
      "answerer_id": 2,
      "text": "ypm1",
      "like": 0
    }
  ],
  "question_list": [
    {
      "id": 2,
      "questioner_id": 2,
      "title": "hello world",
      "introduce": "hello1",
      "topic": "hello"
    }
  ],
  "status": 200
}
```
---
#### 修改问题
```
  PUT /question/change_question/{question_id}
```
| 参数              | 位置     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |
| `question_id`   | path   | `int`    | **必选**, 问题id  |
| `title`         | body   | `string` | 问题标题          |
| `introduce`     | body   | `string` | 问题详细描述        |
| `topic`         | body   | `string` | 问题分类          |

返回参数：

| 参数       | 类型       | 描述  |
|:---------|:---------|:----|
| `status` | `int`    | 状态码 |
| `info`   | `string` | 信息  |

成功返回示例
```json
{
  "info": "success",
  "status": 200
}
```
---
#### 修改回答
```
  PUT /question/change_answer/{answer_id}
```
| 参数              | 位置     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |
| `answer_id`     | path   | `int`    | **必选**, 回答id  |
| `text`          | body   | `string` | 回答内容          |

返回参数：

| 参数       | 类型       | 描述  |
|:---------|:---------|:----|
| `status` | `int`    | 状态码 |
| `info`   | `string` | 信息  |

成功返回示例
```json
{
  "info": "success",
  "status": 200
}
```
---
#### 删除问题
从产品的角度讲，只有在用户创建的问题没有回答的情况下，用户才能够删除问题。所以删除问题时要求该问题下没有回答。
```
  DELETE /question/delete_question/{question_id}
```
| 参数              | 位置     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |
| `question_id`   | path   | `int`    | **必选**, 问题id  |

返回参数：

| 参数       | 类型       | 描述  |
|:---------|:---------|:----|
| `status` | `int`    | 状态码 |
| `info`   | `string` | 信息  |

成功返回示例
```json
{
  "info": "success",
  "status": 200
}
```
---
#### 删除回答
```
  DELETE /question/delete_answer/{answer_id}
```
| 参数              | 位置     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |
| `answer_id`     | path   | `int`    | **必选**, 回答id  |

返回参数：

| 参数       | 类型       | 描述  |
|:---------|:---------|:----|
| `status` | `int`    | 状态码 |
| `info`   | `string` | 信息  |

成功返回示例
```json
{
  "info": "success",
  "status": 200
}
```