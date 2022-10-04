package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var (
		ctlGroup    = widgets.NewQGroupBox2("Widgets Test for Controls", nil)
		ctlLabel    = widgets.NewQLabel2("Type:", nil, 0)
		ctlComboBox = widgets.NewQComboBox(nil)
		ctlLineEdit = widgets.NewQLineEdit(nil)
	)
	ctlComboBox.AddItems([]string{"Left", "Center", "Right"})
	ctlLineEdit.SetPlaceholderText("Test")

	ctlComboBox.ConnectCurrentIndexChanged(func(index int) { ctlChanged(ctlLineEdit, index) })

	var ctlLayout = widgets.NewQGridLayout2()
	ctlLayout.AddWidget2(ctlLabel, 0, 0, 0)
	ctlLayout.AddWidget2(ctlComboBox, 0, 1, 0)
	ctlLayout.AddWidget3(ctlLineEdit, 1, 0, 1, 2, 0)
	ctlGroup.SetLayout(ctlLayout)

	var layout = widgets.NewQGridLayout2()
	layout.AddWidget2(ctlGroup, 2, 0, 0)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Test Controls")

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

	window.Show()

	widgets.QApplication_Exec()
}

func ctlChanged(ctlLineEdit *widgets.QLineEdit, index int) {
	switch index {
	case 0:
		{

			ctlLineEdit.SetAlignment(core.Qt__AlignLeft)
		}

	case 1:
		{
			ctlLineEdit.SetAlignment(core.Qt__AlignCenter)
		}

	case 2:
		{
			ctlLineEdit.SetAlignment(core.Qt__AlignRight)
		}
	}
}
