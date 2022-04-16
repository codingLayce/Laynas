import 'package:client/model/file_node.dart';
import 'package:client/screens/components/app_bar.dart';
import 'package:client/screens/components/directory_view.dart';
import 'package:client/services/files_manager.dart';
import 'package:flutter/material.dart';

class MainPage extends StatelessWidget {
  const MainPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: LaynasAppBar.get("Stockage"),
      body: FutureBuilder<FileNode>(
        future: FileManager().getFileTree(),
        builder: (context, snapshot) {
          if (snapshot.hasError) {
            return _error();
          }

          if (snapshot.hasData) {
            return _content(context, snapshot.data!);
          }

          return _loading();
        }
      ),
    );
  }

  _content(BuildContext context, FileNode root) {
    return DirectoryView(directory: root);
  }

  _loading() {
    return const Center(child: CircularProgressIndicator());
  }

  _error() {
    return const Center(child: Text("Error while retrieving the file tree"));
  }
}