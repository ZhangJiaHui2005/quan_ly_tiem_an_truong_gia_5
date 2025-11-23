import 'dart:convert';

import 'package:frontend/models/item.dart';
import 'package:http/http.dart' as http;

class ItemsServices {
  static const String baseUrl = "http://localhost:8000";

  static Future<List<Item>> fetchItems() async {
    final url = Uri.parse("$baseUrl/api/items");

    final response = await http.get(url);

    if (response.statusCode == 200) {
      final data = json.decode(response.body);
      final List items = data["items"];

      return items.map((item) => Item.fromJson(item)).toList();
    } else {
      throw Exception("Failed to load items");
    }
  }
}