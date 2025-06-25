package main

import "fmt"

func demo4() {
	fmt.Println("============= demo 4 ==============")
	{
		// 创建map类型

		// 创建空map
		userinfo := make(map[string]string)
		userinfo["username"] = "张三"
		userinfo["age"] = "20"
		userinfo["sex"] = "男"
		fmt.Printf("%T: %v\n", userinfo, userinfo) //map[string]string: map[age:20 sex:男 username:张三]

		// 创建map时初始化元素
		userinfo1 := map[string]string{
			"username": "李四",
			"age":      "18",
			"sex":      "男",
		}
		fmt.Println(userinfo1)

		for k, v := range userinfo {
			fmt.Printf("%v %v\n", k, v)
		}

		v, ok := userinfo["age"]
		fmt.Println(v, ok) // 20 true

		// 使用delete()内建函数删除map中的元素
		delete(userinfo, "age")
		v, ok = userinfo["age"]
		fmt.Println(v, ok) //  false

		usersinfo := make([]map[string]string, 3, 3)
		fmt.Println(usersinfo[0] == nil) // true

		usersinfo[0] = make(map[string]string)
		usersinfo[0]["name"] = "小A"
		usersinfo[0]["age"] = "20"

		usersinfo[1] = make(map[string]string)
		usersinfo[1]["name"] = "小B"
		usersinfo[1]["age"] = "30"

		for _, user := range usersinfo {
			for k, v := range user {
				fmt.Printf("%v-%v ", k, v)
			}
			fmt.Println()
		}

		userinfo2 := make(map[string][]string)
		userinfo2["hobby"] = []string{
			"吃饭", "睡觉", "coding",
		}
		userinfo2["city"] = []string{
			"北京", "天津", "上海", "济南",
		}
		fmt.Println(userinfo2) //map[city:[北京 天津 上海 济南] hobby:[吃饭 睡觉 coding]]

	}

}
