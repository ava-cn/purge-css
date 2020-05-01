/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var (
	originFile string // 源文件
	distFile   string // 目标文件
	parseLine  int    // 分析出的行
)

// 正则表达式获取页面class和id的值
const classAndIDRegex = `(id|class)=["'](.*?)["']`

// whiteListFilterCmd represents the whiteListFilter command
var whiteListFilterCmd = &cobra.Command{
	Use:   "white-list-filter",
	Short: "提取页面class和id。",
	Long: `从提供的源文件中提取需要的class和id。
去重、按字母排序后覆盖式写入到目标文件。`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			fileContent []byte   // 文件内容
			result      []string // 正则匹配后的内容
			err         error    // 错误
		)

		// 获取页面文件内容
		if fileContent, err = readFile(originFile); err != nil {
			goto ERR
		}

		// 正则表达式分析出页面对应class和id的值并写入到对应的临时文件
		result = regexContent(fileContent)

		// 去重
		result = removeRepByMap(result)

		// 输出排序
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})

		// 写入到目标文件
		if err = writeFile(distFile, result); err != nil {
			goto ERR
		}

		fmt.Printf("已经分析出需要的class或id，总共%d个，保存在%s里", parseLine, distFile)
		return

	ERR:
		fmt.Println(err.Error())
	},
}

func init() {
	rootCmd.AddCommand(whiteListFilterCmd)
	whiteListFilterCmd.PersistentFlags().StringVarP(&originFile, "origin", "o", "./code.html", "提供原始文件")
	whiteListFilterCmd.PersistentFlags().StringVarP(&distFile, "dist", "d", "./dist.txt", "提供目标文件")
}

// removeRepByMap: 通过map主键唯一的特性过滤重复元素
func removeRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		// fmt.Println(e)
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// readFile 读文件
func readFile(fileName string) (result []byte, err error) {
	if result, err = ioutil.ReadFile(fileName); err != nil {
		return nil, err
	}
	return result, nil
}

// writeFile 写文件
func writeFile(fileName string, content []string) (err error) {
	var (
		newFile *os.File
		str     string
	)
	if newFile, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
		fmt.Printf("failed to create new file err: %s", err.Error())
		return
	}
	defer newFile.Close()

	for _, str = range content {
		if str == "" { // 去除空格
			continue
		}
		if _, err = newFile.WriteString(str + "\n"); err != nil {
			return
		}
		parseLine++
	}

	return
}

// 正则匹配内容
func regexContent(fileContent []byte) (result []string) {
	var (
		re       *regexp.Regexp
		matches  [][][]byte
		match    [][]byte
		subMatch string
	)
	re = regexp.MustCompile(classAndIDRegex)
	matches = re.FindAllSubmatch(fileContent, -1)
	for _, match = range matches {
		// 按照空格分割class或id字符串
		for _, subMatch = range strings.Split(string(match[2]), " ") {
			result = append(result, subMatch)
		}
	}
	return
}
