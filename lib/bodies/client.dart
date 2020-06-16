class Client {
  final String name;
  final String secret;
  final String url;
  final List<String> resources;
  final bool terms;
  final bool codes;

  Client(this.name, this.secret, this.url, this.resources, this.terms, this.codes);

  Map<String, dynamic> toJson() {
    return {
      "Name": this.name,
      "Secret": this.secret,
      "Url": this.url,
      "AllowedResources": this.resources,
      "TermsEnabled": this.terms,
      "CodesEnabled": this.codes
    };
  }
}