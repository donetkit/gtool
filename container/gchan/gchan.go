package gchan

import "time"

func BatchProcessorWithTimeout[T any](input <-chan T, batchSize int, timeout time.Duration) <-chan []T {
	output := make(chan []T)

	go func() {
		defer close(output)
		batch := make([]T, 0, batchSize)
		timer := time.NewTimer(timeout)
		timer.Stop() // 初始时停止定时器
		defer timer.Stop()

		sendBatch := func() {
			if len(batch) > 0 {
				output <- batch
				batch = make([]T, 0, batchSize) // 保持容量避免重复分配
			}
			timer.Stop()
			// 确保排空可能残留的定时信号
			select {
			case <-timer.C:
			default:
			}
		}

		for {
			select {
			case v, ok := <-input:
				if !ok {
					sendBatch()
					return
				}

				batch = append(batch, v)

				// 达到批次大小时立即发送
				if len(batch) >= batchSize {
					sendBatch()
				} else {
					// 重置定时器（先停后启）
					if !timer.Stop() {
						select {
						case <-timer.C:
						default:
						}
					}
					timer.Reset(timeout)
				}

			case <-timer.C:
				sendBatch()
			}
		}
	}()

	return output
}

func BatchProcessor[T any](input <-chan T, batchSize int) <-chan []T {
	output := make(chan []T)

	go func() {
		defer close(output)

		// 使用带预分配的切片
		batch := make([]T, 0, batchSize)

		for item := range input {
			batch = append(batch, item)

			// 直接比较长度更直观
			if len(batch) == batchSize {
				// 发送副本避免数据竞争
				output <- append([]T{}, batch...)
				batch = make([]T, 0, batchSize) // 重置时保持容量
			}
		}

		// 处理剩余元素
		if len(batch) > 0 {
			output <- batch
		}
	}()

	return output
}
