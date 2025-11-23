import 'package:frontend/models/category.dart';

class Item {
  final int id;
  final String name;
  final double price;
  final int categoryID;
  final Category? category;

  Item({
    required this.id,
    required this.name,
    required this.price,
    required this.categoryID,
    required this.category,
  });

  factory Item.fromJson(Map<String, dynamic> json) {
    return Item(
      id: json["id"],
      name: json["name"],
      price: json["price"].toDouble(),
      categoryID: json["category_id"],
      category: json["category"] != null
          ? Category.fromJson(json["category"])
          : null,
    );
  }
}
