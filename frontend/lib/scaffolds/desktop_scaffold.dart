import 'package:flutter/material.dart';
import 'package:frontend/pages/bills_page.dart';
import 'package:frontend/pages/categories_page.dart';
import 'package:frontend/pages/dashboard_page.dart';
import 'package:frontend/pages/items_page.dart';

class DesktopScaffold extends StatefulWidget {
  const DesktopScaffold({super.key});

  @override
  State<DesktopScaffold> createState() => _DesktopScaffoldState();
}

class _DesktopScaffoldState extends State<DesktopScaffold> {
  int selectedIndex = 0;

  final List<Widget> pages = [
    DashboardPage(),
    BillsPage(),
    ItemsPage(),
    CategoriesPage(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.grey[900],
        title: Text(
          "Truong Gia 5 Management App",
          style: TextStyle(color: Colors.white),
        ),
        centerTitle: true,
      ),
      backgroundColor: Colors.grey[300],
      body: Row(
        children: [
          Drawer(
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadiusGeometry.only(
                topRight: Radius.circular(16),
                bottomRight: Radius.circular(16),
              ),
            ),
            shadowColor: Colors.black,
            backgroundColor: Colors.grey[300],
            child: Column(
              children: [
                ListTile(
                  leading: Icon(Icons.dashboard),
                  title: Text("Dashboard"),
                  selected: selectedIndex == 0,
                  onTap: () => setState(() {
                    selectedIndex = 0;
                  }),
                ),

                ListTile(
                  leading: Icon(Icons.note),
                  title: Text("Bills"),
                  selected: selectedIndex == 1,
                  onTap: () => setState(() {
                    selectedIndex = 1;
                  }),
                ),

                ListTile(
                  leading: Icon(Icons.gif_box),
                  title: Text("Items"),
                  selected: selectedIndex == 2,
                  onTap: () => setState(() {
                    selectedIndex = 2;
                  }),
                ),

                ListTile(
                  leading: Icon(Icons.category),
                  title: Text("Categories"),
                  selected: selectedIndex == 3,
                  onTap: () => setState(() {
                    selectedIndex = 3;
                  }),
                ),
              ],
            ),
          ),

          Expanded(
            child: IndexedStack(index: selectedIndex, children: pages),
          ),
        ],
      ),
    );
  }
}
