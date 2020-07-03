import 'dart:convert';
import 'dart:html';

import 'package:mango_secure/bodies/resource.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

Future<HttpRequest> createResource(Resource obj) async {
  var apiroute = getEndpoint("secure");
  var url = "${apiroute}/resources";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateResource(Key k, Resource obj) async {
  var apiroute = getEndpoint("secure");
  var url = "${apiroute}/resources/${k.toJson()}";

  return invokeService("PUT", url, jsonEncode(obj.toJson()));
}
