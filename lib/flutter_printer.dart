import 'dart:async';
import 'dart:io';

import 'package:flutter/services.dart';

class FlutterPrinter {
  static const MethodChannel _channel =
      const MethodChannel('flutter_printer');

  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }

  static void startChrome(String url,String prescriptionLsId,String token,){
    if (Platform.isAndroid) {

    } else if (Platform.isIOS) {

    } else if(Platform.isMacOS){
      _channel.invokeMethod('startChrome',{"url": url,"prescriptionLsId": prescriptionLsId,"token": token});
    } else if(Platform.isWindows){
      _channel.invokeMethod('startChromeWindows',{"url": url,"prescriptionLsId": prescriptionLsId,"token": token});
    } else if(Platform.isLinux){

    } else if(Platform.isFuchsia){

    }
  }
}
