import 'dart:convert';
import 'dart:html';

import 'package:mango_secure/bodies/profile.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

Future<HttpRequest> createProfile(Profile obj) async {
  var apiroute = getEndpoint("secure");
  var url = "${apiroute}/profile";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateProfile(Key k, Profile obj) async {
  var apiroute = getEndpoint("secure");
  var url = "${apiroute}/profile/${k.toJson()}";

  return invokeService("PUT", url, jsonEncode(obj.toJson()));
}