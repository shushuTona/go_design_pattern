# Iterator pattern

https://en.wikipedia.org/wiki/Iterator_pattern

> The iterator pattern decouples algorithms from containers

イテレーターパターンを使用することで、繰り返し処理の対象からアルゴリズムを分離することができる。

## 何を解決することができるのか

- 繰り返し処理のアルゴリズムをカプセル化することができる
- 対象データの構造を知らなくても、共通化された繰り返し処理を実行することができる

## 構成要素

- Aggregate interface
    - Iterator interfaceの生成処理が定義されているインターフェース
- Iterator interface
    - Aggregate interfaceを走査するための処理が定義されているインターフェース
- ConcreteAggregate
    - Aggregate interfaceを実装して、ConcreteIteratorを生成する。
- ConcreteIterator
    - Iterator interfaceを実装して、ConcreteAggregateを走査する。
