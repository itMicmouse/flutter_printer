# flutter_printer

This Go package implements the host-side of the Flutter [flutter_printer](https://github.com/itMicmouse/flutter_printer) plugin.

## Usage

Import as:

```go
import flutter_printer "github.com/itMicmouse/flutter_printer/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&flutter_printer.FlutterPrinterPlugin{}),
```
