import 'dart:math';

import 'package:client/model/file_node.dart';
import 'package:client/screens/directory_page.dart';
import 'package:flutter/material.dart';

class FileView extends StatefulWidget {
  final FileNode file;

  const FileView({Key? key, required this.file}) : super(key: key);

  @override
  State<StatefulWidget> createState() => _FileView();
}

class _FileView extends State<FileView> {
  @override
  Widget build(BuildContext context) {
    return ListTile(
      leading: _buildLeadingIcon(),
      trailing: _buildTrailing(),
      title: _buildTitle(),
      subtitle: _buildSubtitle(),
      onTap: _buildOnTap(),
    );
  }

  _buildTrailing() {
    return const IconButton(icon: Icon(Icons.download), onPressed: null);
  }

  _buildOnTap() {
    if (widget.file.isDir) {
      return () => {
        Navigator.push(context, MaterialPageRoute(builder: (context) => DirectoryPage(directory: widget.file)))
      };
    }

    return null;
  }

  _buildSubtitle() {
    if (widget.file.isDir) {
      return Text("Fichiers : ${widget.file.files.length}");
    }

    return Text(formatBytes(widget.file.size, 0));
  }

  _buildLeadingIcon() {
    return Icon(widget.file.isDir ? Icons.folder : Icons.insert_drive_file);
  }

  _buildTitle() {
    return Text(widget.file.name == "root" ? "Stockage" : widget.file.name);
  }

  static String formatBytes(int bytes, int decimals) {
    if (bytes <= 0) return "0 o";
    const suffixes = ["o", "Mo", "Ko", "Go", "To", "Po", "Eo", "Zo", "Yo"];
    var i = (log(bytes) / log(1024)).floor();
    return ((bytes / pow(1024, i)).toStringAsFixed(decimals)) +
        ' ' +
        suffixes[i];
  }

}