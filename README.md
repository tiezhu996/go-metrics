# metrics

一个用 Go 写的内存指标聚合服务，演示分层、并发聚合、分桶统计与上下文取消。

## 功能
- 记录样本、按桶聚合、汇总统计
- 样本分页查询
- 并发聚合 worker 池，支持 context 取消

## 目录结构
```
cmd/metrics/        程序入口
internal/config/    环境配置
internal/model/     模型与纯工具函数
internal/store/     内存存储（样本 + 锁 + 统计）
internal/service/   业务逻辑
internal/worker/    聚合 worker 池
```

## 运行与测试
```bash
go build ./...
go test ./...
go run ./cmd/metrics
```

## 环境变量
| 变量 | 说明 | 默认值 |
|------|------|--------|
| `METRICS_WORKERS` | worker 数量 | `2` |
| `METRICS_BUCKET_SIZE` | 分桶大小 | `2` |
