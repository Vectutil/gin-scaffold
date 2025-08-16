package robot

import "context"

func CallQWAssistant(ctx context.Context, data, QWRobotMsgType string) {
	SendFeishuRobot(ctx, data)
	//SendQWRobot(ctx, data, QWRobotMsgType)
}
