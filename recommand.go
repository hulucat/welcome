package main

import (
	"github.com/hulucat/utils"
)

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func recommandFor1(person *Person) (p *Product, err error) {
	utils.Debugf("Recommand for 1 person: %d, %s", person.Age, person.Gender)
	if person.Age < 10 {
		p = &Product{
			Id:   "27",
			Name: "统一冰糖雪梨",
			Img:  "http://img.ubox.cn/box-tmp/m/27/27.jpg?t=20140305101413",
		}
		return
	}
	switch person.Gender {
	case "Male":
		if person.Age < 18 {
			p = &Product{
				Id:   "1",
				Name: "可口可乐",
				Img:  "http://img.ubox.cn/box-tmp/m/1/1.jpg?t=20140305101413",
			}
		} else if person.Age < 30 {
			p = &Product{
				Id:   "5976",
				Name: "杜蕾斯热感超薄3只装",
				Img:  "http://img.ubox.cn/box-tmp/m/5976/5976.jpg?t=20140305101413",
			}
		} else if person.Age < 40 {
			p = &Product{
				Id:   "19251",
				Name: "取悦潮品4D经典名器少女",
				Img:  "http://img.ubox.cn/box-tmp/m/19251/19251.jpg?t=20140305101413",
			}
		} else if person.Age < 60 {
			p = &Product{
				Id:   "18471",
				Name: "XYZ男士劲能延时喷雾",
				Img:  "http://img.ubox.cn/box-tmp/m/18471/18471.jpg?t=20140305101413",
			}
		} else {
			p = &Product{
				Id:   "19250",
				Name: "取悦潮品4D经典名器处女",
				Img:  "http://img.ubox.cn/box-tmp/m/19250/19250.jpg?t=20140305101413",
			}
		}
	case "Female":
		if person.Age < 18 {
			p = &Product{
				Id:   "18459",
				Name: "费洛蒙男女通用香水",
				Img:  "http://img.ubox.cn/box-tmp/m/18459/18459.jpg?t=20140305101413",
			}
		} else if person.Age < 30 {
			p = &Product{
				Id:   "18528",
				Name: "好奇而已纯白诱惑",
				Img:  "http://img.ubox.cn/box-tmp/m/18528/18528.jpg?t=20140305101413",
			}
		} else if person.Age < 40 {
			p = &Product{
				Id:   "19308",
				Name: "美国进口MOVO迷恋型紧",
				Img:  "http://img.ubox.cn/box-tmp/m/19308/19308.jpg?t=20140305101413",
			}
		} else if person.Age < 60 {
			p = &Product{
				Id:   "18395",
				Name: "喵星人伊咪Y防水阴蒂按摩",
				Img:  "http://img.ubox.cn/box-tmp/m/18395/18395.jpg?t=20140305101413",
			}
		} else {
			p = &Product{
				Id:   "27",
				Name: "统一冰糖雪梨",
				Img:  "http://img.ubox.cn/box-tmp/m/27/27.jpg?t=20140305101413",
			}
		}
	}
	return
}

func recommandFor2(people []*Person) (p *Product, err error) {
	return &Product{
		Id:   "1",
		Name: "可口可乐",
		Img:  "http://img.ubox.cn/box-tmp/m/1/1.jpg?t=20140305101413",
	}, nil
}

func recommandForMulti(people []*Person) (p *Product, err error) {
	return &Product{
		Id:   "24",
		Name: "娃哈哈纯净水596ml",
		Img:  "http://img.ubox.cn/box-tmp/m/24/24.jpg?t=20140305101413",
	}, nil
}

func Recommand(filePath string, people []*Person) (err error) {
	dict := make(map[string]interface{}) //people: []*Person, product: Product
	var product *Product
	if people != nil && len(people) > 0 {
		switch len(people) {
		case 1:
			product, err = recommandFor1(people[0])
		case 2:
			product, err = recommandFor2(people)
		default:
			product, err = recommandForMulti(people)
		}
		if err != nil {
			utils.Errorf("Error recommand: %s", err.Error())
		} else {
			utils.Debugf("Product: %s", product.Name)
		}
	}
	dict["file"] = filePath
	dict["people"] = people
	dict["product"] = product
	if WebSocketConn != nil {
		if err = WebSocketConn.WriteJSON(dict); err != nil {
			utils.Errorf("Error write json to websocket: %s", err.Error())
		}
	} else {
		utils.Debugf("Websocket is nil")
	}
	return
}
