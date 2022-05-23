package person

/**
 * @Author: Alan Yin
 * @Date: 2022/5/20
 * @Desc:
 */

type Person struct {
	age int64
}

func (p *Person) older(other *Person) bool {
	return p.age > other.age
}

func main() {

}
