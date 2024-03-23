# Adapter pattern

https://en.wikipedia.org/wiki/Adapter_pattern

> allows the interface of an existing class to be used as another interface.

既存classのインターフェースを別のインターフェースとして使用できるようにする。

## 何を解決することができるのか

- 互換性の無い既存classを変更することなく、求める要件の機能を実装することができる。
    - テストコードが存在する動作の安定した既存class（動作の担保）を再利用することができる。
- 互換性の無い既存classを組み合わせることができる。

## 構成要素

- Adaptee
    - 既存の処理を持っているclass
- Adapter
    - Adapteeを再利用・組み合わせることで処理を作成するclass

Adapter patternは利用される側のAdapteeの処理が流用できるように作られている前提が必要。
