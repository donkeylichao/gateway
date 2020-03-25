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

/* 计算使用的数据容器 */
var comb []interface{}

/**
 获取当前请求匹配的转发设置
 */
func GetGatewayService(requestParam map[string]interface{}) interface{} {
	//获取当前请求路径数组
	currentRoute := GetCurrentRoute(requestParam)

	//获取数据库请求路径配置
	apiData := GetRouteConfig()

	for _,v := range currentRoute.([]interface{}) {
		apiDataMap := apiData.(map[string]interface{})
		if apiDataMap[v.(string)] != nil {
			return parseServicePath(requestParam["path"].(string),apiDataMap[v.(string)])
		}
	}
	return nil
}

/**
 获取数据库设置的所有url配置数组
 */
func GetRouteConfig() interface{} {

	routeConfig := beego.AppConfig.String("route.cache")
	if c := help.Redis.Get(routeConfig); c == nil {
		var url models.ServiceUrl
		var api models.ServiceApi
		urldata := url.List()
		apidata := api.List()

		returnData := getHandleApi(apidata, urldata)
		saveData, _ := json.Marshal(returnData)
		help.Redis.Put(routeConfig, saveData, time.Hour)
		return returnData
	} else {
		var data interface{}
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
func getHandleApi(api []*models.ServiceApi, url []*models.ServiceUrl) map[string]interface{} {
	data := make(map[string]interface{})
	u := getHandleUrl(url)
	for _, a := range api {
		formatA := make(map[string]interface{})
		formatA["Method"] = a.Method
		formatA["ServiceUrl"] = u[a.ServiceName]
		formatA["ApiAlias"] = a.ApiAlias
		formatA["ApiPath"] = a.ApiPath
		formatA["InnerPath"] = a.InnerPath

		data["["+a.Method+"]"+a.ApiAlias] = formatA
	}
	return data
}

/**
 获取格式化的url数据
 */
func getHandleUrl(url []*models.ServiceUrl) map[string]string {
	data := make(map[string]string)
	for _, u := range url {
		data[u.ServiceName] = u.ServiceUrl
	}
	return data
}

/**
 解析请求路由
 */
func GetCurrentRoute(requestParam map[string]interface{}) interface{} {

	method := requestParam["method"]
	path := requestParam["path"]
	placeHolder := beego.AppConfig.String("route.parser_placeholder")
	pathData := make(map[int]interface{})
	combine := strings.Split(strings.Trim(path.(string),"/"), "/")
	count := len(combine)

	if path == "/" {
		pathData[0] = "/"
		return pathData
	}

	//计算替换的数组下标
	pathData = processCombine(count, placeHolder)

	//拼接格式化的key数组
	var returnPath []interface{}
	originPath := "[" +method.(string)+"]" + path.(string)
	returnPath = append(returnPath,originPath)
	for _,v := range pathData {
		temp := make([]string, count)
		copy(temp,combine)
		for k,i := range v.(map[int]string) {
			temp[k] = i
		}
		combinePath := "["+method.(string)+"]"
		for _,v := range temp {
			combinePath += "/"+v
		}
		returnPath = append(returnPath,combinePath)
	}
	return returnPath
}

/**
 获取匹配路由替换数组
 */
func processCombine(count int, placeholder string) map[int]interface{} {
	combinineString := ""

	for i := 0; i < count ; i++ {
		combinineString += strconv.Itoa(i)
	}

	SliceClear(&comb) //清空数据，每次调用都初始化
	combinineArray := combinations(combinineString)//递归获取替换的下标数组
	newComBinineArray := combinineArray.([]interface{})
	combine := make(map[int]interface{})

	for k, v := range newComBinineArray {
		index := make(map[int]string)
		for _, v := range strings.Split(v.(string), ",") {
			kindex,_ := strconv.Atoi(v)
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
	if len(combine) <= 1 {

		comb = append(comb, combine)
	} else {
		strFirst := combine[0:1]
		content := combine[1: len(combine)]
		combTemp := combinations(content)
		newCombTemp := combTemp.([]interface{})
		comb = append(comb, strFirst)
		for _, v := range newCombTemp {
			comb = append(comb, (strFirst + "," + v.(string)))
		}
	}

	return comb
}

/**
 计算有参数的路由转发地址
 */
func parseServicePath(path string ,matchRoute interface{}) interface{} {

	//请求path
	pathOrgin := strings.Trim(path,"/")
	//网关配置path
	routePath := strings.Trim(matchRoute.(map[string]interface{})["ApiAlias"].(string),"/")
	//转发配置path
	apiPath := strings.Trim(matchRoute.(map[string]interface{})["ApiPath"].(string),"/")

	//分别存放原始数据的map
	pathOrginSlice := strings.Split(pathOrgin,"/")
	routePathSlice := strings.Split(routePath,"/")
	apiPathSlice := strings.Split(apiPath,"/")

	var pathReplaceSlice []string
	for k,v := range routePathSlice {
		if v[0:1] == "[" {
			pathReplaceSlice = append(pathReplaceSlice,pathOrginSlice[k])
		}
	}

	returnPath := ""
	for k,v := range apiPathSlice {
		if v[0:1] == "[" {
			apiPathSlice[k] = pathReplaceSlice[0]
			pathReplaceSlice = pathReplaceSlice[1:]
		}
		returnPath +=  "/"+apiPathSlice[k]
	}

	//组装request请求转发的真实路径
	return matchRoute.(map[string]interface{})["ServiceUrl"].(string) + returnPath
}

/**
 清空切片内容
 */
func SliceClear(s *[]interface{}) {
	*s = (*s)[0:0]
}

