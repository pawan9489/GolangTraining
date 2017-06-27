package main

import (
	"crypto/sha1"
	"fmt"
)

func main()  {
	s := `
	 The Code of Alabama and Executive Order. It has an authorized strength of 1,000 members and is organized on the United States Army structural pattern. The ASDF is under the control of the Governor of Alabama, as the state's Commander in Chief,
	`
	fmt.Println(s)
	fmt.Println(len(s))
	h := sha1.New()
	fmt.Println(h)
	h.Write([]byte(s))
	fmt.Println(h)
	bs := h.Sum(nil)
	fmt.Println(bs)
	fmt.Printf("%x\n",bs)

	temp_arr := []int{226,246 ,124 ,119 ,35 ,104, 172, 222, 238, 106, 34 ,66 ,197, 53, 198, 204, 40, 216, 224, 237}
	var sum int
	for _,v := range temp_arr {
		sum += v
	}
	fmt.Println(sum)
	fmt.Printf("%x\n",sum)
}
