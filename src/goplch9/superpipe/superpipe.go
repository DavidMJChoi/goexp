package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Pipeline阶段函数类型
type Stage func(in <-chan int) <-chan int

// 恒等阶段 - 原样传递值
func identityStage(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v
		}
		close(out)
	}()
	return out
}

// 创建具有n个阶段的pipeline
func createPipeline(n int) (chan<- int, <-chan int) {
	// 创建输入通道
	input := make(chan int)

	// 构建pipeline阶段
	var current <-chan int = input
	for i := 0; i < n; i++ {
		current = identityStage(current)
	}

	// 返回输入和输出通道
	return input, current
}

// 测量pipeline性能
func measurePipeline(n int, value int) time.Duration {
	input, output := createPipeline(n)

	// 开始计时
	start := time.Now()

	// 发送值
	input <- value

	// 接收值
	result := <-output

	// 计算耗时
	duration := time.Since(start)

	if result != value {
		log.Fatalf("数值错误：期望 %d, 得到 %d", value, result)
	}

	close(input)
	return duration
}

// 通过创建pipeline直到失败来测试内存限制
func findMaxPipelineStages() int {
	const increment = 1000
	var stages int

	for stages = increment; ; stages += increment {
		log.Printf("测试 %d 个阶段...", stages)

		// 创建前检查内存
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		beforeHeap := m.HeapAlloc

		// 尝试创建pipeline
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("在 %d 阶段发生panic: %v", stages, r)
				}
			}()

			input, output := createPipeline(stages)

			// 用单个值测试
			go func() {
				input <- 42
				close(input)
			}()

			<-output

			// 创建后检查内存
			runtime.ReadMemStats(&m)
			afterHeap := m.HeapAlloc
			log.Printf("内存使用: %.2f MB", float64(afterHeap-beforeHeap)/(1024*1024))
		}()

		// 强制垃圾回收
		runtime.GC()
		time.Sleep(100 * time.Millisecond)
	}
}

// 测试不同阶段数量的实验
func runExperiments(maxStages int) {
	// 测试值
	testValues := []int{1, 42, 100}

	fmt.Println("\n=== Pipeline性能分析 ===")
	fmt.Printf("CPU核心数: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// 测试递增的阶段数量
	for stages := 1; stages <= maxStages; stages *= 2 {
		fmt.Printf("\n--- %d 个pipeline阶段 ---\n", stages)

		var totalDuration time.Duration
		var measurements []time.Duration

		// 多次测量
		for i := 0; i < 5; i++ {
			for _, val := range testValues {
				duration := measurePipeline(stages, val)
				measurements = append(measurements, duration)
				totalDuration += duration
				time.Sleep(10 * time.Millisecond) // 让goroutine清理
			}
		}

		// 计算统计
		avgDuration := totalDuration / time.Duration(len(measurements))
		minDuration := measurements[0]
		maxDuration := measurements[0]
		for _, d := range measurements[1:] {
			if d < minDuration {
				minDuration = d
			}
			if d > maxDuration {
				maxDuration = d
			}
		}

		fmt.Printf("平均传输时间: %v\n", avgDuration)
		fmt.Printf("最小传输时间: %v\n", minDuration)
		fmt.Printf("最大传输时间: %v\n", maxDuration)
		fmt.Printf("每阶段开销: %v\n", avgDuration/time.Duration(stages))

		// 计算理论最大吞吐量
		theoreticalMax := time.Second / avgDuration
		fmt.Printf("理论最大吞吐量: %.0f 值/秒\n", float64(theoreticalMax))
	}
}

// 带缓冲的优化版本
func createBufferedPipeline(n int, bufferSize int) (chan<- int, <-chan int) {
	input := make(chan int)
	var current <-chan int = input

	for i := 0; i < n; i++ {
		// 使用带缓冲的通道提高吞吐量
		out := make(chan int, bufferSize)
		go func(in <-chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(current, out)
		current = out
	}

	return input, current
}

// 测量缓冲大小的影响
func benchmarkBufferImpact() {
	fmt.Println("\n=== 缓冲大小对性能的影响 ===")

	for stages := 1; stages <= 10000; stages *= 10 {
		for bufferSize := 1; bufferSize <= 100; bufferSize *= 10 {
			// 创建pipeline并测量
			input, output := createBufferedPipeline(stages, bufferSize)

			start := time.Now()
			go func() { input <- 42 }()
			<-output
			duration := time.Since(start)
			close(input)

			fmt.Printf("阶段数: %d, 缓冲大小: %d, 传输时间: %v\n",
				stages, bufferSize, duration)
		}
	}
}

func main() {
	// 命令行参数
	maxStages := flag.Int("max", 10000, "要测试的最大pipeline阶段数")
	memoryTest := flag.Bool("memory", false, "运行内存限制测试")
	bufferTest := flag.Bool("buffer", false, "运行缓冲测试")
	flag.Parse()

	if *memoryTest {
		fmt.Println("警告：内存测试将尝试找到系统限制")
		fmt.Println("这可能导致系统无响应！")
		fmt.Print("继续？(y/n): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			os.Exit(0)
		}

		findMaxPipelineStages()
		return
	}

	if *bufferTest {
		benchmarkBufferImpact()
		return
	}

	runExperiments(*maxStages)
}
