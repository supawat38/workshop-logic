package main

import (
	"fmt"
	"strconv"
)

func ExamEnternumber() {

	//รับค่าจาก input
	var number uint
	fmt.Print("Enter your number : ")
	fmt.Scanf("%d", &number)

	//ต้องกรอกตัวเลขเท่านั้น
	if number == 0 {
		fmt.Println("Error : Number only.")
		return
	}

	//(ข้อ1) ถ้าตัวเลขน้อยกว่า 6 ตัว
	numberis_string := strconv.FormatUint(uint64(number), 10)
	numberis_string_len := len(numberis_string)
	if numberis_string_len < 6 {
		fmt.Println("Error : Number must be more than 6 digits.")
		return
	}

	//สร้างเป็น array uint ก่อน
	tmp_number := make([]uint, 0)
	for _, char := range numberis_string {
		numberconverttouint, _ := strconv.ParseUint(string(char), 10, 32)
		tmp_number = append(tmp_number, uint(numberconverttouint))
	}

	// loop เช็คตัวอักษรติดกันห้าม เลขเดียวกัน
	var keepnumber uint
	tmpresult := map[uint]uint{}
	for index, number := range tmp_number {

		//ไม่เช็คตัวที่ 0 และ ตัวสุดท้าย
		if index != 0 && index != (numberis_string_len-1) {

			//ตัวก่อนหน้ามัน 1 ลำดับ
			var Check_digit_digit_one uint
			Check_digit_digit_one = tmp_number[index-1]

			//ตัวถัดไป 1 ลำดับ
			var Check_digit_digit_two uint
			if index+1 >= numberis_string_len {
				Check_digit_digit_two = 0
			} else {
				Check_digit_digit_two = tmp_number[index+1]
			}

			//เลขตัวมันเอง +1 , -1
			var Number_plus = number + 1
			var Number_minus = number - 1

			// fmt.Println("บวก ROUND : ", number, " => ตัวเลขก่อนหน้ามัน ", Check_digit_digit_one, " ตัวเลขมันบวกหนึ่ง ", Number_plus, ", ตัวเลขหลังมัน ", Check_digit_digit_two)
			// fmt.Println("ลบ ROUND : ", number, " => ตัวเลขก่อนหน้ามัน ", Check_digit_digit_one, " , ตัวเลขมันลบหนึ่ง ", Number_minus, " , ตัวเลขหลังมัน ", Check_digit_digit_two)

			// (ข้อ3) เลขห้ามเรียงทั้ง จากมากไปน้อย และ จากน้อยไปมาก
			if (Number_plus == Check_digit_digit_two && Number_minus == Check_digit_digit_one) ||
				(Number_plus == Check_digit_digit_one && Number_minus == Check_digit_digit_two) {
				fmt.Println("Error : Number of tiers ")
				return
			}
		}

		// เก็บเข้า array กรณี ตัวปัจจุบัน กับ ตัวก่อนหน้าซ้ำกัน
		if keepnumber == number {
			tmpresult[number] = tmpresult[number] + 1
		}

		// เก็บตัวเลขก่อนหน้าเอาไว้
		keepnumber = number
	}

	// (ข้อ4) เลขชุดห้ามซ้ำเกินสองชุด
	if len(tmpresult) > 2 {
		fmt.Println("Error : Number duplicate number set")
		return
	}

	// (ข้อ2) check ซ้ำกันสองตัวไม่ให้ผ่าน , แต่มากกว่าสองตัวผ่าน
	for _, count := range tmpresult {
		if count > 2 {
			fmt.Println("Error : Number duplicate number")
			return
		}
	}

	fmt.Println("!!!!! Validate success !!!!!")
}
