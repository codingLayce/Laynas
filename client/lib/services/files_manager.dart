import 'dart:convert';

import 'package:client/model/file_node.dart';
import 'package:http/http.dart' as http;

const backendURI = "http://localhost:20000";
const secretToken = "QWxlZDE3";

class FileManager {
  static final FileManager _instance = FileManager._internal();

  factory FileManager() {
    return _instance;
  }

  Future<FileNode> getFileTree() async {
    final response = await http.get(Uri.parse(backendURI + "/files"), headers: {"Authorization": secretToken});
    
    if (response.statusCode == 200) {
      final decoded = jsonDecode(response.body);
      return FileNode.fromJson(decoded);
    }

    return Future.error("status code isn't 200");
  }

  FileManager._internal();
}