package main

import "fmt"

func main(){
	var p1x, p2x, p3x, p4x,p1y, p2y,p3y,p4y float64
	fmt.Println("请输入第一条线的两个点：")
	p1x, p1y = getxy()
	p2x, p2y = getxy()
	fmt.Println("请输入第二条线的两个点：")
	p3x, p3y = getxy()
	p4x, p4y = getxy()
	k1 := calculateK(p2y, p1y, p2x, p1x)
	k2 := calculateK(p4y, p3y, p4x, p3x)
	getResult(k1, k2)


}
func getxy()(float64,float64){
	var x, y float64
	fmt.Print("录入x：")
	fmt.Scanln(&x)
	fmt.Print("录入y：")
	fmt.Scanln(&y)
	return x, y
}
func calculateK(y2, y1, x2, x1 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}

func getResult(k1, k2 float64) {
	if k1 == k2 {
		fmt.Println("两条线平行")
	} else {
		fmt.Println("两条线不平行")
	}
}