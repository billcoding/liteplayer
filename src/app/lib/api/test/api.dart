import 'package:httper/httper.dart';
import 'model.dart';

abstract class TestApi {
  static Future<Model<Test>> get({
    Callback200<Test>? callback200,
    CallbackOther<Test>? callbackOther,
  }) async =>
      await Http.get(
        'http://localhost:8080/api/echo',
        Test.fromMap,
        callback200: callback200,
        callbackOther: callbackOther,
      );
}
