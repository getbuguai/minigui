package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2"
)

func buildPrefWindow() fyne.Window {
	indexWindow := mainApp.NewWindow("Index Windows")
	indexWindow.SetIcon(pngIconResource)
	indexWindow.SetFixedSize(true)
	indexWindow.CenterOnScreen()
	indexWindow.Resize(fyne.NewSize(600, 400))
	indexWindow.SetCloseIntercept(func() { // 关闭的操作执行的任务
		indexWindow.Hide()
	})

	indexWindow.SetContent(
		container.NewScroll(
			canvas.NewText(
				fmt.Sprintf("text 中文 %s", version),
				color.Black)))

	return indexWindow
}
