# ChatGPT-CLI

基于OpenAI GPT-3.5 模型API封装的控制台应用，可以根据token数量设置实现相同于官网的聊天功能。  
不会在Token达到上限后中断聊天。  
可以在本地文件中保存、回复、共享会话，让你免受OpenAI网站繁忙导致的历史会话无法打开故障。  
A console application based on OpenAI GPT-3.5 model API, which can achieve the same chat function as the official website.  
The chat will not be interrupted or reset after the token limit is reached.

支持以下功能

- [x]  账户切换
- [x]  对话本地保存
- [x]  对话恢复
- [x]  对话内容编辑
- [x]  对话分享


The following functions are supported:
- [x]  Account switching
- [x]  Local saving of conversations
- [x]  Conversation recovery
- [x]  Editing of conversation content
- [x]  Sharing of conversations.

## 使用方法 Usage
根据你的操作系统，从Release中下载最新版本的程序，从命令行/终端运行即可。  
默认使用中文，如果需要使用英文，可以在命令行中设置环境变量`LANG=en`。  

According to your operating system, download the latest version of the program from the Release section, and run it from the command line/terminal.  
By default, it uses Chinese. If you need to use English, you can set the environment variable LANG=en in the command line.
### windows
```bash
# 中文
.\chatgpt-cli.exe
# english
set LANG=en
.\chatgpt-cli.exe
```
### linux/macos
```bash
# 中文
./chatgpt-cli
# english
LANG=en ./chatgpt-cli
```

## 配置说明 Configuration

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

- `Host` The API domain name. By default, it uses the OpenAI official website address. If you want to use other accelerated domains, you can set it here.  
- `Proxy` The proxy server address. If you need to use a proxy, you can configure it here.  
- `DefaultConversionConfig` The default configuration used for each session. Currently, separate configuration for each individual session is not supported, so you have to manually modify the settings in the session directory if needed.  
- `Model` The API model used for the session. Currently, the Chat mode is only limited to gpt-3.5-turbo.
- `Temperature` One of the OpenAI Chat parameters that controls the model's randomness. Please refer to the OpenAI documentation for more information.
- `TopP` One of the OpenAI Chat parameters that controls the model's randomness. Please refer to the OpenAI documentation for more information.
- `MaxTokens` **Replaces the OpenAI parameter. It is the maximum number of tokens allowed per round of conversation. Historical sessions that exceed the token limit will be removed from the request. The larger the MaxTokens value, the stronger the AI's identity and the more expensive the price.**
- `N` the number of replies returned by the AI at a time. Returning two replies at once will save more tokens than making two separate requests for the same amount of replies.

## 背景 Background

最近ChatGPT的页面不太稳定，历史会话总是打不开，花了很多时间磨合(划掉)的角色都无法找回，所以这些会话还是保存在本地比较好。

而且由于API可以编辑AI返回的会话，通过修改会话的方式，可以更快让AI进入设定的角色（划掉）。
