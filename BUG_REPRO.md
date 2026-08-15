# BUG_REPRO

## Bug 是什么
SortSamples 排序方向反、BuildBuckets 分桶 off-by-one 漏最后一条、OrderIDs 返回内部切片、worker 丢第一条，导致顺序错、漏样本、列表串改。

## 如何触发
`go test ./...`

## 错误信息
- TestSortSamples / TestListBucketsOrder 失败。
- TestOrderIDsFresh / TestRunSummary 失败。
