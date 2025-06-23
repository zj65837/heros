/**
* @Author: Ramoncjs
* @Date: 2021/8/20 21:03
 */
package loginutil

import (
	"errors"
	"fmt"
)

// 定义服务器 URL 的结构体
type urlStruct struct {
	QuickLoginUrl string // 快速登录 URL
	GameUrl       string // 游戏服务器 URL
}

// 全局变量：存储服务器列表信息
var srv map[string]interface{}

// 全局变量：存储区服代码到 URL 的映射
var rt = make(map[string]urlStruct)

// 定义服务器主机地址列表的类型
type BZSRVLIST map[string][]string

// 全局变量：存储服务器主机地址列表
var bsvrlst BZSRVLIST

// ChooseServer 根据区服代码选择对应的服务器 URL
// 参数 servercode: 区服代码 (如 "g1", "h2" 等)
// 返回值: 快速登录URL, 游戏URL, 错误信息
func ChooseServer(servercode string) (string, string, error) {
	// 遍历服务器列表 srv
	for k, v := range srv {
		// 获取服务器信息中的第6个元素（索引5）作为区服名称
		switch v.([]interface{})[5].(string) {
		// 根据区服名称映射到对应的区服代码
		case "官方1区":
			rt["g1"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "官方2区":
			rt["g2"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "官方3区":
			rt["g3"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服1区":
			rt["h1"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服2区":
			rt["h2"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服3区":
			rt["h3"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服4区":
			rt["h4"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服5区":
			rt["h5"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服6区":
			rt["h6"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服7区":
			rt["h7"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服8区":
			rt["h8"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服9区":
			rt["h9"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服10区":
			rt["h10"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服11区":
			rt["h11"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服12区":
			rt["h12"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}
		case "混服13区":
			rt["h13"] = urlStruct{QuickLoginUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[2].(string)), GameUrl: fmt.Sprintf("%s:%s", bsvrlst[k][0], v.([]interface{})[3].(string))}

			// 未知区服名称处理
		default:
			return "", "", errors.New("[-] 区服代码错误.")
		}
	}
	
	// 返回请求的区服代码对应的 URL
	return rt[fmt.Sprintf("%s", servercode)].QuickLoginUrl, rt[fmt.Sprintf("%s", servercode)].GameUrl, nil

}
