# BUG_REPRO

## Bug 是什么
ValidSample 判断反向、Record/Count 去掉锁、service 去掉校验、worker 汇总去掉锁且 wg.Add 放进 goroutine，导致样本记重、统计错误并有数据竞争。

## 如何触发
`go test -race ./...`

## 错误信息
- TestValidSample / TestRecordSummary 失败（无效样本被接受）。
- `go test -race` 报 data race。
