package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var (
	tinyOriginFile string // 去除重复和排序的文件路径
	shouldSort     bool   // 是否需要排序
)

// tinyCmd represents the tiny command
var tinyCmd = &cobra.Command{
	Use:   "tiny",
	Short: "去除文件重复行并排序。",
	Long: `文件去除重复，排序
注意：文件内容会被修改`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			fileContent          []byte
			result               []string
			tinyBeforeLine       int      // 修剪前行数
			tinyBeforeUnfoldLine int      // 修剪前展开行数
			tinyAfterLine        int      // 修剪后行数
			duplicates           []string // 重复行内容
			err                  error
		)
		// 读取源文件
		if fileContent, err = readFile(tinyOriginFile); err != nil {
			fmt.Println("源文件读取错误")
			goto ERR
		}
		tinyBeforeLine = len(strings.Split(string(fileContent), "\n")) // 修改前文件行数
		// 切割文件内容
		result = splitContents(fileContent)
		tinyBeforeUnfoldLine = len(result) // 统计修剪前展开行数

		// 去除重复
		result, duplicates = removeDuplicateElement(result)
		tinyAfterLine = len(result) // 统计修剪前行数

		// 排序
		if shouldSort {
			sort.Slice(result, func(i, j int) bool {
				return result[i] < result[j]
			})
		}

		// 重新写入到源文件
		if err = writeFile(tinyOriginFile, result); err != nil {
			fmt.Println("目标文件写入错误")
			goto ERR
		}

		fmt.Printf(`运行成功！

修剪前共%d行，修剪前展开%d行，修剪后%d行，修复重复%d行。
其中重复的为：%v

保存在%s里
`,
			tinyBeforeLine,
			tinyBeforeUnfoldLine,
			tinyAfterLine,
			tinyBeforeUnfoldLine-tinyAfterLine,
			duplicates,
			tinyOriginFile,
		)
		return
	ERR:
		fmt.Println(err.Error())
	},
}

func init() {
	rootCmd.AddCommand(tinyCmd)
	tinyCmd.Flags().StringVarP(&tinyOriginFile, "origin", "o", "./dist.txt", "提供去重和排序的文件")
	tinyCmd.Flags().BoolVarP(&shouldSort, "sort", "s", true, "去重后是否需要排序")
}

// 切割文件内容
func splitContents(fileContent []byte) (result []string) {
	var (
		item string
		tmp  string
	)
	// 通过 \n 分割
	for _, item = range strings.Split(string(fileContent), "\n") {
		item = strings.Trim(item, "")
		if item == "" {
			continue
		}
		// 通过空格分割
		for _, tmp = range strings.Split(item, " ") {
			if item == "" {
				continue
			}
			result = append(result, tmp)
		}

	}

	return
}
