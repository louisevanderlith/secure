class Resource {
  final String name;
  final String displayname;
  final String secret;
  final List<String> needs;

  Resource(this.name, this.displayname, this.secret, this.needs);

  Map<String, dynamic> toJson() {
    return {
      "Name": this.name,
      "DisplayName": this.displayname,
      "Secret": this.secret,
      "Needs": this.needs
    };
  }
}
