import 'package:flutter/material.dart';

class TopPage extends StatelessWidget {
  const TopPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Text(
        "aiueo",
        style: TextStyle(
            fontSize: 80, fontWeight: FontWeight.w600, color: Colors.white),
      ),
      backgroundColor: Color.fromRGBO(206, 229, 208, 1),
    );
  }
}
