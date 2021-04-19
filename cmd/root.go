package cmd

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/marchar/fabric-entry/server"
	"github.com/marchar/fabric-entry/utils"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// 根命令
var rootCmd = &cobra.Command{
	Use:   "entry",
	Short: "启动Fabric网络管理服务器",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// 启动网络服务器
		startServer()
	},
}

// 启动网络服务器
func startServer() {
	var address = viper.GetString("server.address")
	var webServer = server.New(address)
	logrus.Infof("正在%s上启动服务器监听...\n", address)
	var wg = &sync.WaitGroup{}

	// 等待Interrupt信号，并优雅关闭服务器
	utils.WaitSignals(
		func() {
			wg.Add(1)
			defer wg.Done()
			webServer.Start()
		},
		os.Interrupt,
	)

	// 优雅关闭服务器
	webServer.Stop(context.Background())
	logrus.Info("正在关闭HTTP服务器...")

	// 等待服务器退出
	wg.Wait()
	logrus.Info("HTTP服务器已关闭")
}

// 执行命令
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件 (默认读取 $HOME/.cmd.yaml)")
}

// 初始化配置
func initConfig() {
	if cfgFile != "" {
		// 执行参数中指定的配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 从主目录中查找配置
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// 添加查找配置文件的路径
		viper.AddConfigPath(home)
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // 读取环境变量

	// 找到配置文件后读入
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "配置文件路径:", viper.ConfigFileUsed())
	}
}
