package main

func reverseString(s []byte) {
	var temp byte
	//双指针
	if (len(s) == 1) {
		return
	}
	var index = len(s) / 2
	//0 1 2 3 4
	//5
	//0 1
	var end = len(s) - 1
	for i := 0; i < index; i++ {
		temp = s[i];
		s[i] = s[end-i]
		s[end-i] = temp
	}

}
