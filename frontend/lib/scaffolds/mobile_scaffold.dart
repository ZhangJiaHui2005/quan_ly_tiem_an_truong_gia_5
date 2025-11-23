import 'package:flutter/material.dart';

class MobileScaffold extends StatefulWidget {
  const MobileScaffold({super.key});

  @override
  State<MobileScaffold> createState() => _MobileScaffoldState();
}

class _MobileScaffoldState extends State<MobileScaffold> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.grey[900],
      ),

      backgroundColor: Colors.grey[300],
      drawer: Drawer(
        backgroundColor: Colors.grey[300],
        child: Column(
          children: [
            DrawerHeader(
              child: Text("Food Restaurant Bill Management App"),
            ),

            ListTile(
              leading: Icon(Icons.dashboard),
              title: Text("Dashboard"),
            ),

            ListTile(
              leading: Icon(Icons.note),
              title: Text("Bills"),
            ),

            ListTile(
              leading: Icon(Icons.gif_box),
              title: Text("Items"),
            ),

            ListTile(
              leading: Icon(Icons.category),
              title: Text("Categories"),
            ),
          ],
        ),

        
      ),
    );
  }
}