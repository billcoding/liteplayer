import 'package:app/widgets/entertainment.dart';
import 'package:flutter/material.dart';
import 'widgets/movie.dart';
import 'widgets/tv.dart';

class Home extends StatefulWidget {
  const Home({super.key});

  @override
  State<Home> createState() => _HomeState();
}

const navBars = <_NavBarWrapper>[
  _NavBarWrapper(TVPage(), Icon(Icons.tv), 'TV'),
  _NavBarWrapper(MoviePage(), Icon(Icons.movie), 'Movie'),
  _NavBarWrapper(EntertainmentPage(), Icon(Icons.videocam), 'Entertainment'),
];
var widgets = navBars.map((a) => a.widget).toList();
var items = navBars.map((a) => a.item()).toList();

class _HomeState extends State<Home> {
  static int selectedIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Welcome to MyApp')),
      body: Center(child: widgets[selectedIndex]),
      bottomNavigationBar: BottomNavigationBar(
          items: items,
          currentIndex: selectedIndex,
          selectedItemColor: Colors.amber.shade500,
          onTap: (index) {
            selectedIndex = index;
            setState(() {});
          }),
    );
  }
}

class _NavBarWrapper {
  final Widget widget;
  final Icon icon;
  final String label;
  const _NavBarWrapper(this.widget, this.icon, this.label);
  BottomNavigationBarItem item() =>
      BottomNavigationBarItem(icon: icon, label: label);
}
