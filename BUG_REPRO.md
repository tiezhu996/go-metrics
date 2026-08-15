# BUG_REPRO

## Bug 是什么
MergeSummary 丢 Failed、store 重复记录静默返回 nil、service 吞掉记录错误、worker 聚合失败不累加，导致重复无报错、失败漏统计。

## 如何触发
`go test ./...`

## 错误信息
- TestMergeSummary 失败。
- TestRecordGet 失败（重复未报错）。
- TestRunFailed 失败（Failed=0 want 1）。
