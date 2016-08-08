package config

type DCConfig struct {
        List []DCList `json:List`
}

type DCList struct {
	Name  string
        Master   []System
//        Slave    []System
}

type System struct {
        Ip       string
}
