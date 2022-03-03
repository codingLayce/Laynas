class FileNode {
  final String name;
  final int size;
  final bool isDir;
  final List<FileNode> files;

  FileNode(this.name, this.size, this.isDir, this.files);

  factory FileNode.fromJson(dynamic json) {
    List jsonFiles = json["files"];
    
    List<FileNode> files = jsonFiles.map((e) => FileNode.fromJson(e)).toList();

    return FileNode(json["name"], json["size"], json["isDir"], files);
  }
}