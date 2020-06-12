import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_secure/bodies/login.dart';
import 'package:mango_secure/bodies/register.dart';
import 'package:mango_ui/requester.dart';

Future<HttpRequest> sendLogin(Login obj) async {
  var apiroute = getEndpoint("secure");
  var url = "${apiroute}/login";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> sendForgot(String identity) async {
  var apiroute = getEndpoint("secure");
  var url = "${apiroute}/forgot";
  final data = jsonEncode(identity);

  return invokeService("POST", url, data);
}

Future<HttpRequest> sendRegister(Register obj) async {
  var apiroute = getEndpoint("secure");
  var url = "${apiroute}/register";
  final data = jsonEncode(obj.toJson());

  return invokeService("POST", url, data);
}
