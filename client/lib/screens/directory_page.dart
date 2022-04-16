import 'package:client/model/file_node.dart';
import 'package:client/screens/components/directory_view.dart';
import 'package:flutter/material.dart';

import 'components/app_bar.dart';

class DirectoryPage extends StatelessWidget {
  final FileNode directory;

  const DirectoryPage({Key? key, required this.directory}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: LaynasAppBar.get(directory.name),
        body: DirectoryView(directory: directory)
    );
  }

}