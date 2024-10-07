package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义订单结构体
type Order struct {
	OrderID   string
	ProductID string
	Quantity  int
	UserID    string
}

// 定义库存管理器
// Mutex是GO中一个互斥锁，本地内存中使用的锁，用于防止多个goroutines同时访问共享资源时不会发生冲突和数据不一致
type InventoryManager struct {
	mu       sync.Mutex
	stockMap map[string]int // 商品ID -> 库存数量
}

// 检查库存是否足够
func (im *InventoryManager) CheckStock(productID string, quantity int) bool {
	im.mu.Lock()         // 锁定资源
	defer im.mu.Unlock() // defer 是 Go 的一种机制，确保无论函数是正常结束还是发生错误退出，锁都会被释放，防止死锁问题。
	stock, exists := im.stockMap[productID]
	if !exists {
		return false
	}
	return stock >= quantity
}

// 扣减库存
func (im *InventoryManager) DeductStock(productID string, quantity int) bool {
	im.mu.Lock()
	defer im.mu.Unlock()
	stock, exists := im.stockMap[productID]
	if !exists || stock < quantity {
		return false
	}
	im.stockMap[productID] -= quantity
	return true
}

// 定义支付处理器
type PaymentProcessor struct{}

// 模拟支付处理
func (pp *PaymentProcessor) ProcessPayment(order Order) bool {
	fmt.Printf("订单 %s：正在处理支付...\n", order.OrderID)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 模拟支付耗时
	fmt.Printf("订单 %s：支付成功！\n", order.OrderID)
	return true
}

// 定义物流处理器
type ShippingProcessor struct{}

// 模拟物流发货
func (sp *ShippingProcessor) ShipOrder(order Order) {
	fmt.Printf("订单 %s：正在安排物流发货...\n", order.OrderID)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 模拟物流耗时
	fmt.Printf("订单 %s：物流已发出！\n", order.OrderID)
}

// 订单处理函数
func processOrder(order Order, im *InventoryManager, pp *PaymentProcessor, sp *ShippingProcessor, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("开始处理订单 %s\n", order.OrderID)

	// 创建一个错误通道，用于在各个步骤中传递错误信息
	errChan := make(chan string, 1)

	// 检查库存
	if !im.CheckStock(order.ProductID, order.Quantity) {
		fmt.Printf("订单 %s：库存不足！\n", order.OrderID)
		return
	}

	// 并发执行扣减库存和支付处理
	var stepWg sync.WaitGroup
	stepWg.Add(2)

	// 扣减库存
	go func() {
		defer stepWg.Done()
		if !im.DeductStock(order.ProductID, order.Quantity) {
			errChan <- fmt.Sprintf("订单 %s：扣减库存失败！", order.OrderID)
		} else {
			fmt.Printf("订单 %s：库存已扣减。\n", order.OrderID)
		}
	}()

	// 支付处理
	go func() {
		defer stepWg.Done()
		success := pp.ProcessPayment(order)
		if !success {
			errChan <- fmt.Sprintf("订单 %s：支付失败！", order.OrderID)
		}
	}()

	// 等待扣减库存和支付处理完成
	stepWg.Wait()
	close(errChan) // 关闭错误通道

	// 检查是否有错误
	if errMsg, ok := <-errChan; ok {
		fmt.Println(errMsg)
		return
	}

	// 物流发货
	sp.ShipOrder(order)

	fmt.Printf("订单 %s：处理完成！\n", order.OrderID)
}

func main() {
	// 初始化库存管理器
	im := &InventoryManager{
		stockMap: map[string]int{
			"product_1": 100,
			"product_2": 200,
			"product_3": 150,
		},
	}

	// 初始化支付和物流处理器
	pp := &PaymentProcessor{}
	sp := &ShippingProcessor{}

	// 模拟接收多个订单
	orders := []Order{
		{OrderID: "order_1", ProductID: "product_1", Quantity: 2, UserID: "user_1"},
		{OrderID: "order_2", ProductID: "product_2", Quantity: 1, UserID: "user_2"},
		{OrderID: "order_3", ProductID: "product_3", Quantity: 5, UserID: "user_3"},
		{OrderID: "order_4", ProductID: "product_1", Quantity: 200, UserID: "user_4"}, // 超过库存，测试库存不足的情况
	}

	var wg sync.WaitGroup

	// 并发处理订单
	for _, order := range orders {
		wg.Add(1)
		go processOrder(order, im, pp, sp, &wg)
	}

	wg.Wait()

	fmt.Println("所有订单处理完毕！")
}
