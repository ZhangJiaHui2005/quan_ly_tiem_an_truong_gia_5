import 'package:flutter/material.dart';

class ResponsiveLayout extends StatelessWidget {
  final Widget mobileScaffold;
  final Widget desktopScaffold;

  const ResponsiveLayout({super.key, required this.desktopScaffold, required this.mobileScaffold});

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(builder: (context, constrants) {
      if (constrants.maxWidth < 500) {
        return mobileScaffold;
      } else {
        return desktopScaffold;
      }
    });
  }
}