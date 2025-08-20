package common

import (
	"context"
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"gin-scaffold/pkg/robot"
	"gin-scaffold/pkg/utils"
	"gin-scaffold/pkg/windows_send"
	"math"
)

func CallRobot(currentPrice, changePrice, ratio, stockCode, stockName string) {
	botUrl := config.Cfg.FSRobot.StockRobot

	if ratio == "" {
		fmt.Println("未找到涨跌幅数据")
		return
	}
	if val, ok := StockValue[stockCode]; ok {
		subValue := utils.ExtractPercentageNumber(val) - utils.ExtractPercentageNumber(ratio)
		msg := fmt.Sprintf("[%s]%s(%s) 当前涨幅:%s  当前价:%s", ratio, stockName, stockCode, changePrice, currentPrice)
		if subValue > 0.5 {
			msg = fmt.Sprintf("↓[%s]%s(%s) 当前涨幅:%s  当前价:%s", ratio, stockName, stockCode, changePrice, currentPrice)
		}
		if subValue < -0.5 {
			msg = fmt.Sprintf("↑[%s]%s(%s) 当前涨幅:%s  当前价:%s", ratio, stockName, stockCode, changePrice, currentPrice)
		}
		logger.Logger.Info(msg)
		if math.Abs(subValue) >= 0.5 {
			robot.SendFeishuRobotWithUrl(context.Background(), botUrl, msg, robot.MsgTypeText)
			windows_send.SendMsgToWindows(WindowsName, msg)
			StockValue[stockCode] = ratio
		}
		return
	}
	StockValue[stockCode] = ratio
	msg := fmt.Sprintf("[%s]%s(%s) 当前涨幅:%s  当前价:%s", ratio, stockName, stockCode, changePrice, currentPrice)
	robot.SendFeishuRobotWithUrl(context.Background(), botUrl, msg, robot.MsgTypeText)
	windows_send.SendMsgToWindows(WindowsName, msg)
	logger.Logger.Info(msg)
}
