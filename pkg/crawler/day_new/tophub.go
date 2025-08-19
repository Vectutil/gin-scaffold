package day_new

import (
	"context"
	"gin-scaffold/pkg/crawler/day_new/huxiu"
	"gin-scaffold/pkg/crawler/day_new/thepaper"
)

func RunNews() {
	ctx := context.Background()
	huxiu.HuxiuNews(ctx)
	thepaper.WorldNews(ctx)
}
