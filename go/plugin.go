package flutter_printer

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"os/exec"
)

const channelName = "flutter_printer"

// FlutterPrinterPlugin implements flutter.Plugin and handles method.
type FlutterPrinterPlugin struct{}

var _ flutter.Plugin = &FlutterPrinterPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *FlutterPrinterPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	channel.HandleFunc("startChrome", p.startChrome)
    channel.HandleFunc("startChromeWindows", p.startChromeWindows)
	return nil
}

func (p *FlutterPrinterPlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}

//mac 调用
func (p *FlutterPrinterPlugin) startChrome(arguments interface{}) (reply interface{}, err error) {
	url := arguments.(map[interface{}]interface{})["url"].(string)
	prescriptionLsId := arguments.(map[interface{}]interface{})["prescriptionLsId"].(string)
	changeval := arguments.(map[interface{}]interface{})["token"].(string)
	exec.Command(`open`, url+"?prescriptionLsId="+prescriptionLsId+"&changeval="+changeval).Start()
	return "go-flutter " + flutter.PlatformVersion, nil
}

//windows 调用
func (p *FlutterPrinterPlugin) startChromeWindows(arguments interface{}) (reply interface{}, err error) {
	url := arguments.(map[interface{}]interface{})["url"].(string)
	prescriptionLsId := arguments.(map[interface{}]interface{})["prescriptionLsId"].(string)
	changeval := arguments.(map[interface{}]interface{})["token"].(string)

	// 有GUI调用
	exec.Command(`cmd`, `/c`, `start`, url+"?prescriptionLsId="+prescriptionLsId+"^&changeval="+changeval).Start()

	return "go-flutter " + flutter.PlatformVersion, nil
}
