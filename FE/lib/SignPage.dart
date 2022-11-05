import 'package:amplify_auth_cognito/amplify_auth_cognito.dart';
import 'package:amplify_authenticator/amplify_authenticator.dart';
import 'package:amplify_flutter/amplify_flutter.dart';
import 'package:flutter/material.dart';

import 'amplifyconfiguration.dart';

class SignPage extends StatefulWidget {
  const SignPage({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return _SignPageState();
  }
}

class _SignPageState extends State {
  @override
  void initState() {
    super.initState();
    _configureAmplify();
  }

  void _configureAmplify() async {
    try {
      await Amplify.addPlugin(AmplifyAuthCognito());
      await Amplify.configure(amplifyconfig);
      print('Successfully configured');
    } on Exception catch (e) {
      print('Error configuring Amplify: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Authenticator(
      child: MaterialApp(
        builder: Authenticator.builder(),
        home: Scaffold(
          appBar: AppBar(title: const Text('ホーム')),
          body: const Center(
            child: Text('ログイン成功！', style: TextStyle(fontSize: 32.0)),
          ),
        ),
      ),
    );
  }
}
