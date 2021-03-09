package abstractfactory

import "fmt"

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDeatailDAO interface {
	SaveOrderDetail()
}

type DaoFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDeatailDAO
}

type RMBMainDAO struct {
}

func (d *RMBMainDAO) SaveOrderMain() {
	fmt.Printf("order main save!\n")
}

type RMBDetailDA0 struct {
}

func (d *RMBDetailDA0) SaveDetailDAO() {
	fmt.Printf("order detail save.!")
}
