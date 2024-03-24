# Factory method pattern

https://en.wikipedia.org/wiki/Factory_method_pattern

> the factory method pattern is a creational pattern that uses factory methods to deal with the problem of creating objects without having to specify the exact class of the object that will be created.

ファクトリ メソッド パターンは、作成されるオブジェクトの正確なクラスを指定することなく、ファクトリ メソッドを使用してオブジェクト作成の問題に対処する作成パターンです。

- テンプレート メソッドパターンを使用して、インスタンス生成のロジックを抽象的に定義して、具体的な処理内容はサブクラスで定義する。

- `new` による直接的なインスタンス生成をインスタンス生成メソッドに置き換えることで、1つのclassへの依存を解消することができる。

## 何を解決することができるのか

- オブジェクトのインスタンス生成方法をと具体的な生成内容を切り分けることができる。

## 構成要素

- Creator interface
    - インスタンスを生成するファクトリメソッドとインスタンス生成に必要なメソッドを定義する
    - 具体的な処理内容をConcreteCreatorが定義する
- Product interface
    - Creator interfaceで生成されるclass
- ConcreteCreator
    - Creator interfaceを実装したclass
- ConcreteProduct
    - Product interfaceを実装したclass
