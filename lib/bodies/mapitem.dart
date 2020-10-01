import 'dart:html';

List<MapItem> getMapItems(String nameElem, String valueElem) {
  var isLoaded = false;
  var result = new List<MapItem>();
  var index = 0;

  do {
    var item = new MapItem('${nameElem}${index}', "${valueElem}${index}");

    isLoaded = item.loaded;
    if (isLoaded) {
      result.add(item);
    }

    index++;
  } while (isLoaded);

  return result;
}

class MapItem {
  TextInputElement txtName;
  TextInputElement txtValue;
  bool _loaded;

  MapItem(String nameId, String valueId) {
    txtName = querySelector(nameId);
    txtValue = querySelector(valueId);

    _loaded = txtName != null && txtValue != null;
  }

  bool get loaded {
    return _loaded;
  }

  String get name {
    return txtName.value;
  }

  String get value {
    return txtValue.value;
  }

  Map<String, dynamic> toJson() {
    return {
      "K": this.name,
      "V": this.value,
    };
  }
}
