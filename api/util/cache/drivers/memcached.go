package drivers

type Memcached struct {
	init func()
}

func (mem Memcached) Get(key string) {
	init := mem.init
	init()
	{

	}
}

func (mem Memcached) Set(key string, data interface{}, timeout int8) {

}

func (mem Memcached) Inc(key string) {

}

func (mem Memcached) IncBy(key string, step int) {

}

func (mem Memcached) Dec(key string, step int8) {

}

func (mem Memcached) Del(key string) {

}

func (mem Memcached) Pull(key string) {

}
