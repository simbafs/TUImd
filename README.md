# TUImd

[bubbletea](https://github.com/charmbracelet/bubbletea) 測試

## 遇到的問題：
~~charmbracelet/bubbletea 沒有提供類似 state 的概念，子元件抓不到父元件上的屬性，會變得很複雜~~用 msg 解決

然後需要一個批次處理 Init/Update 的方法

## Save File
1. components/cmd 發出 Msg.SaveFile("")
2. components/source、components/tab 收到 Msg.SaveFile("")
	1. 因為 msg 為空，components/source 忽略
	2. 因為 msg 為空
		1. 若 m.filename 為空，compoents/tab 發出 Msg.ShowMsg("No file name")
		2. 若 m.filename 不為空，components/tab 發出 Msg.SaveFile(m.filename)
3. components/source、components/tab 收到 Msg.SaveFile(m.filename)
	1. 因為 msg 為不為空，components/source 發出 util.SaveFile(filename, body)
	2. 因為 msg 為不為空，components/tab 忽略
