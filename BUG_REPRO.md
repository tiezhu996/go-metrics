# BUG_REPRO

## Bug 是什么
BuildBuckets / OrderIDs / ListBuckets 返回共享底层数组子切片，worker 又用 bucket[:len(bucket)-1] 切掉最后一条，导致样本漏掉、列表串改。

## 如何触发
`go test ./...`

## 错误信息
- TestBuildBucketsFresh / TestOrderIDsFresh 失败。
- TestRunSummary 失败（Count 计数不对）。
