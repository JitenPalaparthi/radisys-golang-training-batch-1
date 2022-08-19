package main

import "fmt"

func main() {

	fmt.Println(sumOf(10, 20, 30, 40))
	sum, _ := sumOf(10.10, 20.20, 30.30, 40.44)
	fmt.Printf("%0.2f\n", sum)

	if sum1, err := sumOf(10, 20, 30, true, "Hello"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum1)
	}
	fmt.Println(sumOf(10, 20, 30.30, 40.44))

	var num1 uint8 = 1
	var num2 uint16 = 2
	var num3 uint32 = 3
	var num4 uint64 = 4
	var num5 float32 = 5.44
	var num6 int8 = -1
	var num7 int16 = -2

	nums := make([]interface{}, 9)
	nums[0] = num1
	nums[1] = num2
	nums[2] = num3
	nums[3] = num4
	nums[4] = num5
	nums[5] = num6
	nums[6] = num7
	nums[7] = 1 //
	nums[8] = 1.1

	if sum2, err := sumOf(nums...); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sum of various types:%0.2f", sum2)
	}

}

func sumOf(vals ...any) (any, error) {
	var sum1 int = 0
	var sum2 float64 = 0
	for _, val := range vals {
		switch val.(type) {
		case int:
			sum1 = sum1 + val.(int)
		case int64:
			sum1 = sum1 + int(val.(int64))
		case int32:
			sum1 = sum1 + int(val.(int32))
		case int16:
			sum1 = sum1 + int(val.(int16))
		case int8:
			sum1 = sum1 + int(val.(int8))
		case uint64:
			sum1 = sum1 + int(val.(uint64))
		case uint32:
			sum1 = sum1 + int(val.(uint32))
		case uint16:
			sum1 = sum1 + int(val.(uint16))
		case uint8:
			sum1 = sum1 + int(val.(uint8))
		case float64:
			sum2 = sum2 + val.(float64)
		case float32:
			sum2 = sum2 + float64(val.(float32))
		default:
			return 0, fmt.Errorf("invalid type in one of the arguments")
		}
	}
	if sum1 != 0 && sum2 == 0 {
		return sum1, nil
	} else if sum2 != 0 && sum1 == 0 {
		return sum2, nil
	} else if sum1 != 0 && sum2 != 0 {
		return sum2 + float64(sum1), nil
	}

	return nil, nil
}
