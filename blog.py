# 型指定と出力、型確認
print("型指定と出力、型確認")
num = 42

print("改行あり")
print(num)
print(type(num))
print() # 改行

# 改行あり
print("改行あり")
print(42)
print(42)
print() # 改行

# 改行なし
print("改行なし")
print(42, end="")
print(42, end="")
print("\n") # 改行

# フォーマット済み文字列
print("フォーマット済み文字列")
age = 42
print("父は{}歳である".format(age))
print(f"父は{age}歳である")
print("\n") # 改行

# キーボードからの入力を取得
print("キーボードからの入力を取得")
key = input()
print(key)