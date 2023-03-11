# ChatGPT-CLI

基于OpenAI GPT-3.5 模型API封装的控制台应用

支持以下功能

- [x]  账户切换
- [x]  对话本地保存
- [x]  对话恢复
- [x]  对话内容编辑
- [x]  对话分享

## 配置说明

```json
// 默认配置
{
  "Host": "https://api.openai.com",
  "Proxy": "",
  "DefaultConversionConfig": {
    "Model": "gpt-3.5-turbo",
    "Temperature": 1,
    "TopP": 1,
    "MaxTokens": 2048,
    "N": 2
  }
}
```

- `Host` API域名，默认使用OpenAI官网地址，如果你要使用其它加速域名，在这里设置
- `Proxy` 代理服务器地址，如果你需要使用代理，在这里配置
- `DefaultConversionConfig` 每个会话默认使用的配置项,暂未开放对单独会话的配置，如果有需要，去会话目录下手动改一下吧
- `Model` 会话使用的API模型，目前Chat模式仅限`gpt-3.5-turbo`
- `Temperature` OpenAI Chat参数之一，控制模型随机性，参见OpenAI文档
- `TopP` OpenAI Chat参数之一，控制模型随机性，参见OpenAI文档
- `MaxTokens` **取代了OpenAI的参数。一轮会话发起时最多的Token数量，超过Token限制的历史会话会被剔出请求。越大的MaxTokens数量，AI的身份代入越强，但价格也越贵。**
- `N` AI一次返回的回复数量，一次返回两条会比相同请求发起两次省量的多。

## 背景

最近ChatGPT的页面不太稳定，历史会话总是打不开，花了很多时间磨合(划掉)的角色都无法找回，所以这些会话还是保存在本地比较好。

而且由于API可以编辑AI返回的会话，通过修改会话的方式，可以更快让AI进入设定的角色（划掉）。
