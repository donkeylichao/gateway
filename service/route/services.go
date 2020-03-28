package route

import (
	"gateway/models"
	"gateway/help"
	"time"
	"github.com/astaxie/beego"
	"encoding/json"
	"strings"
	"strconv"
)

/**
 获取当前请求匹配的转发设置
 */
func GetGatewayService(requestParam map[string]interface{}) string {
	//获取当前请求路径数组
	currentRoute := getCurrentRoute(requestParam)

	//获取数据库请求路径配置
	apiData := getRouteConfig()

	for _, v := range currentRoute {
		if apiData[v.(string)] != nil {
			return parseServicePath(requestParam["path"].(string), apiData[v.(string)])
		}
	}
	return ""
}

/**
 获取数据库设置的所有url配置数组
 */
func getRouteConfig() map[string]interface{} {

	routeConfig := beego.AppConfig.String("route::cache")
	if c := help.Redis.Get(routeConfig); c == nil {
		var url models.ServiceUrl
		var api models.ServiceApi
		urldata := url.IdAndUrlList()
		apidata := api.List()

		returnData := getHandleApi(apidata, urldata)
		saveData, _ := json.Marshal(returnData)
		help.Redis.Put(routeConfig, saveData, time.Hour)
		return returnData
	} else {
		var data map[string]interface{}
		if err := json.Unmarshal(c.([]byte), &data); err != nil {
			panic(err)
		} else {
			return data
		}
	}
}

/**
 获取格式化后的api数据
 */
func getHandleApi(api []*models.ServiceApi, url map[int]string) map[string]interface{} {
	data := make(map[string]interface{})
	for _, a := range api {
		formatA := make(map[string]interface{})
		formatA["Method"] = a.Method
		formatA["ServiceUrl"] = url[a.ServiceUrlId]
		formatA["ApiAlias"] = a.ApiAlias
		formatA["ApiPath"] = a.ApiPath

		data["["+a.Method+"]"+a.ApiAlias] = formatA
	}
	return data
}

/**
 解析请求路由
 */
func getCurrentRoute(requestParam map[string]interface{}) []interface{} {

	method := requestParam["method"]
	path := requestParam["path"]
	placeHolder := beego.AppConfig.String("route::parser_placeholder")

	combine := strings.Split(strings.Trim(path.(string), "/"), "/")
	count := len(combine)

	if path == "/" {
		return []interface{}{"[" + method.(string) + "]/"}
	}

	//计算替换的数组下标
	pathData := processCombine(count, placeHolder)

	//拼接格式化的key数组
	var returnPath []interface{}
	originPath := "[" + method.(string) + "]" + path.(string)
	returnPath = append(returnPath, originPath)
	for _, v := range pathData {
		temp := make([]string, count)
		copy(temp, combine)
		for k, i := range v.(map[int]string) {
			temp[k] = i
		}
		methodPath := "[" + method.(string) + "]"
		for _, v := range temp {
			methodPath += "/" + v
		}
		returnPath = append(returnPath, methodPath)
	}
	return returnPath
}

/**
 获取匹配路由替换数组
 */
func processCombine(count int, placeholder string) map[int]interface{} {
	combinineString := ""

	for i := 0; i < count; i++ {
		combinineString += strconv.Itoa(i)
	}

	combinineArray := combinations(combinineString) //递归获取替换的下标数组
	newComBinineArray := combinineArray.([]interface{})
	combine := make(map[int]interface{})

	for k, v := range newComBinineArray {
		index := make(map[int]string)
		for _, v := range strings.Split(v.(string), ",") {
			kindex, _ := strconv.Atoi(v)
			index[kindex] = placeholder
		}
		combine[k] = index
	}

	return combine
}

/**
 计算可替换的路由
 */
func combinations(combine string) interface{} {

	if combine == "" {
		return nil
	}

	var comb []interface{}
	if len(combine) <= 1 {
		comb = append(comb, combine)
	} else {
		first := combine[0:1]
		remain := combine[1:]
		combTemp := combinations(remain)
		newCombTemp := combTemp.([]interface{})
		comb = append(comb, first)
		comb = append(comb, newCombTemp...)
		for _, v := range newCombTemp {
			comb = append(comb, (first + "," + v.(string)))
		}
	}

	return comb
}

/**
 计算有参数的路由转发地址
 */
func parseServicePath(path string, matchRoute interface{}) string {

	//请求path
	pathOrgin := strings.Trim(path, "/")
	//网关配置path
	routePath := strings.Trim(matchRoute.(map[string]interface{})["ApiAlias"].(string), "/")
	//转发配置path
	apiPath := strings.Trim(matchRoute.(map[string]interface{})["ApiPath"].(string), "/")

	//分别存放原始数据的map
	pathOrginSlice := strings.Split(pathOrgin, "/")
	routePathSlice := strings.Split(routePath, "/")
	apiPathSlice := strings.Split(apiPath, "/")

	var pathReplaceSlice []string
	placeHolder := beego.AppConfig.String("route::parser_placeholder")
	for k, v := range routePathSlice {
		if v == placeHolder {
			pathReplaceSlice = append(pathReplaceSlice, pathOrginSlice[k])
		}
	}

	returnPath := ""
	itemUrl := ""
	for _, v := range apiPathSlice {
		if v == placeHolder {
			itemUrl = pathReplaceSlice[0]
			pathReplaceSlice = pathReplaceSlice[1:]
		} else {
			itemUrl = v
		}
		returnPath += "/" + itemUrl
	}

	//组装request请求转发的真实路径
	return matchRoute.(map[string]interface{})["ServiceUrl"].(string) + returnPath
}
