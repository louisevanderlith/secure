import 'package:mango_secure/bodies/client.dart';
import 'package:mango_secure/bodies/contact.dart';
import 'package:mango_ui/keys.dart';

class Profile {
  final String title;
  final String description;
  final List<Contact> contacts;
  final Key imageKey;
  final List<Client> clients;
  final Map<String, String> endpoints;
  final Map<String, String> codes;
  final Map<String, String> terms;

  Profile(this.title, this.description, this.contacts, this.imageKey,
      this.clients, this.endpoints, this.codes, this.terms);

  Map<String, dynamic> toJson() {
    return {
      "Title": this.title,
      "Description": this.description,
      "Contacts": this.contacts,
      "ImageKey": this.imageKey,
      "Clients": this.clients,
      "Endpoints": this.endpoints,
      "Codes": this.codes,
      "Terms": this.terms
    };
  }
}
