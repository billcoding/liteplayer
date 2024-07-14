dartClassName="$1"

if [ -z "$dartClassName" ]; then
    echo "Please provide the dart class name."
    exit 0
fi

./.bin/json2dart --dart-class-name="$dartClassName" --dart-file-name=json2dart.dart --json-file-path=json2dart.json