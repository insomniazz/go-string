package string

func Ucfirst(name string) string {
	var uppstr string

	//如果只是ASCII可以直接遍历字符串，要支持非ASCII字符就要转换为【】rune
	//for i, i2 := range name {
	//	fmt.Println(i)
	//	fmt.Println(i2)
	//}

	//将字符串转为Unicode对应编码顺序的切片，支持非ASCII的字符
	sName := []rune(name)

	for k, v := range sName {
		if k == 0 && v >= 97 && v <= 122 {
			v -= 32
		}
		uppstr += string(v)
	}

	return uppstr
}
