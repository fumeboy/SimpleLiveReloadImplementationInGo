# SimpleLiveReloadImplementationInGo

## 需求

HTML 等资源文件修改后，在浏览器实时刷新

## 实现

HTML 加载一个 js 程序用于和服务器创建 websocket 通信。

服务器监听文件修改事件，文件修改后， 通过 websocket 通知浏览器需要刷新资源。
