package main

import (
	"fmt"
	"strconv"
)

func main(){
	//连续输入三人的姓名、性别、身高、体重、年龄信息
	var value1 [5]string
	fmt.Println("请输入第一个人的姓名：")
	fmt.Scanln(&value1[0])
	fmt.Println("请输入第一个人的性别：")
	fmt.Scanln(&value1[1])
	fmt.Println("请输入第一个人的身高：")
	fmt.Scanln(&value1[2])
	fmt.Println("请输入第一个人的体重：")
	fmt.Scanln(&value1[3])
	fmt.Println("请输入第一个人的年龄：")
	fmt.Scanln(&value1[4])
	var value2 [5]string
	fmt.Println("请输入第二个人的姓名：")
	fmt.Scanln(&value2[0])
	fmt.Println("请输入第二个人的性别：")
	fmt.Scanln(&value2[1])
	fmt.Println("请输入第二个人的身高：")
	fmt.Scanln(&value2[2])
	fmt.Println("请输入第二个人的体重：")
	fmt.Scanln(&value2[3])
	fmt.Println("请输入第二个人的年龄：")
	fmt.Scanln(&value2[4])
	var value3 [5]string
	fmt.Println("请输入第三个人的姓名：")
	fmt.Scanln(&value3[0])
	fmt.Println("请输入第三个人的性别：")
	fmt.Scanln(&value3[1])
	fmt.Println("请输入第三个人的身高：")
	fmt.Scanln(&value3[2])
	fmt.Println("请输入第三个人的体重：")
	fmt.Scanln(&value3[3])
	fmt.Println("请输入第三个人的年龄：")
	fmt.Scanln(&value3[4])


	fmt.Printf("第一个人的体脂评价为：%v/n",fat(value1))
	fmt.Printf("第二个人的体脂评价为：%v/n",fat(value2))
	fmt.Printf("第三个人的体脂评价为：%v/n",fat(value3))




}
func fat(value [5]string) [3]string{
	//var tall float64
	//var age int
	var sexWeight int
	var sex string
	//var weight int
	var name string
	name = value[0]
	sex = value[1]
	tall, terr:= strconv.ParseFloat(value[2], 64)
	weight, werr := strconv.ParseInt(value[3],10,64)
	age, aerr := strconv.ParseInt(value[4],10,64)

	if sex == "男"{
		sexWeight = 1
	}else{
		sexWeight = 0
	}
	var bmi float64 = float64(weight)/(tall* tall)
	strbmi := strconv.FormatFloat(bmi, 'f', 1, 64)
	result := [3]string{}
	if terr !=nil || werr!=nil || aerr!=nil {
		result[0] = name
		result[1] = strbmi
		result[2] = "数据不全"
		return result
	}else{
		var fatRate float64 = (1.2 * bmi + 0.23 * float64(age) - 5.4 -10.8 *float64(sexWeight))/100
		var fcha int
		if sex =="男"{
			if age >= 18 && age<= 39{
				if fatRate <= 0.1{
					fcha = 0
				}else if fatRate > 0.1 && fatRate <= 0.16{
					fcha = 1
				}else if fatRate > 0.16 && fatRate <= 0.21{
					fcha = 2
				}else if fatRate > 0.21 && fatRate <= 0.26{
					fcha = 3
				}else{
					fcha = 4
				}
			}else if age >= 40 && age <= 59{
				if fatRate <= 0.11{
					fcha = 0
				}else if fatRate > 0.11 && fatRate <= 0.17{
					fcha = 1
				}else if fatRate > 0.17 && fatRate <= 0.22{
					fcha = 2
				}else if fatRate > 0.22 && fatRate <= 0.27{
					fcha =3
				}else{
					fcha = 4
				}
			}else{
				if fatRate <= 0.13{
					fcha = 0
				}else if fatRate > 0.13 && fatRate <= 0.19{
					fcha = 1
				}else if fatRate > 0.19 && fatRate <= 0.24{
					fcha = 2
				}else if fatRate > 0.24 && fatRate <= 0.29{
					fcha =3
				}else{
					fcha = 4
				}


			}
		}else{
			if age >= 18 && age<= 39{
				if fatRate <= 0.2{
					fcha = 0
				}else if fatRate > 0.2 && fatRate <= 0.27{
					fcha = 1
				}else if fatRate > 0.27 && fatRate <= 0.34{
					fcha = 2
				}else if fatRate > 0.34 && fatRate <= 0.39{
					fcha =3
				}else{
					fcha = 4
				}
			}else if age >= 40 && age <= 59{
				if fatRate <= 0.21{
					fcha = 0
				}else if fatRate > 0.21 && fatRate <= 0.28{
					fcha = 1
				}else if fatRate > 0.28 && fatRate <= 0.35{
					fcha = 2
				}else if fatRate > 0.35 && fatRate <= 0.4{
					fcha =3
				}else{
					fcha = 4
				}
			}else{
				if fatRate <= 0.22{
					fcha = 0
				}else if fatRate > 0.22 && fatRate <= 0.29{
					fcha = 1
				}else if fatRate > 0.29 && fatRate <= 0.36{
					fcha = 2
				}else if fatRate > 0.36 && fatRate <= 0.41{
					fcha =3
				}else{
					fcha = 4
				}


			}

		}
		var suggestion  string
		if fcha == 0{
			suggestion = "偏瘦，请调整饮食"
		}else if fcha == 1{
			suggestion = "标准请继续保持"
		}else if fcha == 2{
			suggestion = "偏重，请控制饮食"
		}else if fcha == 3{
			suggestion = "肥胖，请加强锻炼"
		}else {
			suggestion = "严重肥胖，请引起注意"
		}

		result[0] = name
		result[1] = strbmi
		result[2] = suggestion
		return result

	}


}

//	var tall float64
//	var age int
//	var sexWeight int
//	var sex string
//	if sex == "男"{
//		sexWeight = 1
//	}else {
//		sexWeight = 0
//	}
//	var bmi float64 = weight / (tall * tall)
//	var fatRate float64 = 1.2 * bmi + 0.23 * float64(age) - 5.4 -10.8 *float64(sexWeight)
//	fmt.Println("体脂率是：",fatRate)
//
//	if sex == "男"{
//		if age >= 18 && age<= 39{
//
//		} else if age >= 40 && age <= 59{
//
//		}
//	}
//
//}
