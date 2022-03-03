import 'package:client/model/file_node.dart';
import 'package:client/screens/components/file_view.dart';
import 'package:client/screens/directory_page.dart';
import 'package:flutter/material.dart';

class DirectoryView extends StatefulWidget {
  final FileNode directory;
  
  const DirectoryView({Key? key, required this.directory}) : super(key: key);

  @override
  State<StatefulWidget> createState() => _DirectoryView();
}

class _DirectoryView extends State<DirectoryView> {
  @override
  Widget build(BuildContext context) {
    return ListView.separated(
      padding: const EdgeInsets.all(8),
      itemCount: widget.directory.files.length,
      itemBuilder: (BuildContext context, int index) {
        return FileView(file: widget.directory.files[index]);
      },
      separatorBuilder: (BuildContext context, int index) => const SizedBox(height: 10),
    );
  }

}