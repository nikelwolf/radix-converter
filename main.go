package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"

	"github.com/nikelwolf/radix-converter-go-qt/converter"
)

type RadixConverter struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string, uint64, uint64) string `slot:"convertButtonClicked"`
}

func (rc *RadixConverter) init() {
	rc.ConnectConvertButtonClicked(rc.convertButtonClickHandler)
}

func (rc *RadixConverter) convertButtonClickHandler(number string, from, to uint64) string {
	result, err := converter.ConvertNumberToAnotherRadix(number, from, to)
	if err != nil {
		return err.Error()
	}
	return result
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("Material")

	engine := qml.NewQQmlApplicationEngine(nil)

	rc := NewRadixConverter(nil)
	rc.QmlRegisterType()
	engine.RootContext().SetContextProperty("_converter", rc)

	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
