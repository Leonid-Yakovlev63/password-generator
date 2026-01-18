# Простой генератор паролей написанный на Go


# Аргументы программы
Возможность сгенерировать сразу несколько паролей, доступен вариант содержищий случаные симовлы (минимальная длина 12 символов) и на основе seed-фразы (12 слов) из словаря BIP39, кол-во паролей должно быть положительным числом.

Варианты использования:
```bash
./password-generator

./password-generator -help

./password-generator -c 5

./password-generator -count 10

./password-generator -count 10 -mode seed

./password-generator -c 15 -m seed

./password-generator -c 20 -m random # random по умолчанию


```

# Полезная информация
## Работа с Makefile в NixOS
```bash
nix-shell -p gnumake
```
