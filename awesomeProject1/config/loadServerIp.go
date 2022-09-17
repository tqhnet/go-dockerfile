package config

var Ip string

func init() {
	//var err error
	//Ip, err = getLocalIp()
	//if err != nil {
	//	panic(err)
	//}
	//
	//if Ip == "" {
	//	panic("get local ip failed")
	//}
}

////获取本地ip地址
//func getLocalIp() (string, error) {
//	addresses, err := net.InterfaceAddrs()
//	if err != nil {
//		return "", err
//	}
//
//	reg := regexp.MustCompile(`^192|10.*`)
//	for _, addr := range addresses {
//		//检查ip地址判断是否回环地址
//		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
//			if ipNet.IP.To4() != nil && reg.MatchString(ipNet.IP.String()) {
//				return ipNet.IP.String(), nil
//			}
//		}
//	}
//	return "", errors.New("can not find the ip address")
//}
