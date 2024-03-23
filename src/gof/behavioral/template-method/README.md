# Template method pattern

https://en.wikipedia.org/wiki/Template_method_pattern

> The intent of the template method is to define the overall structure of the operation, while allowing subclasses to refine, or redefine, certain steps

テンプレート メソッドの目的は、サブクラスが特定のステップを改良または再定義できるようにしながら、操作の全体的な構造を定義することです。

- 大枠の処理構成（テンプレート）を抽象的に定義して、構成内で実行される具体的な処理内容をサブクラスで定義する。

- サブクラスで定義する必要がある処理は、下記などで定義する
    - abstract methods
    - hook methods（スーパークラス内でオーバーライドされる前提の処理内容が空のメソッドを定義する）

> This pattern is an example of inversion of control because the high-level code no longer determines what algorithms to run

このパターンは、高レベルのコードが実行するアルゴリズムを決定しなくなるため、制御の逆転の一例

> The template pattern is useful when working with auto-generated code.

テンプレートパターンはコード生成と相性が良い。

## 何を解決することができるのか

- 処理の構成を共通化することでコードの重複を避けることができる。
- 具体的な処理のバリエーションをサブクラスに分離することができる。

## 構成要素

- Abstract
    - 大枠の処理構成をテンプレートメソッドとして定義する。
- Concreate
    - テンプレートメソッド内で実行されるステップ処理（＝具体的な処理）を定義する。
