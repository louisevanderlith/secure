class Register {
  final String name;
  final String email;
  final String password;
  final String confirm;

  Register(this.name, this.email, this.password, this.confirm);

  Map<String, dynamic> toJson() {
    return {
      "Name": this.name,
      "Email": this.email,
      "Password": this.password,
      "PasswordRepeat": this.confirm
    };
  }
}
