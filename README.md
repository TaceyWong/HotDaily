<div align="center">
<img alt="logo" height="120" src="./public/favicon.png" width="120"/>
<h2>今日热榜</h2>
<p>一个聚合热门数据的 API 接口</p>
<small> <a href="https://github.com/imsyy/DailyHotApi">DailyHotApi</a>的Golang实现</small>
</div>

## 🚩 特性

- 极快响应，便于开发
- 支持 RSS 模式和 JSON 模式
- 支持多种部署方式
- 简明的路由目录，便于新增

## 👀 示例

> 这里是使用该 API 的示例站点  
> 示例站点可能由于访问量或者长久未维护而访问异常  
> 若您也使用了本 API 搭建了网站，欢迎提交您的站点链接

- [今日热榜 - https://hot.imsyy.top/](https://hot.imsyy.top/)

## 🎉 总览

> 🟢 状态正常 / 🟠 可能失效 / ❌ 无法使用 / ⚠️ 需要科学上网

| **站点**         | **类别**     | **调用名称**   | **状态** |
| ---------------- | ------------ | -------------- | -------- |
| 哔哩哔哩         | 热门榜       | bilibili       | 🟢       |
| AcFun            | 排行榜       | acfun          | 🟢       |
| 微博             | 热搜榜       | weibo          | 🟢       |
| 知乎             | 热榜         | zhihu          | 🟢       |
| 知乎日报         | 推荐榜       | zhihu-daily    | 🟢       |
| 百度             | 热搜榜       | baidu          | 🟢       |
| 抖音             | 热点榜       | douyin         | 🟢       |
| 豆瓣电影         | 新片榜       | douban-movie   | 🟢       |
| 豆瓣讨论小组     | 讨论精选     | douban-group   | 🟢       |
| 百度贴吧         | 热议榜       | tieba          | 🟢       |
| 少数派           | 热榜         | sspai          | 🟢       |
| IT之家           | 热榜         | ithome         | 🟠       |
| IT之家「喜加一」 | 最新动态     | ithome-xijiayi | 🟠       |
| 简书             | 热门推荐     | jianshu        | 🟠       |
| 澎湃新闻         | 热榜         | thepaper       | 🟢       |
| 今日头条         | 热榜         | toutiao        | 🟢       |
| 36 氪            | 热榜         | 36kr           | 🟢       |
| 51CTO            | 推荐榜       | 51cto          | 🟢       |
| 稀土掘金         | 热榜         | juejin         | 🟢       |
| 腾讯新闻         | 热点榜       | qq-news        | 🟢       |
| 网易新闻         | 热点榜       | netease-news   | 🟢       |
| 英雄联盟         | 更新公告     | lol            | 🟢       |
| 原神             | 最新消息     | genshin        | 🟢       |
| 崩坏3            | 最新动态     | honkai         | 🟢       |
| 崩坏：星穹铁道   | 最新动态     | starrail       | 🟢       |
| 微信读书         | 飙升榜       | weread         | 🟢       |
| NGA              | 热帖         | ngabbs         | 🟢       |
| V2EX             | 主题榜       | v2ex           | ⚠️       |
| HelloGitHub      | Trending     | hellogithub    | 🟢       |
| 中央气象台       | 全国气象预警 | weatheralarm   | 🟢       |
| 中国地震台       | 地震速报     | earthquake     | 🟢       |

## ⚙️ 部署



## 👷 开发


## ⚠️ 须知

- 本项目为了避免频繁请求官方数据，默认对数据做了缓存处理，默认为 `60` 分钟，如需更改，请自行修改配置
- 本项目部分接口使用了 **页面爬虫**，若违反对应页面的相关规则，请 **及时通知我去除该接口**

## 📢 免责声明

- 本项目提供的 `API` 仅供开发者进行技术研究和开发测试使用。使用该 `API` 获取的信息仅供参考，不代表本项目对信息的准确性、可靠性、合法性、完整性作出任何承诺或保证。本项目不对任何因使用该 `API` 获取信息而导致的任何直接或间接损失负责。本项目保留随时更改 `API` 接口地址、接口协议、接口参数及其他相关内容的权利。本项目对使用者使用 `API` 的行为不承担任何直接或间接的法律责任
- 本项目并未与相关信息提供方建立任何关联或合作关系，获取的信息均来自公开渠道，如因使用该 `API` 获取信息而产生的任何法律责任，由使用者自行承担
- 本项目对使用 `API` 获取的信息进行了最大限度的筛选和整理，但不保证信息的准确性和完整性。使用 `API` 获取信息时，请务必自行核实信息的真实性和可靠性，谨慎处理相关事项
- 本项目保留对 `API` 的随时更改、停用、限制使用等措施的权利。任何因使用本 `API` 产生的损失，本项目不负担任何赔偿和责任

<!-- ## 😘 鸣谢

特此感谢为本项目提供支持与灵感的项目

- [RSSHub](https://github.com/DIYgod/RSSHub) -->

<!-- ## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=imsyy/DailyHotApi&type=Date)](https://star-history.com/#imsyy/DailyHotApi&Date) -->

https://editor.swagger.io/
https://swagger.io/docs/specification