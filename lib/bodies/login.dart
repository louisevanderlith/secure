class Login {
  final String client;
  final String name;
  final String password;

  Login(this.client, this.name, this.password);

  Map<String, dynamic> toJson() {
    return {
      "Client": this.client,
      "Username": this.name,
      "Password": this.password
    };
  }
}
