## API Docs

### 登录接口

* 作用：用户登录验证，如果验证通过，则设置相应的Cookie
* 路径：/auth/login
* 方法：POST
* 参数

| Parameter | Description |
|-----------|-------------|
| userid    | 用户 id       |
| password  | 用户密码        |

* 参数格式：application/x-www-form-urlencoded
* 返回值示例（JSON格式）

```json
{
  "code": "200",
  "msg": ""
}
```

* 返回值

| Response | Description |
|----------|-------------|
| code     | 返回值，200表示成功 |
| msg      | 失败的原因       |

### 翻译接口

* 作用：翻译功能调用该接口进行翻译
* 路径：/api/translate
* 方法：POST
* 参数

| Parameter | Description                      |
|-----------|----------------------------------|
| source    | 需要翻译的文本                          |
| target    | 目标语言 "zh"（中文）或"en"（英文）或"yue"（粤语） |

* 参数格式：application/x-www-form-urlencoded
* 返回值示例（JSON格式）

```json
{
  "dst": "你怎么看？"
}
```

* 返回值

| Response | Description |
|----------|-------------|
| dst      | 翻译后的文本      |

### 即时通讯模块令牌刷新

* 作用：用户端调用腾讯云IM需要有一个服务器端签发的具有有效期的令牌，本接口起到检查令牌有效期的作用，如果过期/无效，则发放新的令牌
* 路径：/chat/refresh
* 方法：POST
* 参数

| Parameter | Description |
|-----------|-------------|
| sig       | 当前存储在用户端的令牌 |

* 参数格式：application/x-www-form-urlencoded
* 返回值示例（JSON格式）

```json
{
  "ok": false,
  "sig": "just an example"
}
```

* 返回值

| Response | Description                        |
|----------|------------------------------------|
| ok       | true：原有的sig有效，false：原有的sig无效，更新sig |
| sig      | 当ok为false时，使用该sig来调用通讯模块           |

### 文件下载接口

* 文件模块中文件下载调用该接口
* 路径：/files/download?filename=${filename}&group=${group}
* 方法：GET
* 参数

| Parameter | Description                           |
|-----------|---------------------------------------|
| filename  | 需要下载的文件名                              |
| group     | 用户所属的小组（一个用户可能属于多个小组），该字段制定了下载哪个小组的文件 |

* 参数格式：application/x-www-form-urlencoded
* 返回： 相应的文件

### 文件删除接口

* 路径：/files/delete
* 方法：POST
* 参数

| Parameter | Description  |
|-----------|--------------|
| filename  | 需要删除的文件名     |
| group     | 需要删除文件所属的小组号 |

* 参数格式：application/x-www-form-urlencoded
* 返回值示例（JSON格式）

```json
{
  "ok": true,
  "reason": ""
}
```

* 返回值

| Response | Description          |
|----------|----------------------|
| ok       | true：删除成功，false：删除失败 |
| reason   | 删除失败的原因              |

### 创建新文档接口

* 路径：/docs/new
* 方法：POST
* 参数

| Parameter | Description   |
|-----------|---------------|
| markdown  | markdown格式的文本 |
| group     | 新文档所属的小组号     |
| title     | 文档的标题         |

* 参数格式示例(JSON):

```json
{
  "markdown": "# hello world",
  "group": "1",
  "title": "hello"
}
```

* 返回值示例（JSON格式）

```json
{
  "id": "119",
  "ok": true,
  "markdown": "",
  "title": ""
}
```

* 返回值

| Response | Description |
|----------|-------------|
| id       | 新文档被分配的id   |
| ok       | 文档是否被成功创建   |
| 其他字段     | 本接口暂未用到     |

### 获取对应id文档内容接口

* 路径：/docs/byid?group=1&id=114
* 方法：GET
* 参数

| Parameter | Description           |
|-----------|-----------------------|
| id        | 需要获取文档的id             |
| group     | 用户所属的小组（一个用户可能属于多个小组） |

* 返回示例（JSON）

```json
{
  "id": "",
  "ok": true,
  "markdown": "# hello world",
  "title": "hello"
}
```

* 返回

| Response | Description     |
|----------|-----------------|
| id       | 未用到             |
| ok       | 是否成功获取相应文档      |
| markdown | Markdown格式的本地文档 |
| title    | 文档的标题           |
