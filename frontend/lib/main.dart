import 'package:flutter/material.dart';
import 'package:frontend/layout/responsive_layout.dart';
import 'package:frontend/scaffolds/desktop_scaffold.dart';
import 'package:frontend/scaffolds/mobile_scaffold.dart';

void main(List<String> args) {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: ResponsiveLayout(
        desktopScaffold: DesktopScaffold(),
        mobileScaffold: MobileScaffold(),
      ),
    );
  }
}
